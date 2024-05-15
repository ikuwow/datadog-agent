// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

package environments

import (
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/components"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/e2e"
	"github.com/DataDog/test-infra-definitions/resources/local/docker"
)

// DockerLocal is a *local* Docker environment that contains a Host, FakeIntake and Agent configured to talk to each other.
type DockerLocal struct {
	Local *docker.Environment
	// Components
	RemoteHost *components.RemoteHost
	FakeIntake *components.FakeIntake
	Agent      *components.RemoteHostAgent
	Updater    *components.RemoteHostUpdater
}

var _ e2e.Initializable = &DockerLocal{}

// Init initializes the environment
func (e *DockerLocal) Init(_ e2e.Context) error {
	return nil
}
