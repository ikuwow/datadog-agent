// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !linux

package ebpf

import (
	"github.com/prometheus/client_golang/prometheus"
)

type LockContentionCollector struct{}

// NewLockContentionCollector returns nil
func NewLockContentionCollector() *LockContentionCollector {
	return nil
}
