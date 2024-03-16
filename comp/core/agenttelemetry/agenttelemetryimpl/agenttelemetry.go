// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// ---------------------------------------------------
//
// This is experimental code and is subject to change.
//
// ---------------------------------------------------

// Package agenttelemetryimpl provides the implementation of the agenttelemetry component.
package agenttelemetryimpl

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/DataDog/datadog-agent/comp/core/agenttelemetry"
	"github.com/DataDog/datadog-agent/comp/core/config"
	"github.com/DataDog/datadog-agent/comp/core/log"
	"github.com/DataDog/datadog-agent/comp/core/status"
	"github.com/DataDog/datadog-agent/comp/core/telemetry"
	"github.com/DataDog/datadog-agent/comp/metadata/host"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
	"golang.org/x/exp/maps"

	dto "github.com/prometheus/client_model/go"
	"go.uber.org/fx"
)

// Embed one or more rendering templated into this binary as a resource
// to be used at runtime.

//go:embed status_templates
var templatesFS embed.FS

// Module defines the fx options for this component.
func Module() fxutil.Module {
	return fxutil.Component(
		fx.Provide(newAtel))
}

type atel struct {
	cfgComp    config.Component
	logComp    log.Component
	telComp    telemetry.Component
	statusComp status.Component
	hostComp   host.Component

	enabled bool
	sender  sender
	runner  runner
	atelCfg *Config

	cancelCtx context.Context
	cancel    context.CancelFunc
}

// FX-in compatibility
type dependencies struct {
	fx.In

	Log       log.Component
	Config    config.Component
	Telemetry telemetry.Component
	Status    status.Component
	Host      host.Component

	Lifecycle fx.Lifecycle
}

// Interfacing with runner.
type job struct {
	a        *atel
	profiles []*Profile
	schedule Schedule
}

func (j job) Run() {
	j.a.run(j.profiles)
}

// Passing metrics to sender Interfacing with sender
type agentmetric struct {
	name    string
	metrics []*dto.Metric
	family  *dto.MetricFamily
}

func createSender(
	cfgComp config.Component,
	logComp log.Component,
	hostComp host.Component,
) (sender, error) {
	client := newSenderClientImpl(cfgComp)
	sender, err := newSenderImpl(cfgComp, logComp, hostComp, client)
	if err != nil {
		logComp.Errorf("Failed to create agent telemetry sender: %s", err.Error())
	}
	return sender, err
}

func createAtel(
	cfgComp config.Component,
	logComp log.Component,
	telComp telemetry.Component,
	statusComp status.Component,
	hostComp host.Component,
	sender sender,
	runner runner) *atel {
	// Parse Agent Telemetry Configuration configuration
	atelCfg, err := parseConfig(cfgComp)
	if err != nil {
		logComp.Errorf("Failed to parse agent telemetry config: %s", err.Error())
		return &atel{}
	}
	if !atelCfg.Enabled {
		logComp.Info("Agent telemetry is disabled")
		return &atel{}
	}

	return &atel{
		enabled:    true,
		cfgComp:    cfgComp,
		logComp:    logComp,
		telComp:    telComp,
		statusComp: statusComp,
		hostComp:   hostComp,
		sender:     sender,
		runner:     runner,
		atelCfg:    atelCfg,
	}
}

func newAtel(deps dependencies) agenttelemetry.Component {
	sender, err := createSender(deps.Config, deps.Log, deps.Host)
	if err != nil {
		return &atel{}
	}

	runner := newRunnerImpl()

	// Wire up the agent telemetry provider (TODO: use FX for sender, client and runner?)
	a := createAtel(
		deps.Config,
		deps.Log,
		deps.Telemetry,
		deps.Status,
		deps.Host,
		sender,
		runner,
	)

	// If agent telemetry is enabled, add the start and stop hooks
	if a.enabled {
		// Instruct FX to start and stop the agent telemetry
		deps.Lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return a.start()
			},
			OnStop: func(ctx context.Context) error {
				return a.stop()
			},
		})
	}

	return a
}

