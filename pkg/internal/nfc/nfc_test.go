package nfc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewNfcTagsService(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	nfcService := NewNfcTagsService(request)
	if nfcService == nil {
		t.Error("NewNfcTagsService() returned nil")
	}
	if nfcService.request != request {
		t.Error("NewNfcTagsService() request not set correctly")
	}
}

func TestNfc_SeenTags(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/nfc/seen_tags" {
			t.Errorf("Expected path /nfc/seen_tags, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"tag_type": 2, "tag_id": "ABC123", "last_seen": 1640995200},
			{"tag_type": 2, "tag_id": "DEF456", "last_seen": 1640995300},
			{"tag_type": 2, "tag_id": "", "last_seen": 1640995400}
		]`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	nfcService := NewNfcTagsService(request)
	tags, err := nfcService.SeenTags()
	if err != nil {
		t.Errorf("Nfc.SeenTags() error = %v", err)
	}

	// Should filter out the tag with empty ID
	if len(tags) != 2 {
		t.Errorf("Nfc.SeenTags() returned %d tags, expected 2 (empty ID should be filtered)", len(tags))
	}

	// Check first tag
	if tags[0].Type != 2 {
		t.Errorf("First tag Type = %d, expected 2", tags[0].Type)
	}
	if tags[0].Id != "ABC123" {
		t.Errorf("First tag Id = %s, expected 'ABC123'", tags[0].Id)
	}
	if tags[0].LastSeen != 1640995200 {
		t.Errorf("First tag LastSeen = %d, expected 1640995200", tags[0].LastSeen)
	}

	// Check second tag
	if tags[1].Type != 2 {
		t.Errorf("Second tag Type = %d, expected 2", tags[1].Type)
	}
	if tags[1].Id != "DEF456" {
		t.Errorf("Second tag Id = %s, expected 'DEF456'", tags[1].Id)
	}
}

func TestNfc_Config(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/nfc/config" {
			t.Errorf("Expected path /nfc/config, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"authorized_tags": [
				{"user_id": 1, "tag_type": 2, "tag_id": "ABC123"},
				{"user_id": 2, "tag_type": 2, "tag_id": "DEF456"}
			]
		}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	nfcService := NewNfcTagsService(request)
	config, err := nfcService.Config()
	if err != nil {
		t.Errorf("Nfc.Config() error = %v", err)
	}

	if len(config.AuthorizedTags) != 2 {
		t.Errorf("Nfc.Config() returned %d authorized tags, expected 2", len(config.AuthorizedTags))
	}

	// Check first authorized tag
	firstTag := config.AuthorizedTags[0]
	if firstTag.UserId != 1 {
		t.Errorf("First authorized tag UserId = %d, expected 1", firstTag.UserId)
	}
	if firstTag.Type != 2 {
		t.Errorf("First authorized tag Type = %d, expected 2", firstTag.Type)
	}
	if firstTag.Id != "ABC123" {
		t.Errorf("First authorized tag Id = %s, expected 'ABC123'", firstTag.Id)
	}

	// Check second authorized tag
	secondTag := config.AuthorizedTags[1]
	if secondTag.UserId != 2 {
		t.Errorf("Second authorized tag UserId = %d, expected 2", secondTag.UserId)
	}
	if secondTag.Id != "DEF456" {
		t.Errorf("Second authorized tag Id = %s, expected 'DEF456'", secondTag.Id)
	}
}

func TestNfc_StartCharging(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/nfc/inject_tag_start" {
			t.Errorf("Expected path /nfc/inject_tag_start, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	nfcService := NewNfcTagsService(request)
	userTag := UserTag{
		UserId: 1,
		Type:   2,
		Id:     "ABC123",
	}

	err := nfcService.StartCharging(userTag)
	if err != nil {
		t.Errorf("Nfc.StartCharging() error = %v", err)
	}
}

func TestNfc_StopCharging(t *testing.T) {
	// Create a test server that handles both seen_tags and inject_tag_stop
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/nfc/seen_tags" {
			// Return mock seen tags
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"tag_type": 2, "tag_id": "ABC123", "last_seen": 1640995200}]`))
		} else if r.URL.Path == "/nfc/inject_tag_stop" {
			if r.Method != "PUT" {
				t.Errorf("Expected PUT request for inject_tag_stop, got %s", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true}`))
		} else {
			t.Errorf("Unexpected path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	request := warp.Request{
		Warp:        server.URL,
		ContentType: warp.JSON,
	}

	nfcService := NewNfcTagsService(request)
	err := nfcService.StopCharging()
	if err != nil {
		t.Errorf("Nfc.StopCharging() error = %v", err)
	}
}

func TestTagStruct(t *testing.T) {
	tag := Tag{
		Type:     2,
		Id:       "ABC123",
		LastSeen: 1640995200,
	}

	if tag.Type != 2 {
		t.Errorf("Tag.Type = %d, expected 2", tag.Type)
	}
	if tag.Id != "ABC123" {
		t.Errorf("Tag.Id = %s, expected 'ABC123'", tag.Id)
	}
	if tag.LastSeen != 1640995200 {
		t.Errorf("Tag.LastSeen = %d, expected 1640995200", tag.LastSeen)
	}
}

func TestUserTagStruct(t *testing.T) {
	userTag := UserTag{
		UserId: 42,
		Type:   2,
		Id:     "DEF456",
	}

	if userTag.UserId != 42 {
		t.Errorf("UserTag.UserId = %d, expected 42", userTag.UserId)
	}
	if userTag.Type != 2 {
		t.Errorf("UserTag.Type = %d, expected 2", userTag.Type)
	}
	if userTag.Id != "DEF456" {
		t.Errorf("UserTag.Id = %s, expected 'DEF456'", userTag.Id)
	}
}

func TestAuthorizedTagsStruct(t *testing.T) {
	userTag1 := UserTag{UserId: 1, Type: 2, Id: "ABC123"}
	userTag2 := UserTag{UserId: 2, Type: 2, Id: "DEF456"}
	
	authorizedTags := AuthorizedTags{
		AuthorizedTags: []UserTag{userTag1, userTag2},
	}

	if len(authorizedTags.AuthorizedTags) != 2 {
		t.Errorf("AuthorizedTags.AuthorizedTags length = %d, expected 2", len(authorizedTags.AuthorizedTags))
	}

	if authorizedTags.AuthorizedTags[0] != userTag1 {
		t.Error("AuthorizedTags.AuthorizedTags[0] not set correctly")
	}
	if authorizedTags.AuthorizedTags[1] != userTag2 {
		t.Error("AuthorizedTags.AuthorizedTags[1] not set correctly")
	}
}
