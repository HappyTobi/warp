package meter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewMeterService(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	meterService := NewMeterService(request)
	if meterService == nil {
		t.Error("NewMeterService() returned nil")
	}
	if meterService.request != request {
		t.Error("NewMeterService() request not set correctly")
	}
}

func TestMeter_Values(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/meter/values" {
			t.Errorf("Expected path /meter/values, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"power": 3500.5,
			"energy_rel": 12345.67,
			"energy_abs": 87654.32,
			"phases_active": [true, true, false],
			"phases_connected": [true, true, true]
		}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	meterService := NewMeterService(request)
	result, err := meterService.Values()
	if err != nil {
		t.Errorf("Meter.Values() error = %v", err)
	}

	// Check power value
	if result["power"] != 3500.5 {
		t.Errorf("Meter.Values() power = %v, expected 3500.5", result["power"])
	}

	// Check energy_rel value
	if result["energy_rel"] != 12345.67 {
		t.Errorf("Meter.Values() energy_rel = %v, expected 12345.67", result["energy_rel"])
	}

	// Check energy_abs value
	if result["energy_abs"] != 87654.32 {
		t.Errorf("Meter.Values() energy_abs = %v, expected 87654.32", result["energy_abs"])
	}

	// Check phases_active array
	phasesActive, ok := result["phases_active"].([]interface{})
	if !ok {
		t.Error("Meter.Values() phases_active is not an array")
		return
	}
	
	if len(phasesActive) != 3 {
		t.Errorf("Meter.Values() phases_active length = %d, expected 3", len(phasesActive))
	}

	if phasesActive[0] != true {
		t.Errorf("Meter.Values() phases_active[0] = %v, expected true", phasesActive[0])
	}
	if phasesActive[1] != true {
		t.Errorf("Meter.Values() phases_active[1] = %v, expected true", phasesActive[1])
	}
	if phasesActive[2] != false {
		t.Errorf("Meter.Values() phases_active[2] = %v, expected false", phasesActive[2])
	}

	// Check phases_connected array
	phasesConnected, ok := result["phases_connected"].([]interface{})
	if !ok {
		t.Error("Meter.Values() phases_connected is not an array")
		return
	}
	
	if len(phasesConnected) != 3 {
		t.Errorf("Meter.Values() phases_connected length = %d, expected 3", len(phasesConnected))
	}

	for i, phase := range phasesConnected {
		if phase != true {
			t.Errorf("Meter.Values() phases_connected[%d] = %v, expected true", i, phase)
		}
	}
}

func TestMeter_Values_EmptyResponse(t *testing.T) {
	// Create a test server that returns empty JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	meterService := NewMeterService(request)
	result, err := meterService.Values()
	if err != nil {
		t.Errorf("Meter.Values() error = %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Meter.Values() returned %d values, expected 0", len(result))
	}
}

func TestMeter_Values_ErrorResponse(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	meterService := NewMeterService(request)
	_, err := meterService.Values()
	if err == nil {
		t.Error("Meter.Values() expected error for 500 response")
	}
}
