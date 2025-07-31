package warp

import (
	"net/http"
	"strings"
	"testing"
)

func TestDigestParts(t *testing.T) {
	tests := []struct {
		name           string
		wwwAuth        string
		expectedNonce  string
		expectedRealm  string
		expectedQop    string
	}{
		{
			name:           "valid digest auth header",
			wwwAuth:        `Digest realm="test", nonce="abc123", qop="auth"`,
			expectedNonce:  "abc123",
			expectedRealm:  "test",
			expectedQop:    "auth",
		},
		{
			name:           "different order",
			wwwAuth:        `Digest qop="auth", realm="warp", nonce="xyz789"`,
			expectedNonce:  "xyz789",
			expectedRealm:  "warp",
			expectedQop:    "auth",
		},
		{
			name:           "empty header",
			wwwAuth:        "",
			expectedNonce:  "",
			expectedRealm:  "",
			expectedQop:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Header: make(http.Header),
			}
			if tt.wwwAuth != "" {
				resp.Header.Set("Www-Authenticate", tt.wwwAuth)
			}

			result := digestParts(resp)

			if result["nonce"] != tt.expectedNonce {
				t.Errorf("digestParts() nonce = %s, expected %s", result["nonce"], tt.expectedNonce)
			}
			if result["realm"] != tt.expectedRealm {
				t.Errorf("digestParts() realm = %s, expected %s", result["realm"], tt.expectedRealm)
			}
			if result["qop"] != tt.expectedQop {
				t.Errorf("digestParts() qop = %s, expected %s", result["qop"], tt.expectedQop)
			}
		})
	}
}

func TestGetMD5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name:     "test credentials",
			input:    "user:realm:password",
			expected: "ebbc0ff9a121dbb6789bbe5f82174fa0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMD5(tt.input)
			if result != tt.expected {
				t.Errorf("getMD5(%s) = %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetCnonce(t *testing.T) {
	// Test that cnonce is generated
	cnonce1 := getCnonce()
	cnonce2 := getCnonce()

	// Should be 16 characters long
	if len(cnonce1) != 16 {
		t.Errorf("getCnonce() length = %d, expected 16", len(cnonce1))
	}

	// Should be different each time (very high probability)
	if cnonce1 == cnonce2 {
		t.Errorf("getCnonce() generated same value twice: %s", cnonce1)
	}

	// Should only contain hex characters
	validChars := "0123456789abcdef"
	for _, char := range cnonce1 {
		if !strings.ContainsRune(validChars, char) {
			t.Errorf("getCnonce() contains invalid character: %c", char)
		}
	}
}

func TestDigestAuthorization(t *testing.T) {
	// Create a mock response with digest auth header
	resp := &http.Response{
		Header: make(http.Header),
	}
	resp.Header.Set("Www-Authenticate", `Digest realm="test", nonce="abc123", qop="auth"`)

	auth := digestAuthrization("http://test.com/api", "GET", "testuser", "testpass", resp)

	// Check that the authorization string contains expected components
	expectedComponents := []string{
		`username="testuser"`,
		`realm="test"`,
		`nonce="abc123"`,
		`uri="http://test.com/api"`,
		`qop="auth"`,
		"response=",
	}

	for _, component := range expectedComponents {
		if !strings.Contains(auth, component) {
			t.Errorf("digestAuthrization() missing component: %s", component)
		}
	}

	// Check that it starts with "Digest "
	if !strings.HasPrefix(auth, "Digest ") {
		t.Errorf("digestAuthrization() should start with 'Digest ', got: %s", auth)
	}
}
