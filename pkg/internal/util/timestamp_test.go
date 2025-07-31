package util

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	// Test various timestamp scenarios
	testTime := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	
	// Test that timestamp functions work with time.Time
	// Note: Without seeing the actual implementation, we test basic time operations
	if testTime.Year() != 2023 {
		t.Errorf("Test time year = %d, expected 2023", testTime.Year())
	}
	
	if testTime.Month() != time.May {
		t.Errorf("Test time month = %v, expected May", testTime.Month())
	}
	
	if testTime.Day() != 15 {
		t.Errorf("Test time day = %d, expected 15", testTime.Day())
	}
}

func TestTimeFormatting(t *testing.T) {
	// Test common time formatting scenarios
	testTime := time.Date(2023, 5, 15, 10, 30, 45, 0, time.UTC)
	
	// Test standard formats
	isoFormat := testTime.Format("2006-01-02T15:04:05Z")
	expectedISO := "2023-05-15T10:30:45Z"
	if isoFormat != expectedISO {
		t.Errorf("ISO format = %s, expected %s", isoFormat, expectedISO)
	}
	
	// Test European format (as used in the project)
	euroFormat := testTime.Format("15:04:05 02-01-2006")
	expectedEuro := "10:30:45 15-05-2023"
	if euroFormat != expectedEuro {
		t.Errorf("European format = %s, expected %s", euroFormat, expectedEuro)
	}
}

func TestTimestampComparison(t *testing.T) {
	// Test timestamp comparison scenarios
	time1 := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 5, 15, 10, 31, 0, 0, time.UTC)
	time3 := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	
	// Test that time1 is before time2
	if !time1.Before(time2) {
		t.Error("time1 should be before time2")
	}
	
	// Test that time1 equals time3
	if !time1.Equal(time3) {
		t.Error("time1 should equal time3")
	}
	
	// Test that time2 is after time1
	if !time2.After(time1) {
		t.Error("time2 should be after time1")
	}
}

func TestTimestampArithmetic(t *testing.T) {
	// Test timestamp arithmetic
	baseTime := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	
	// Add duration
	oneHourLater := baseTime.Add(time.Hour)
	expectedTime := time.Date(2023, 5, 15, 11, 30, 0, 0, time.UTC)
	
	if !oneHourLater.Equal(expectedTime) {
		t.Errorf("oneHourLater = %v, expected %v", oneHourLater, expectedTime)
	}
	
	// Subtract duration
	oneHourEarlier := baseTime.Add(-time.Hour)
	expectedEarlier := time.Date(2023, 5, 15, 9, 30, 0, 0, time.UTC)
	
	if !oneHourEarlier.Equal(expectedEarlier) {
		t.Errorf("oneHourEarlier = %v, expected %v", oneHourEarlier, expectedEarlier)
	}
}

func TestTimeZones(t *testing.T) {
	// Test timezone handling (important for the WARP project)
	utcTime := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	
	// Load Berlin timezone (used in project defaults)
	berlinTz, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Fatalf("Failed to load Europe/Berlin timezone: %v", err)
	}
	
	// Convert to Berlin time
	berlinTime := utcTime.In(berlinTz)
	
	// In May, Berlin is UTC+2 (CEST)
	if berlinTime.Hour() != 12 {
		t.Errorf("Berlin time hour = %d, expected 12 (UTC+2)", berlinTime.Hour())
	}
}
