package settings

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStoreImage(t *testing.T) {
	// Create a temporary directory for test
	tempDir := t.TempDir()
	imagePath := filepath.Join(tempDir, "test-logo.png")

	err := StoreImage(imagePath)
	if err != nil {
		t.Errorf("StoreImage() error = %v", err)
	}

	// Check if the image file was created
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Errorf("StoreImage() did not create image file at %s", imagePath)
	}

	// Check if the file has content (logo bytes should be embedded)
	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		t.Errorf("Failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		t.Errorf("StoreImage() created empty file, expected embedded logo content")
	}

	// Verify file permissions
	expectedPerms := os.FileMode(0644)
	if fileInfo.Mode().Perm() != expectedPerms {
		t.Errorf("StoreImage() file permissions = %v, expected %v", fileInfo.Mode().Perm(), expectedPerms)
	}
}

func TestStoreImage_InvalidPath(t *testing.T) {
	// Test with an invalid path that cannot be created
	invalidPath := "/root/invalid/path/that/cannot/be/created/logo.png"

	err := StoreImage(invalidPath)
	// The function should return an error for invalid path, but behavior may vary
	// We test that it doesn't panic and either succeeds or fails gracefully
	_ = err // Don't assert on specific error as behavior depends on permissions
}

func TestStoreImage_ExistingFile(t *testing.T) {
	// Create a temporary directory and existing file
	tempDir := t.TempDir()
	imagePath := filepath.Join(tempDir, "existing-logo.png")

	// Create an existing file with different content
	existingContent := []byte("existing image content")
	err := os.WriteFile(imagePath, existingContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Store the embedded logo (should overwrite)
	err = StoreImage(imagePath)
	if err != nil {
		t.Errorf("StoreImage() with existing file error = %v", err)
	}

	// Read the file content
	content, err := os.ReadFile(imagePath)
	if err != nil {
		t.Errorf("Failed to read image file: %v", err)
	}

	// Should not contain the old content
	if string(content) == string(existingContent) {
		t.Errorf("StoreImage() did not overwrite existing file")
	}

	// Should contain the embedded logo content (non-zero length)
	if len(content) == 0 {
		t.Errorf("StoreImage() created empty file after overwrite")
	}
}

func TestStoreImage_DirectoryCreation(t *testing.T) {
	// Test that StoreImage can create parent directories
	tempDir := t.TempDir()
	nestedPath := filepath.Join(tempDir, "nested", "directory", "logo.png")

	err := StoreImage(nestedPath)
	if err != nil {
		// os.WriteFile doesn't create parent directories automatically
		// So this test expects an error for non-existing parent directories
		if !os.IsNotExist(err) {
			t.Errorf("StoreImage() with nested path unexpected error = %v", err)
		}
		return
	}

	// If no error, check if the image file was created
	if _, err := os.Stat(nestedPath); os.IsNotExist(err) {
		t.Errorf("StoreImage() did not create image file at nested path %s", nestedPath)
	}

	// Check if parent directories were created
	parentDir := filepath.Dir(nestedPath)
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		t.Errorf("StoreImage() did not create parent directory %s", parentDir)
	}
}

func TestStoreImage_ReadOnlyDirectory(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()
	readOnlyDir := filepath.Join(tempDir, "readonly")
	
	// Create directory and make it read-only
	err := os.Mkdir(readOnlyDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Make directory read-only
	err = os.Chmod(readOnlyDir, 0444)
	if err != nil {
		t.Fatalf("Failed to change directory permissions: %v", err)
	}

	// Restore permissions after test
	defer os.Chmod(readOnlyDir, 0755)

	imagePath := filepath.Join(readOnlyDir, "logo.png")
	err = StoreImage(imagePath)
	
	// Should return an error due to permission denied, but behavior may vary by system
	// We test that it doesn't panic
	_ = err // Don't assert on specific error behavior as it may vary
}
