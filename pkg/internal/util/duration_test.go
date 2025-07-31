package util

import (
	"testing"
)

func TestChargeDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration uint32
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "00:00:00",
		},
		{
			name:     "one second",
			duration: 1,
			expected: "00:00:01",
		},
		{
			name:     "one minute",
			duration: 60,
			expected: "00:01:00",
		},
		{
			name:     "one hour",
			duration: 3600,
			expected: "01:00:00",
		},
		{
			name:     "complex duration",
			duration: 3661, // 1 hour, 1 minute, 1 second
			expected: "01:01:01",
		},
		{
			name:     "large duration",
			duration: 86461, // 24 hours, 1 minute, 1 second
			expected: "24:01:01",
		},
		{
			name:     "typical charge duration",
			duration: 7380, // 2 hours, 3 minutes
			expected: "02:03:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChargeDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("ChargeDuration(%d) = %s, expected %s", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationFormat(t *testing.T) {
	tests := []struct {
		name     string
		time     float64
		expected string
	}{
		{
			name:     "single digit",
			time:     5.0,
			expected: "05",
		},
		{
			name:     "double digit",
			time:     15.0,
			expected: "15",
		},
		{
			name:     "zero",
			time:     0.0,
			expected: "00",
		},
		{
			name:     "nine",
			time:     9.0,
			expected: "9",
		},
		{
			name:     "ten",
			time:     10.0,
			expected: "10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := durationFormat(tt.time)
			if result != tt.expected {
				t.Errorf("durationFormat(%f) = %s, expected %s", tt.time, result, tt.expected)
			}
		})
	}
}
