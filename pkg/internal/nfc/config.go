package nfc

func (n *Nfc) Config() (AuthorizedTags, error) {
	n.request.Path = "nfc/config"

	data, err := n.request.Get()
	if err != nil {
		return AuthorizedTags{}, err
	}

	return deserialize[AuthorizedTags](data)
}
