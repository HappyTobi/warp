package evse

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewEvseService(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	if evseService == nil {
		t.Error("NewEvseService() returned nil")
	}
	if evseService.request != request {
		t.Error("NewEvseService() request not set correctly")
	}
}

func TestEvse_State(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/evse/state" {
			t.Errorf("Expected path /evse/state, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"iec61851_state": 2, "charger_state": 1}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	result, err := evseService.State()
	if err != nil {
		t.Errorf("Evse.State() error = %v", err)
	}

	if result["iec61851_state"] != float64(2) {
		t.Errorf("Evse.State() iec61851_state = %v, expected 2", result["iec61851_state"])
	}

	if result["charger_state"] != float64(1) {
		t.Errorf("Evse.State() charger_state = %v, expected 1", result["charger_state"])
	}
}

func TestEvse_CurrentChargePower(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/evse/global_current" {
			t.Errorf("Expected path /evse/global_current, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"current": 16}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	result, err := evseService.CurrentChargePower()
	if err != nil {
		t.Errorf("Evse.CurrentChargePower() error = %v", err)
	}

	if result.Current != 16 {
		t.Errorf("Evse.CurrentChargePower() current = %d, expected 16", result.Current)
	}
}

func TestEvse_UpdateChargePower(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/evse/global_current_update" {
			t.Errorf("Expected path /evse/global_current_update, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	err := evseService.UpdateChargePower(20)
	if err != nil {
		t.Errorf("Evse.UpdateChargePower() error = %v", err)
	}
}

func TestEvse_ReadExternalCurrent(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/evse/external_current" {
			t.Errorf("Expected path /evse/external_current, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"current": 12}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	current, err := evseService.ReadExternalCurrent()
	if err != nil {
		t.Errorf("Evse.ReadExternalCurrent() error = %v", err)
	}

	if current != 12 {
		t.Errorf("Evse.ReadExternalCurrent() = %d, expected 12", current)
	}
}

func TestEvse_SetExternalCurrent(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/evse/external_current" {
			t.Errorf("Expected path /evse/external_current, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	err := evseService.SetExternalCurrent(15)
	if err != nil {
		t.Errorf("Evse.SetExternalCurrent() error = %v", err)
	}
}

func TestEvse_GetExternalClearOnDisconnect(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/evse/external_clear_on_disconnect" {
			t.Errorf("Expected path /evse/external_clear_on_disconnect, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"clear_on_disconnect": true}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	clearOnDisconnect, err := evseService.GetExternalClearOnDisconnect()
	if err != nil {
		t.Errorf("Evse.GetExternalClearOnDisconnect() error = %v", err)
	}

	if !clearOnDisconnect {
		t.Errorf("Evse.GetExternalClearOnDisconnect() = %v, expected true", clearOnDisconnect)
	}
}

func TestEvse_EnableClearOnDisconnect(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/evse/external_clear_on_disconnect" {
			t.Errorf("Expected path /evse/external_clear_on_disconnect, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	evseService := NewEvseService(request)
	err := evseService.EnableClearOnDisconnect()
	if err != nil {
		t.Errorf("Evse.EnableClearOnDisconnect() error = %v", err)
	}
}

func TestChargePowerStruct(t *testing.T) {
	chargePower := ChargePower{Current: 16}
	if chargePower.Current != 16 {
		t.Errorf("ChargePower.Current = %d, expected 16", chargePower.Current)
	}
}

func TestExternalCurrentStruct(t *testing.T) {
	externalCurrent := ExternelCurrent{Current: 20}
	if externalCurrent.Current != 20 {
		t.Errorf("ExternelCurrent.Current = %d, expected 20", externalCurrent.Current)
	}
}

func TestClearOnDisconnectStruct(t *testing.T) {
	clearOnDisconnect := ClearOnDisconnect{Enabled: true}
	if !clearOnDisconnect.Enabled {
		t.Errorf("ClearOnDisconnect.Enabled = %v, expected true", clearOnDisconnect.Enabled)
	}
}
