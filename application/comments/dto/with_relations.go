package comment_dto

import (
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type CommentWithRelationsDTO struct {
	CommentDTO
	User  UserExcerpt  `json:"user"`
	Photo PhotoExcerpt `json:"photo"`
}

type UserExcerpt struct {
	ID       int                    `json:"id"`
	Username value_objects.Username `json:"username"`
	Email    value_objects.Email    `json:"email"`
}

type PhotoExcerpt struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

func MapFromEntityWithRelations(comment entities.Comment) CommentWithRelationsDTO {
	return CommentWithRelationsDTO{
		CommentDTO: MapFromEntity(comment),
		User: UserExcerpt{
			ID:       comment.User.ID,
			Username: comment.User.Username,
			Email:    comment.User.Email,
		},
		Photo: PhotoExcerpt{
			ID:       comment.Photo.ID,
			UserID:   comment.Photo.UserID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoURL: comment.Photo.PhotoURL,
		},
	}
}
