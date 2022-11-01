package comment_payloads

import "github.com/jellydator/validation"

type CommentGetByIDPayload struct {
	UserID int `json:"user_id"`
	ID     int `json:"id" param:"id"`
}

func (p *CommentGetByIDPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
	)
}
