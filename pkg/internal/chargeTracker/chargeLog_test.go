package chargeTracker

import (
	"testing"
	"time"

	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func TestNewChargeLog(t *testing.T) {
	request := &warp.Request{
		Warp: "http://test.com",
		Path: "test/path",
	}

	chargeLog := NewChargeLog(request)
	if chargeLog == nil {
		t.Error("NewChargeLog() returned nil")
	}
	if chargeLog.request != request {
		t.Error("NewChargeLog() request not set correctly")
	}
}

func TestMapUserNameToId(t *testing.T) {
	userMapping := map[int]string{
		1: "user1",
		2: "user2", 
		3: "admin",
	}

	tests := []struct {
		name     string
		userId   int
		expected string
	}{
		{
			name:     "existing user",
			userId:   1,
			expected: "user1",
		},
		{
			name:     "another existing user",
			userId:   3,
			expected: "admin",
		},
		{
			name:     "non-existing user",
			userId:   99,
			expected: "unknown user",
		},
		{
			name:     "zero user id",
			userId:   0,
			expected: "unknown user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapUserNameToId(tt.userId, userMapping)
			if result != tt.expected {
				t.Errorf("mapUserNameToId(%d) = %s, expected %s", tt.userId, result, tt.expected)
			}
		})
	}
}

func TestDeserialize_EmptyData(t *testing.T) {
	emptyData := []byte{}
	testUsers := []*users.User{
		{Id: 1, Username: "user1"},
		{Id: 2, Username: "user2"},
	}
	
	charges, err := deserialize(emptyData, testUsers, []Filter{})
	if err != nil {
		t.Errorf("deserialize() with empty data error = %v", err)
	}
	
	if charges == nil {
		t.Error("deserialize() returned nil charges")
	}
	
	if len(charges.Charges) != 0 {
		t.Errorf("deserialize() returned %d charges, expected 0", len(charges.Charges))
	}
}

func TestDeserialize_WithFilters(t *testing.T) {
	// Create test users
	testUsers := []*users.User{
		{Id: 1, Username: "user1"},
		{Id: 2, Username: "user2"},
	}

	// Create test filters
	userFilter := &UserFilter{filterValue: "user1"}
	filters := []Filter{userFilter}

	// Test with empty data (filters should not cause issues)
	emptyData := []byte{}
	charges, err := deserialize(emptyData, testUsers, filters)
	if err != nil {
		t.Errorf("deserialize() with filters error = %v", err)
	}
	
	if charges == nil {
		t.Error("deserialize() returned nil charges")
	}
}

// Test helper function to create test charge data
func createTestChargeData() *Charge {
	return &Charge{
		Time:            time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC),
		User:            "testuser",
		PowerMeterStart: 100.5,
		PowerMeterEnd:   120.7,
		Duration:        "02:30:00",
	}
}

func TestChargeStruct(t *testing.T) {
	charge := createTestChargeData()

	if charge.User != "testuser" {
		t.Errorf("Charge.User = %s, expected 'testuser'", charge.User)
	}

	if charge.PowerMeterStart != 100.5 {
		t.Errorf("Charge.PowerMeterStart = %f, expected 100.5", charge.PowerMeterStart)
	}

	if charge.PowerMeterEnd != 120.7 {
		t.Errorf("Charge.PowerMeterEnd = %f, expected 120.7", charge.PowerMeterEnd)
	}

	if charge.Duration != "02:30:00" {
		t.Errorf("Charge.Duration = %s, expected '02:30:00'", charge.Duration)
	}

	expectedTime := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)
	if !charge.Time.Equal(expectedTime) {
		t.Errorf("Charge.Time = %v, expected %v", charge.Time, expectedTime)
	}
}

func TestChargesStruct(t *testing.T) {
	charge1 := createTestChargeData()
	charge2 := &Charge{
		Time:            time.Date(2023, 6, 15, 14, 45, 0, 0, time.UTC),
		User:            "anotheruser",
		PowerMeterStart: 200.0,
		PowerMeterEnd:   250.0,
		Duration:        "03:15:00",
	}

	charges := &Charges{
		Charges: []*Charge{charge1, charge2},
	}

	if len(charges.Charges) != 2 {
		t.Errorf("Charges.Charges length = %d, expected 2", len(charges.Charges))
	}

	if charges.Charges[0] != charge1 {
		t.Error("Charges.Charges[0] not set correctly")
	}

	if charges.Charges[1] != charge2 {
		t.Error("Charges.Charges[1] not set correctly")
	}
}
