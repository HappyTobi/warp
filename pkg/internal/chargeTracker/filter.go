package chargeTracker

import (
	"time"
)

func Filters(filter ...Filter) []Filter {
	filters := make([]Filter, 0)
	for _, f := range filter {
		if f != nil {
			filters = append(filters, f)
		}
	}
	return filters
}

func NewFilter(filterName string, filterValue string, filterNumber int) Filter {
	switch filterName {
	case "month":
		if filterNumber >= 1 && filterNumber <= 12 {
			return &MonthFilter{filterValue: filterNumber}
		}
	case "year":
		if filterNumber >= 2000 {
			return &YearFilter{filterValue: filterNumber}
		}
	case "user":
		if len(filterValue) > 0 {
			return &UserFilter{filterValue: filterValue}
		}
	}

	return nil
}

func (mf *MonthFilter) Filter(charge *Charge) bool {
	return charge.Time.Month() == time.Month(mf.filterValue)
}

func (yF *YearFilter) Filter(charge *Charge) bool {
	return charge.Time.Year() == yF.filterValue
}

func (uF *UserFilter) Filter(charge *Charge) bool {
	return charge.User == uF.filterValue
}
