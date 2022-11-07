package social_media_payloads

import (
	"github.com/jellydator/validation"
)

type SocialMediaGetAllByOwnerPayload struct {
	UserID int `json:"user_id"`
}

func (p *SocialMediaGetAllByOwnerPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
	)
}
