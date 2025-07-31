package info

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewInfoService(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	if infoService == nil {
		t.Error("NewInfoService() returned nil")
	}
	if infoService.request != request {
		t.Error("NewInfoService() request not set correctly")
	}
}

func TestInfo_DisplayName(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/info/display_name" {
			t.Errorf("Expected path /info/display_name, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"display_name": "WARP Charger"}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	result, err := infoService.DisplayName()
	if err != nil {
		t.Errorf("Info.DisplayName() error = %v", err)
	}

	if result["display_name"] != "WARP Charger" {
		t.Errorf("Info.DisplayName() = %v, expected display_name: 'WARP Charger'", result)
	}
}

func TestInfo_Features(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/info/features" {
			t.Errorf("Expected path /info/features, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"features": ["nfc", "meter", "evse"]}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	result, err := infoService.Features()
	if err != nil {
		t.Errorf("Info.Features() error = %v", err)
	}

	expected := `{"features": ["nfc", "meter", "evse"]}`
	if string(result) != expected {
		t.Errorf("Info.Features() = %s, expected %s", string(result), expected)
	}
}

func TestInfo_Modules(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/info/modules" {
			t.Errorf("Expected path /info/modules, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"modules": {"evse": "1.0", "nfc": "2.1"}}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	result, err := infoService.Modules()
	if err != nil {
		t.Errorf("Info.Modules() error = %v", err)
	}

	modules, ok := result["modules"].(map[string]interface{})
	if !ok {
		t.Error("Info.Modules() did not return expected modules structure")
		return
	}

	if modules["evse"] != "1.0" {
		t.Errorf("Info.Modules() evse version = %v, expected '1.0'", modules["evse"])
	}
}

func TestInfo_Name(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/info/name" {
			t.Errorf("Expected path /info/name, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "warp-ABC123"}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	result, err := infoService.Name()
	if err != nil {
		t.Errorf("Info.Name() error = %v", err)
	}

	if result["name"] != "warp-ABC123" {
		t.Errorf("Info.Name() = %v, expected name: 'warp-ABC123'", result)
	}
}

func TestInfo_Version(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/info/version" {
			t.Errorf("Expected path /info/version, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"firmware": "2.1.5", "spiffs": "1.0.0"}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	infoService := NewInfoService(request)
	result, err := infoService.Version()
	if err != nil {
		t.Errorf("Info.Version() error = %v", err)
	}

	if result["firmware"] != "2.1.5" {
		t.Errorf("Info.Version() firmware = %v, expected '2.1.5'", result["firmware"])
	}

	if result["spiffs"] != "1.0.0" {
		t.Errorf("Info.Version() spiffs = %v, expected '1.0.0'", result["spiffs"])
	}
}
