package evse

import "encoding/json"

func (e *Evse) ReadExternalCurrent() (int, error) {
	e.request.Path = "evse/external_current"

	data, err := e.request.Get()
	if err != nil {
		return -1, err
	}

	var externalCurrent ExternelCurrent
	if err := json.Unmarshal(data, &externalCurrent); err != nil {
		return -1, err
	}

	return externalCurrent.Current, nil
}

func (e *Evse) SetExternalCurrent(val int) error {
	e.request.Path = "evse/external_current"

	externelCurrent := ExternelCurrent{Current: val}

	payload, err := json.Marshal(&externelCurrent)
	if err != nil {
		return err
	}

	if _, err := e.request.Post(payload); err != nil {
		return err
	}

	return nil
}
