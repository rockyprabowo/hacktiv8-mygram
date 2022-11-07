package social_media_dto

import (
	"rocky.my.id/git/mygram/application/common/models"
	"rocky.my.id/git/mygram/domain/entities"
)

type SocialMediaDTO struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	SocialMediaURL string `json:"social_media_url,omitempty"`
	UserID         int    `json:"user_id,omitempty"`
	models.DateTime
}

func MapFromEntity(socialMedia entities.SocialMedia) SocialMediaDTO {
	return SocialMediaDTO{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		DateTime: models.DateTime{
			CreatedAt: &socialMedia.CreatedAt,
			UpdatedAt: &socialMedia.UpdatedAt,
		},
	}
}
