package nfc

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewNfcTagsService(request warp.Request) *Nfc {
	return &Nfc{
		request: request,
	}
}
