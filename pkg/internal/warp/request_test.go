package warp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest_Get(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/test/path" {
			t.Errorf("Expected path /test/path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	req := Request{
		Warp:        server.URL,
		Path:        "test/path",
		ContentType: JSON,
	}

	data, err := req.Get()
	if err != nil {
		t.Errorf("Request.Get() error = %v", err)
	}

	expected := `{"message": "success"}`
	if string(data) != expected {
		t.Errorf("Request.Get() = %s, expected %s", string(data), expected)
	}
}

func TestRequest_GetJson(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key": "value", "number": 42}`))
	}))
	defer server.Close()

	req := Request{
		Warp:        server.URL,
		Path:        "test/path",
		ContentType: JSON,
	}

	result, err := req.GetJson()
	if err != nil {
		t.Errorf("Request.GetJson() error = %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("Request.GetJson() key = %v, expected 'value'", result["key"])
	}

	// JSON numbers are unmarshaled as float64
	if result["number"] != float64(42) {
		t.Errorf("Request.GetJson() number = %v, expected 42", result["number"])
	}
}

func TestRequest_Put(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != string(JSON) {
			t.Errorf("Expected Content-Type %s, got %s", JSON, r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"updated": true}`))
	}))
	defer server.Close()

	req := Request{
		Warp:        server.URL,
		Path:        "test/path",
		ContentType: JSON,
	}

	testData := []byte(`{"test": "data"}`)
	data, err := req.Put(testData)
	if err != nil {
		t.Errorf("Request.Put() error = %v", err)
	}

	expected := `{"updated": true}`
	if string(data) != expected {
		t.Errorf("Request.Put() = %s, expected %s", string(data), expected)
	}
}

func TestRequest_Post(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"created": true}`))
	}))
	defer server.Close()

	req := Request{
		Warp:        server.URL,
		Path:        "test/path",
		ContentType: JSON,
	}

	testData := []byte(`{"test": "data"}`)
	data, err := req.Post(testData)
	if err != nil {
		t.Errorf("Request.Post() error = %v", err)
	}

	expected := `{"created": true}`
	if string(data) != expected {
		t.Errorf("Request.Post() = %s, expected %s", string(data), expected)
	}
}

func TestRequest_AuthenticationRequired(t *testing.T) {
	// Create a test server that requires authentication
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if Authorization header is present
		if r.Header.Get("Authorization") == "" {
			// Return 401 with digest challenge
			w.Header().Set("Www-Authenticate", `Digest realm="test", nonce="abc123", qop="auth"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		// If Authorization header is present, return success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"authenticated": true}`))
	}))
	defer server.Close()

	req := Request{
		Warp:        server.URL,
		Path:        "test/path",
		ContentType: JSON,
		Username:    "testuser",
		Password:    "testpass",
	}

	data, err := req.Get()
	if err != nil {
		t.Errorf("Request.Get() with auth error = %v", err)
	}

	expected := `{"authenticated": true}`
	if string(data) != expected {
		t.Errorf("Request.Get() with auth = %s, expected %s", string(data), expected)
	}
}

func TestRequestConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected interface{}
	}{
		{"GET method", GET, RequestMethod("GET")},
		{"POST method", POST, RequestMethod("POST")},
		{"PUT method", PUT, RequestMethod("PUT")},
		{"DELETE method", DELETE, RequestMethod("DELETE")},
		{"JSON content type", JSON, ContentType("application/json")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, expected %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestRequest_WithRenderer(t *testing.T) {
	mockRenderer := &mockRenderer{}
	req := Request{
		Warp:           "http://test.com",
		Path:           "test/path",
		ContentType:    JSON,
		OutputRenderer: mockRenderer,
	}

	// Just verify the renderer is set correctly
	if req.OutputRenderer != mockRenderer {
		t.Errorf("Request.OutputRenderer not set correctly")
	}
}

// Mock renderer for testing
type mockRenderer struct{}

func (m *mockRenderer) Render(data interface{}) (string, error) {
	return "mocked output", nil
}

func (m *mockRenderer) RenderBytes(data []byte) (string, error) {
	return "mocked bytes output", nil
}
