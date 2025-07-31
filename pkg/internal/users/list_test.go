package users

import (
	"testing"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewUsersService(t *testing.T) {
	request := warp.Request{
		Warp:        "http://test.com",
		Path:        "test/path",
		ContentType: warp.JSON,
	}

	userService := NewUsersService(request)
	if userService == nil {
		t.Error("NewUsersService() returned nil")
	}
	if userService.request != request {
		t.Error("NewUsersService() request not set correctly")
	}
}

func TestDeserialize_EmptyData(t *testing.T) {
	emptyData := []byte{}
	users, err := deserialize(emptyData)
	if err != nil {
		t.Errorf("deserialize() with empty data error = %v", err)
	}
	
	if users == nil {
		t.Error("deserialize() returned nil users")
	}
	
	if len(users) != 0 {
		t.Errorf("deserialize() returned %d users, expected 0", len(users))
	}
}

func TestDeserialize_SingleUser(t *testing.T) {
	// Create test data representing a single user
	// Username: "testuser" (8 bytes + padding), DisplayName: "Test User" (9 bytes + padding)
	testData := make([]byte, 64) // 2 * 32 bytes for username and display name
	
	copy(testData[0:8], "testuser")
	copy(testData[32:41], "Test User")
	
	users, err := deserialize(testData)
	if err != nil {
		t.Errorf("deserialize() error = %v", err)
	}
	
	if len(users) != 1 {
		t.Errorf("deserialize() returned %d users, expected 1", len(users))
	}
	
	user := users[0]
	if user.Username != "testuser" {
		t.Errorf("User.Username = %s, expected 'testuser'", user.Username)
	}
	
	if user.DisplayName != "Test User" {
		t.Errorf("User.DisplayName = %s, expected 'Test User'", user.DisplayName)
	}
	
	if user.Id != 0 {
		t.Errorf("User.Id = %d, expected 0", user.Id)
	}
}

func TestDeserialize_MultipleUsers(t *testing.T) {
	// Create test data representing multiple users
	testData := make([]byte, 128) // 4 * 32 bytes for 2 users
	
	// First user
	copy(testData[0:5], "user1")
	copy(testData[32:42], "First User")
	
	// Second user  
	copy(testData[64:69], "user2")
	copy(testData[96:107], "Second User")
	
	users, err := deserialize(testData)
	if err != nil {
		t.Errorf("deserialize() error = %v", err)
	}
	
	if len(users) != 2 {
		t.Errorf("deserialize() returned %d users, expected 2", len(users))
	}
	
	// Check first user
	if users[0].Username != "user1" {
		t.Errorf("First user.Username = %s, expected 'user1'", users[0].Username)
	}
	if users[0].DisplayName != "First User" {
		t.Errorf("First user.DisplayName = %s, expected 'First User'", users[0].DisplayName)
	}
	if users[0].Id != 0 {
		t.Errorf("First user.Id = %d, expected 0", users[0].Id)
	}
	
	// Check second user
	if users[1].Username != "user2" {
		t.Errorf("Second user.Username = %s, expected 'user2'", users[1].Username)
	}
	if users[1].DisplayName != "Second User" {
		t.Errorf("Second user.DisplayName = %s, expected 'Second User'", users[1].DisplayName)
	}
	if users[1].Id != 1 {
		t.Errorf("Second user.Id = %d, expected 1", users[1].Id)
	}
}

func TestDeserialize_EmptyUsernameIgnored(t *testing.T) {
	// Create test data with empty username (should be ignored)
	testData := make([]byte, 64)
	
	// Empty username, but with display name
	copy(testData[32:42], "Display Name")
	
	users, err := deserialize(testData)
	if err != nil {
		t.Errorf("deserialize() error = %v", err)
	}
	
	// Should be filtered out because username is empty
	if len(users) != 0 {
		t.Errorf("deserialize() returned %d users, expected 0 (empty username should be filtered)", len(users))
	}
}

func TestDeserialize_EmptyDisplayNameIgnored(t *testing.T) {
	// Create test data with empty display name (should be ignored)
	testData := make([]byte, 64)
	
	// Username but empty display name
	copy(testData[0:8], "testuser")
	
	users, err := deserialize(testData)
	if err != nil {
		t.Errorf("deserialize() error = %v", err)
	}
	
	// Should be filtered out because display name is empty
	if len(users) != 0 {
		t.Errorf("deserialize() returned %d users, expected 0 (empty display name should be filtered)", len(users))
	}
}

func TestUserStruct(t *testing.T) {
	user := &User{
		Id:          42,
		Username:    "testuser", 
		DisplayName: "Test User",
	}
	
	if user.Id != 42 {
		t.Errorf("User.Id = %d, expected 42", user.Id)
	}
	
	if user.Username != "testuser" {
		t.Errorf("User.Username = %s, expected 'testuser'", user.Username)
	}
	
	if user.DisplayName != "Test User" {
		t.Errorf("User.DisplayName = %s, expected 'Test User'", user.DisplayName)
	}
}
