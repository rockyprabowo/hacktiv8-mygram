package photo_dto

import (
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type PhotoWithUserDTO struct {
	PhotoDTO
	User UserExcerpt `json:"user"`
}

type UserExcerpt struct {
	Username value_objects.Username `json:"username"`
	Email    value_objects.Email    `json:"email"`
}

func MapFromEntityWithUser(photo entities.Photo) PhotoWithUserDTO {
	return PhotoWithUserDTO{
		PhotoDTO: MapFromEntity(photo),
		User: UserExcerpt{
			Username: photo.User.Username,
			Email:    photo.User.Email,
		},
	}
}
