package nfc

import (
	"encoding/json"
)

func (n *Nfc) SeenTags() ([]Tag, error) {
	n.request.Path = "nfc/seen_tags"

	data, err := n.request.Get()
	if err != nil {
		return nil, err
	}

	tags, err := deserialize[[]Tag](data)
	if err != nil {
		return nil, err
	}

	cleanTags := make([]Tag, 0)
	for _, tag := range tags {
		if len(tag.Id) != 0 {
			cleanTags = append(cleanTags, tag)
		}
	}

	return cleanTags, nil
}

func (n *Nfc) StartCharging(tag UserTag) error {
	n.request.Path = "nfc/inject_tag_start"

	chargeTag := &userTagCharge{Id: tag.Id, TagType: tag.Type}

	payload, err := json.Marshal(chargeTag)
	if err != nil {
		return err
	}

	_, err = n.request.Put(payload)
	return err
}

func (n *Nfc) StopCharging() error {
	tags, err := n.SeenTags()
	if err != nil || len(tags) == 0 {
		return err
	}

	n.request.Path = "nfc/inject_tag_stop"
	//always take the first tag
	chargeTag := &userTagCharge{Id: tags[0].Id, TagType: tags[0].Type}
	payload, err := json.Marshal(&chargeTag)
	if err != nil {
		return err
	}

	_, err = n.request.Put(payload)
	return err
}

func deserialize[T []Tag | AuthorizedTags](data []byte) (T, error) {
	var nfcTags T
	if err := json.Unmarshal(data, &nfcTags); err != nil {
		return nfcTags, err
	}

	return nfcTags, nil
}