func (a *atel) aggregateMetricTags(mCfg *MetricConfig, mt dto.MetricType, ms []*dto.Metric) []*dto.Metric {
	// Nothing to aggregate?
	if len(ms) == 0 {
		return nil
	}

	// Special case when no aggregate tags are defined - aggregate all metrics
	// aggregateMetric will sum all metrics into a single one without copying tags
	if !mCfg.aggregateTagsExists {
		ma := &dto.Metric{}
		for _, m := range ms {
			aggregateMetric(mt, ma, m)
		}

		return []*dto.Metric{ma}
	}

	amMap := make(map[string]*dto.Metric)

	// Initialize total metric
	var totalm *dto.Metric
	if mCfg.AggregateTotal {
		totalm = &dto.Metric{}
	}

	// Enumerate the metric's timeseries and aggregate them
	for _, m := range ms {
		tagsKey := ""

		// if tags are defined, we need to create a key from them by dropping not specified
		// in configuration tags. The key is constructed by conatenating specified tag names and values
		// if the a timeseries has tags is not specified in
		origTags := m.GetLabel()
		if len(origTags) > 0 {
			// sort tags (to have a consistent key for the same tag set)
			tags := cloneLabelsSorted(origTags)

			// create a key from the tags (and drop not specified in the configuration tags)
			var specTags = make([]*dto.LabelPair, 0, len(origTags))
			for _, t := range tags {
				if _, ok := mCfg.aggregateTagsMap[t.GetName()]; ok {
					specTags = append(specTags, t)
					tagsKey += makeLabelPairKey(t)
				}
			}
			if mCfg.AggregateTotal {
				aggregateMetric(mt, totalm, m)
			}

			// finally aggregate the metric on the created key
			if aggm, ok := amMap[tagsKey]; ok {
				aggregateMetric(mt, aggm, m)
			} else {
				// ... or create a new one with specifi value and specified tags
				aggm := &dto.Metric{}
				aggregateMetric(mt, aggm, m)
				aggm.Label = specTags
				amMap[tagsKey] = aggm
			}
		} else {
			// if no tags are specified, we aggregate all metrics into a single one
			if mCfg.AggregateTotal {
				aggregateMetric(mt, totalm, m)
			}
		}
	}

	// Add total metric if needed
	if mCfg.AggregateTotal {
		totalName := "total"
		totalValue := strconv.Itoa(len(ms))
		totalm.Label = []*dto.LabelPair{
			{Name: &totalName, Value: &totalValue},
		}
		amMap[totalName] = totalm
	}

	// Anything to report?
	if len(amMap) == 0 {
		return nil
	}

	// Convert the map to a slice
	return maps.Values(amMap)
}

func isMetricFiltered(p *Profile, mCfg *MetricConfig, mt dto.MetricType, m *dto.Metric) bool {
	// filter out zero values if specified in the profile
	if p.excludeZeroMetric && isZeroValueMetric(mt, m) {
		return false
	}

	// filter out if contains excluded tags
	if len(p.excludeTagsMap) > 0 && areTagsMatching(m.GetLabel(), p.excludeTagsMap) {
		return false
	}

	// filter out if tag does not contain in existing aggregateTags
	if mCfg.aggregateTagsExists && !areTagsMatching(m.GetLabel(), mCfg.aggregateTagsMap) {
		return false
	}

	return true
}

func (a *atel) transformMetricFamily(p *Profile, mfam *dto.MetricFamily) *agentmetric {
	var mCfg *MetricConfig
	var ok bool

	// Check if the metric is included in the profile
	if mCfg, ok = p.metricsMap[mfam.GetName()]; !ok {
		return nil
	}

	// Filter out not supported types
	mt := mfam.GetType()
	if !isSupportedMetricType(mt) {
		return nil
	}

	// Filter the metric according to the profile configuration
	// Currently we only support filtering out zero values if specified in the profile
	var fm []*dto.Metric
	for _, m := range mfam.Metric {
		if isMetricFiltered(p, mCfg, mt, m) {
			fm = append(fm, m)
		}
	}

	amt := a.aggregateMetricTags(mCfg, mt, fm)

	// nothing to report
	if len(fm) == 0 {
		return nil
	}

	return &agentmetric{
		name:    mCfg.Name,
		metrics: amt,
		family:  mfam,
	}
}

