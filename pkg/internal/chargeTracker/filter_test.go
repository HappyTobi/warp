package chargeTracker

import (
	"testing"
	"time"
)

func TestNewFilter(t *testing.T) {
	tests := []struct {
		name        string
		filterName  string
		filterValue string
		filterNumber int
		expectType  string
	}{
		{
			name:        "month filter",
			filterName:  "month",
			filterValue: "",
			filterNumber: 5,
			expectType:  "*chargeTracker.MonthFilter",
		},
		{
			name:        "year filter", 
			filterName:  "year",
			filterValue: "",
			filterNumber: 2023,
			expectType:  "*chargeTracker.YearFilter",
		},
		{
			name:        "user filter",
			filterName:  "user",
			filterValue: "testuser",
			filterNumber: 0,
			expectType:  "*chargeTracker.UserFilter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewFilter(tt.filterName, tt.filterValue, tt.filterNumber)
			if filter == nil {
				t.Errorf("NewFilter() returned nil for %s", tt.filterName)
			}
		})
	}
}

func TestFilters(t *testing.T) {
	monthFilter := NewFilter("month", "", 5)
	yearFilter := NewFilter("year", "", 2023)
	userFilter := NewFilter("user", "testuser", 0)

	filters := Filters(monthFilter, yearFilter, userFilter)

	if len(filters) != 3 {
		t.Errorf("Filters() returned %d filters, expected 3", len(filters))
	}

	if filters[0] != monthFilter {
		t.Error("Filters() first filter not correct")
	}
	if filters[1] != yearFilter {
		t.Error("Filters() second filter not correct")
	}
	if filters[2] != userFilter {
		t.Error("Filters() third filter not correct")
	}
}

func TestMonthFilter_Filter(t *testing.T) {
	filter := &MonthFilter{filterValue: 5} // May

	tests := []struct {
		name     string
		charge   *Charge
		expected bool
	}{
		{
			name: "matching month",
			charge: &Charge{
				Time: time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC),
			},
			expected: true,
		},
		{
			name: "non-matching month",
			charge: &Charge{
				Time: time.Date(2023, 6, 15, 10, 30, 0, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "different year same month",
			charge: &Charge{
				Time: time.Date(2024, 5, 15, 10, 30, 0, 0, time.UTC),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.Filter(tt.charge)
			if result != tt.expected {
				t.Errorf("MonthFilter.Filter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestYearFilter_Filter(t *testing.T) {
	filter := &YearFilter{filterValue: 2023}

	tests := []struct {
		name     string
		charge   *Charge
		expected bool
	}{
		{
			name: "matching year",
			charge: &Charge{
				Time: time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC),
			},
			expected: true,
		},
		{
			name: "non-matching year",
			charge: &Charge{
				Time: time.Date(2024, 5, 15, 10, 30, 0, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "different month same year",
			charge: &Charge{
				Time: time.Date(2023, 12, 31, 23, 59, 0, 0, time.UTC),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.Filter(tt.charge)
			if result != tt.expected {
				t.Errorf("YearFilter.Filter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestUserFilter_Filter(t *testing.T) {
	filter := &UserFilter{filterValue: "testuser"}

	tests := []struct {
		name     string
		charge   *Charge
		expected bool
	}{
		{
			name: "matching user",
			charge: &Charge{
				User: "testuser",
			},
			expected: true,
		},
		{
			name: "non-matching user",
			charge: &Charge{
				User: "otheruser",
			},
			expected: false,
		},
		{
			name: "case sensitive match",
			charge: &Charge{
				User: "TestUser",
			},
			expected: false,
		},
		{
			name: "empty user",
			charge: &Charge{
				User: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.Filter(tt.charge)
			if result != tt.expected {
				t.Errorf("UserFilter.Filter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
