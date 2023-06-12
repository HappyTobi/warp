package evse

import "encoding/json"

func (e *Evse) CurrentChargePower() (ChargePower, error) {
	e.request.Path = "evse/global_current"

	data, err := e.request.Get()
	if err != nil {
		return ChargePower{}, err
	}

	var currentChargePower ChargePower
	if err := json.Unmarshal(data, &currentChargePower); err != nil {
		return ChargePower{}, err
	}

	return currentChargePower, nil
}

func (e *Evse) UpdateChargePower(chargePower int) error {
	e.request.Path = "evse/global_current_update"

	charge := &ChargePower{Current: chargePower}

	payload, err := json.Marshal(charge)
	if err != nil {
		return err
	}

	if _, err := e.request.Put(payload); err != nil {
		return err
	}

	return nil
}
