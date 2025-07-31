package update

import (
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewUpdate(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	warpVersion := "1.0.0"
	firmwareVersion := "2.1.5"

	updateService := NewUpdate(request, warpVersion, firmwareVersion)
	if updateService == nil {
		t.Error("NewUpdate() returned nil")
	}
	if updateService.request != request {
		t.Error("NewUpdate() request not set correctly")
	}
	if updateService.warpVersion != warpVersion {
		t.Error("NewUpdate() warpVersion not set correctly")
	}
	if updateService.firmwareVersion != firmwareVersion {
		t.Error("NewUpdate() firmwareVersion not set correctly")
	}
}

func TestUpdateStruct(t *testing.T) {
	update := &Update{
		Available:      true,
		UpdateVersion:  "1.1.0",
		CurrentVersion: "1.0.0",
	}

	if !update.Available {
		t.Errorf("Update.Available = %v, expected true", update.Available)
	}

	if update.UpdateVersion != "1.1.0" {
		t.Errorf("Update.UpdateVersion = %s, expected '1.1.0'", update.UpdateVersion)
	}

	if update.CurrentVersion != "1.0.0" {
		t.Errorf("Update.CurrentVersion = %s, expected '1.0.0'", update.CurrentVersion)
	}
}

func TestTagsStruct(t *testing.T) {
	tag := &Tags{
		Name:   "v1.0.0",
		ZipUrl: "https://example.com/archive.zip",
		TarUrl: "https://example.com/archive.tar.gz",
	}

	if tag.Name != "v1.0.0" {
		t.Errorf("Tags.Name = %s, expected 'v1.0.0'", tag.Name)
	}

	if tag.ZipUrl != "https://example.com/archive.zip" {
		t.Errorf("Tags.ZipUrl = %s, expected 'https://example.com/archive.zip'", tag.ZipUrl)
	}

	if tag.TarUrl != "https://example.com/archive.tar.gz" {
		t.Errorf("Tags.TarUrl = %s, expected 'https://example.com/archive.tar.gz'", tag.TarUrl)
	}
}

func TestInfoUpdateStruct(t *testing.T) {
	request := warp.Request{
		Warp: "http://test.com",
	}

	infoUpdate := &InfoUpdate{
		request:         request,
		warpVersion:     "1.0.0",
		firmwareVersion: "2.1.5",
	}

	if infoUpdate.request.Warp != "http://test.com" {
		t.Errorf("InfoUpdate.request.Warp = %s, expected 'http://test.com'", infoUpdate.request.Warp)
	}

	if infoUpdate.warpVersion != "1.0.0" {
		t.Errorf("InfoUpdate.warpVersion = %s, expected '1.0.0'", infoUpdate.warpVersion)
	}

	if infoUpdate.firmwareVersion != "2.1.5" {
		t.Errorf("InfoUpdate.firmwareVersion = %s, expected '2.1.5'", infoUpdate.firmwareVersion)
	}
}
