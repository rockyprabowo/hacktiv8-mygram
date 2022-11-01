package comment_dto

import (
	"rocky.my.id/git/mygram/application/common/models"
	"rocky.my.id/git/mygram/domain/entities"
)

type CommentDTO struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message"`
	models.DateTime
}

func MapFromEntity(comment entities.Comment) CommentDTO {
	return CommentDTO{
		ID:      comment.ID,
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
		Message: comment.Message,
		DateTime: models.DateTime{
			CreatedAt: &comment.CreatedAt,
			UpdatedAt: &comment.UpdatedAt,
		},
	}
}
