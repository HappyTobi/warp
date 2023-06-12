package evse

import "encoding/json"

func (e *Evse) StopCharging() error {
	e.request.Path = "evse/stop_charging"

	stopChargePayload := "{}"
	payload, err := json.Marshal(&stopChargePayload)
	if err != nil {
		return err
	}

	_, err = e.request.Put(payload)
	return err
}
