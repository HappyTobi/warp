package evse

func (e *Evse) State() (map[string]interface{}, error) {
	e.request.Path = "evse/state"
	return e.request.GetJson()
}
