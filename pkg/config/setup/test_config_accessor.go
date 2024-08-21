// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build test

package setup

import (
	"sync"

	pkgconfigmodel "github.com/DataDog/datadog-agent/pkg/config/model"
)

// This file contains the accessor for the configuration when running tests. This version offer a way to set the
// different configurations at runtime.

var (
	datadogMutex     = sync.RWMutex{}
	systemProbeMutex = sync.RWMutex{}
)

// Datadog returns the current agent configuration
func Datadog() pkgconfigmodel.Config {
	datadogMutex.RLock()
	defer datadogMutex.RUnlock()
	return datadog
}

// SetDatadog sets the the reference to the agent configuration.
// This is currently used by the legacy converter and config mocks and should not be user anywhere else. Once the
// legacy converter and mock have been migrated we will remove this function.
func SetDatadog(cfg pkgconfigmodel.Config) {
	datadogMutex.Lock()
	defer datadogMutex.Unlock()
	datadog = cfg
}

// SystemProbe returns the current SystemProbe configuration
func SystemProbe() pkgconfigmodel.Config {
	systemProbeMutex.RLock()
	defer systemProbeMutex.RUnlock()
	return systemProbe
}

// SetSystemProbe sets the the reference to the systemProbe configuration.
// This is currently used by the config mocks and should not be user anywhere else. Once the mocks have been migrated we
// will remove this function.
func SetSystemProbe(cfg pkgconfigmodel.Config) {
	systemProbeMutex.Lock()
	defer systemProbeMutex.Unlock()
	systemProbe = cfg
}