func (a *atel) reportAgentMetrics(session *senderSession, p *Profile) {
	// If no metrics are configured nothing to report
	if len(p.metricsMap) == 0 {
		return
	}

	a.logComp.Infof("Collect Agent Metric telemetry for profile %s", p.Name)

	// Gather all prom metrircs. Currently Gather() does not allow filtering by
	// matric name, so we need to gather all metrics and filter them on our own.
	//	pms, err := a.telemetry.Gather(false)
	pms, err := a.telComp.Gather(false)
	if err != nil {
		a.logComp.Errorf("failed to get filtered telemetry metrics: %s", err)
		return
	}

	// ... and filter them according to the profile configuration
	var metrics []*agentmetric
	for _, pm := range pms {
		if am := a.transformMetricFamily(p, pm); am != nil {
			metrics = append(metrics, am)
		}
	}

	// Send the metrics if any were filtered
	if len(metrics) == 0 {
		a.logComp.Info("No Agent Metric telemetry collected")
		return
	}

	// Send the metrics if any were filtered
	a.logComp.Infof("Reporting Agent Metric telemetry for profile %s", p.Name)

	err = a.sender.sendAgentMetricPayloads(session, metrics)
	if err != nil {
		a.logComp.Errorf("failed to get filtered telemetry metrics: %s", err)
	}
}

// renderAgentStatus renders (transform) input status JSON object into output status using the template
func (a *atel) renderAgentStatus(p *Profile, inputStatus map[string]interface{}, statusOutput map[string]interface{}) {
	// Render template if needed
	if p.Status.Template == "none" {
		return
	}

	templateName := "agent-telemetry-" + p.Status.Template + ".tmpl"
	var b = new(bytes.Buffer)
	err := status.RenderText(templatesFS, templateName, b, inputStatus)
	if err != nil {
		a.logComp.Errorf("Failed to collect Agent Status telemetry. Error: %s", err.Error())
		return
	}
	if len(b.Bytes()) == 0 {
		a.logComp.Info("Agent status rendering to agent telemetry payloads return empty payload")
		return
	}

	// Convert byte slice to JSON object
	if err := json.Unmarshal(b.Bytes(), &statusOutput); err != nil {
		a.logComp.Errorf("Failed to collect Agent Status telemetry. Error: %s", err.Error())
		return
	}
}

func (a *atel) addAgentStatusExtra(p *Profile, fullStatus map[string]interface{}, statusOutput map[string]interface{}) {
	for _, builder := range p.statusExtraBuilder {
		// Evaluate JQ expression against the agent status JSON object
		jqResult := builder.jqSource.Run(fullStatus)
		jqValue, ok := jqResult.Next()
		if !ok {
			a.logComp.Errorf("Failed to apply JQ expression for JSON path '%s' to Agent Status payload. Error unknown",
				strings.Join(builder.jpathTarget, "."))
			continue
		}

		// Validate JQ expression result
		if err, ok := jqValue.(error); ok {
			a.logComp.Errorf("Failed to apply JQ expression for JSON path '%s' to Agent Status payload. Error: %s",
				strings.Join(builder.jpathTarget, "."), err.Error())
			continue
		}

		// Validate JQ expression result type
		var attrVal interface{}
		switch jqValueType := jqValue.(type) {
		case int:
			attrVal = jqValueType
		case float64:
			attrVal = jqValueType
		case bool:
			attrVal = jqValueType
		case nil:
			a.logComp.Info("JQ expression return 'nil' value for JSON path '%s'", strings.Join(builder.jpathTarget, "."))
			continue
		case string:
			a.logComp.Errorf("string value (%v) for JSON path '%s' for extra status atttribute is not currently allowed",
				strings.Join(builder.jpathTarget, "."), attrVal)
			continue
		default:
			a.logComp.Errorf("'%v' value (%v) for JSON path '%s' for extra status atttribute is not currently allowed",
				reflect.TypeOf(jqValueType), reflect.ValueOf(jqValueType), strings.Join(builder.jpathTarget, "."))
			continue
		}

		// Add resulting value to the agent status telemetry payload (recursively creating missing JSON objects)
		curNode := statusOutput
		for i, p := range builder.jpathTarget {
			// last element is the attribute name
			if i == len(builder.jpathTarget)-1 {
				curNode[p] = attrVal
				break
			}

			existSubNode, ok := curNode[p]

			// if the node doesn't exist, create it
			if !ok {
				newSubNode := make(map[string]interface{})
				curNode[p] = newSubNode
				curNode = newSubNode
			} else {
				existSubNodeCasted, ok := existSubNode.(map[string]interface{})
				if !ok {
					a.logComp.Errorf("JSON path '%s' points to non-object element", strings.Join(builder.jpathTarget[:i], "."))
					break
				}
				curNode = existSubNodeCasted
			}
		}
	}
}

