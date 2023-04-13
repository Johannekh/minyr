package yr_test

import (
	"testing"

	"github.com/Johannekh/minyr/yr"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "file with 25 lines",
			filename: "kjevik-temp-celsius-20220318-20230318.csv",
			want:     25,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := yr.CountLines(tt.filename)
			if got != tt.want {
				t.Errorf("CountLines(%s) = %d, want %d", tt.filename, got, tt.want)
			}
		})
	}
}

func TestGetAverageTemperature(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		temperature string
		want        string
	}{
		{
			name:        "file with celsius temperature",
			filename:    "kjevik-temp-celsius-20220318-20230318.csv",
			temperature: "celsius",
			want:        "-0.60",
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yr.GetAverageTemperature(tt.filename, tt.temperature)
			if err != nil {
				t.Fatalf("GetAverageTemperature(%s, %s) returned an error: %v", tt.filename, tt.temperature, err)
			}
			if got != tt.want {
				t.Errorf("GetAverageTemperature(%s, %s) = %s, want %s", tt.filename, tt.temperature, got, tt.want)
			}
		})
	}
}
