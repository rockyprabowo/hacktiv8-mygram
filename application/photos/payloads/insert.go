package photo_payloads

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type PhotoInsertPayload struct {
	UserID   int    `json:"user_id" fake:"skip"`
	Title    string `json:"title" fake:"{sentence:3}"`
	Caption  string `json:"caption" fake:"{sentence:3}"`
	PhotoURL string `json:"photo_url" fake:"{url}"`
}

func (p *PhotoInsertPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Caption, validation.Required),
		validation.Field(&p.PhotoURL, validation.Required, is.URL),
	)
}
