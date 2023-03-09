package chargeTracker

import (
	"time"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

type Charge struct {
	Time            time.Time
	User            string
	PowerMeterStart float32
	PowerMeterEnd   float32
	Duration        string
}

type ChargeLog struct {
	request *warp.Request
	Charges []*Charge
}
