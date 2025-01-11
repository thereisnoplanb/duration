package duration

import (
	"testing"
)

func TestDuration_String(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		want     string
	}{
		{
			name:     "0",
			duration: Zero,
			want:     "00:00:00",
		},
		{
			name:     "1 microsecond",
			duration: 1 * Microsecond,
			want:     "00:00:00.0000010",
		},
		{
			name:     "1 millisecond",
			duration: 1 * Millisecond,
			want:     "00:00:00.0010000",
		},
		{
			name:     "1 second",
			duration: 1 * Second,
			want:     "00:00:01",
		},
		{
			name:     "1 minute",
			duration: 1 * Minute,
			want:     "00:01:00",
		},
		{
			name:     "1 hour",
			duration: 1 * Hour,
			want:     "01:00:00",
		},
		{
			name:     "1 day",
			duration: 1 * Day,
			want:     "1.00:00:00",
		},
		{
			name:     "7 days",
			duration: 1 * Week,
			want:     "7.00:00:00",
		},
		{
			name:     "10 days",
			duration: 10 * Day,
			want:     "10.00:00:00",
		},
		{
			name:     "Maximum",
			duration: Maximum,
			want:     "10675199.02:48:05.4775807",
		},
		{
			name:     "Minimum",
			duration: Minimum,
			want:     "-10675199.02:48:05.4775808",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.duration.String(); got != tt.want {
				t.Errorf("Duration.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_Days(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		want     int
	}{
		{
			name:     "1.5 days",
			duration: 1*Day + 12*Hour,
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.duration.Days(); got != tt.want {
				t.Errorf("Duration.Days() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_TotalDays(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		want     float64
	}{
		{
			name:     "1.5 days",
			duration: 1*Day + 12*Hour,
			want:     1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.duration.TotalDays(); got != tt.want {
				t.Errorf("Duration.TotalDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
