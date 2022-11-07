package social_media_payloads

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type SocialMediaInsertPayload struct {
	UserID         int    `json:"user_id" fake:"skip"`
	Name           string `json:"name" fake:"{company}"`
	SocialMediaURL string `json:"social_media_url" fake:"{url}"`
}

func (p *SocialMediaInsertPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.SocialMediaURL, validation.Required, is.URL),
	)
}