func (a *atel) reportAgentStatus(session *senderSession, p *Profile) {
	// If no status is configured nothing to report
	if p.Status == nil {
		return
	}

	a.logComp.Info("Collect Agent Status telemetry for profile %s", p.Name)

	// Current "agent-telemetry-basic.tmpl" uses only "runneStats" and "dogstatsdStats" JSON sections
	// These JSON sections are populated via "collector" and "DogStatsD" status providers sections
	minimumReqSections := []string{"collector", "DogStatsD"}
	statusBytes, err := a.statusComp.GetStatusBySection(minimumReqSections, "json", false)
	if err != nil {
		a.logComp.Errorf("failed to get agent status: %s", err)
		return
	}

	var statusJSON = make(map[string]interface{})
	err = json.Unmarshal(statusBytes, &statusJSON)
	if err != nil {
		a.logComp.Errorf("failed to unmarshall agent status: %s", err)
		return
	}

	// Render Agent Status JSON object (using template if needed and adding extra attributes)
	var statusPayloadJSON = make(map[string]interface{})
	a.renderAgentStatus(p, statusJSON, statusPayloadJSON)
	a.addAgentStatusExtra(p, statusJSON, statusPayloadJSON)
	if len(statusPayloadJSON) == 0 {
		a.logComp.Info("No Agent Status telemetry collected")
		return
	}

	a.logComp.Info("Reporting Agent Status telemetry for profile %s", p.Name)

	// Send the Agent Telemetry status payload
	err = a.sender.sendAgentStatusPayload(session, statusPayloadJSON)
	if err != nil {
		a.logComp.Errorf("failed to send agent status: %s", err)
		return
	}
}

// run runs the agent telemetry for a given profile. It is triggered by the runner
// according to the profiles schedule.
func (a *atel) run(profiles []*Profile) {
	a.logComp.Info("Starting agent telemetry run")
	session := a.sender.startSession(a.cancelCtx)

	for _, p := range profiles {
		a.reportAgentMetrics(session, p)
		a.reportAgentStatus(session, p)
	}

	err := a.sender.flushSession(session)
	if err != nil {
		a.logComp.Errorf("failed to flush agent telemetry session: %s", err)
		return
	}
}

// TODO: implement when CLI tool will be implemented
func (a *atel) GetAsJSON() ([]byte, error) {
	return nil, nil
}

// start is called by FX when the application starts.
func (a *atel) start() error {
	a.logComp.Info("Starting agent telemetry")

	a.cancelCtx, a.cancel = context.WithCancel(context.Background())

	// Start the runner and add the jobs.
	a.runner.start()
	for sh, pp := range a.atelCfg.schedule {
		// get string representation of profiles names
		pnames := make([]string, len(pp))
		for i, p := range pp {
			pnames[i] = p.Name
		}
		a.logComp.Infof("Adding job for schedule[%d, %d, %d] for profile",
			sh.StartAfter, sh.Iterations, sh.Period, " with profiles: ", strings.Join(pnames, ", "))

		a.runner.addJob(job{
			a:        a,
			profiles: pp,
			schedule: sh,
		})
	}

	return nil
}

// stop is called by FX when the application stops.
func (a *atel) stop() error {
	a.logComp.Info("Stopping agent telemetry")
	a.cancel()

	runnerCtx := a.runner.stop()
	<-runnerCtx.Done()

	<-a.cancelCtx.Done()
	a.logComp.Info("Agent telemetry is stopped")
	return nil
}
