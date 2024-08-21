// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !test

package setup

import pkgconfigmodel "github.com/DataDog/datadog-agent/pkg/config/model"

// This file contains the accessor for the configuration when NOT running tests. This version doesn't allow outside
// caller to change the configuration reference at runtime and doesn't use mutexes.

// Datadog returns the current agent configuration
func Datadog() pkgconfigmodel.Config {
	return datadog
}

// SystemProbe returns the current SystemProbe configuration
func SystemProbe() pkgconfigmodel.Config {
	return systemProbe
}
