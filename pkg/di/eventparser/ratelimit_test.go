// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.
package eventparser

import (
	"testing"

	"gotest.tools/assert"
)

func TestRateLimit(t *testing.T) {

	testCases := []struct {
		name           string
		limitPerSecond float64
	}{
		{
			name:           "expected1",
			limitPerSecond: 1.0,
		},
		{
			name:           "expected2",
			limitPerSecond: 5.0,
		},
	}

	for _, testcase := range testCases {

		const timesToRun = 10000
		t.Run(testcase.name, func(t *testing.T) {

			r := newSingleEventRateLimiter(testcase.limitPerSecond)

			for i := 0; i < timesToRun; i++ {
				r.allowOneEvent()
			}

			assert.Equal(t, timesToRun-float64(r.droppedEvents), testcase.limitPerSecond)
			assert.Equal(t, r.droppedEvents, timesToRun-testcase.limitPerSecond)
			assert.Equal(t, r.successfulEvents, testcase.limitPerSecond)
		})
	}
}
