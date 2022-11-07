package social_media_dto

import (
	userDTO "rocky.my.id/git/mygram/application/users/dto"
	"rocky.my.id/git/mygram/domain/entities"
)

type SocialMediaWithUserDTO struct {
	SocialMediaDTO
	User userDTO.UserDTO `json:"user"`
}

func MapFromEntityWithUser(socialMedia entities.SocialMedia) SocialMediaWithUserDTO {
	return SocialMediaWithUserDTO{
		SocialMediaDTO: MapFromEntity(socialMedia),
		User:           userDTO.MapFromEntity(socialMedia.User),
	}
}
