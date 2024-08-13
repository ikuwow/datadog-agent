// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017-present Datadog, Inc.

//go:build !serverless

package listeners

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/core/config"
	log "github.com/DataDog/datadog-agent/comp/core/log/def"
	logmock "github.com/DataDog/datadog-agent/comp/core/log/mock"
	workloadmeta "github.com/DataDog/datadog-agent/comp/core/workloadmeta/def"
	workloadmetafxmock "github.com/DataDog/datadog-agent/comp/core/workloadmeta/fx-mock"
	workloadmetamock "github.com/DataDog/datadog-agent/comp/core/workloadmeta/mock"
	"github.com/DataDog/datadog-agent/pkg/util/containers"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

type wlmListenerSvc struct {
	service Service
	parent  string
}

type testWorkloadmetaListener struct {
	t        *testing.T
	filters  *containerFilters
	store    workloadmeta.Component
	services map[string]wlmListenerSvc
}

//nolint:revive // TODO(CINT) Fix revive linter
func (l *testWorkloadmetaListener) Listen(_ chan<- Service, _ chan<- Service) {
	panic("not implemented")
}

//nolint:revive // TODO(CINT) Fix revive linter
func (l *testWorkloadmetaListener) Stop() {
	panic("not implemented")
}

//nolint:revive // TODO(CINT) Fix revive linter
func (l *testWorkloadmetaListener) Store() workloadmeta.Component {
	return l.store
}

//nolint:revive // TODO(CINT) Fix revive linter
func (l *testWorkloadmetaListener) AddService(svcID string, svc Service, parentSvcID string) {
	l.services[svcID] = wlmListenerSvc{
		service: svc,
		parent:  parentSvcID,
	}
}

//nolint:revive // TODO(CINT) Fix revive linter
func (l *testWorkloadmetaListener) IsExcluded(ft containers.FilterType, annotations map[string]string, name string, image string, ns string) bool {
	return l.filters.IsExcluded(ft, annotations, name, image, ns)
}

func (l *testWorkloadmetaListener) assertServices(expectedServices map[string]wlmListenerSvc) {
	for svcID, expectedSvc := range expectedServices {
		actualSvc, ok := l.services[svcID]
		if !ok {
			l.t.Errorf("expected to find service %q, but it was not generated", svcID)
			continue
		}

		assert.Equal(l.t, expectedSvc, actualSvc)

		delete(l.services, svcID)
	}

	if len(l.services) > 0 {
		l.t.Errorf("got unexpected services: %+v", l.services)
	}
}

func newTestWorkloadmetaListener(t *testing.T) *testWorkloadmetaListener {
	filters, err := newContainerFilters()
	if err != nil {
		t.Fatalf("cannot initialize container filters: %s", err)
	}

	w := fxutil.Test[workloadmetamock.Mock](t, fx.Options(
		fx.Supply(config.Params{}),
		fx.Supply(log.Params{}),
		fx.Provide(func() log.Component { return logmock.New(t) }),
		config.MockModule(),
		fx.Supply(context.Background()),
		fx.Supply(workloadmeta.NewParams()),
		workloadmetafxmock.MockModule(),
	))

	return &testWorkloadmetaListener{
		t:        t,
		filters:  filters,
		store:    w,
		services: make(map[string]wlmListenerSvc),
	}
}
