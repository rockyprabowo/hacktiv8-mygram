package comment_payloads

import "github.com/jellydator/validation"

type CommentInsertPayload struct {
	UserID  int    `json:"user_id" fake:"skip"`
	PhotoID int    `json:"photo_id" fake:"skip"`
	Message string `json:"message" fake:"{sentence:13}"`
}

func (p *CommentInsertPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.PhotoID, validation.Required),
		validation.Field(&p.Message, validation.Required),
	)
}
