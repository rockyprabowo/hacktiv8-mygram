package photo_dto

import (
	"rocky.my.id/git/mygram/application/common/models"
	"rocky.my.id/git/mygram/domain/entities"
)

type PhotoDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
	models.DateTime
}

func MapFromEntity(photo entities.Photo) PhotoDTO {
	return PhotoDTO{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
		DateTime: models.DateTime{
			CreatedAt: &photo.DateTime.CreatedAt,
			UpdatedAt: &photo.DateTime.UpdatedAt,
		},
	}
}
