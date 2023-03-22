package chargeTracker

import (
	"time"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

type Charge struct {
	Time            time.Time `json:"time" yaml:"time"`
	User            string    `json:"user" yaml:"user"`
	PowerMeterStart float32   `json:"powerMeterStart" yaml:"powerMeterStart"`
	PowerMeterEnd   float32   `json:"powerMeterEnd" yaml:"powerMeterEnd"`
	Duration        string    `json:"duration" yaml:"duration"`
}

type Charges struct {
	Charges []*Charge `json:"charges" yaml:"charges"`
}

type ChargeLog struct {
	request *warp.Request
}

type MonthFilter struct {
	filterValue int
}

type YearFilter struct {
	filterValue int
}

type UserFilter struct {
	filterValue string
}

type Filter interface {
	Filter(*Charge) bool
}
