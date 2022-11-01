package user_dto

import (
	"rocky.my.id/git/mygram/application/common/models"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type UserDTO struct {
	ID       int                    `json:"id"`
	Username value_objects.Username `json:"username"`
	Email    value_objects.Email    `json:"email"`
	Age      int                    `json:"age"`
	models.DateTime
}

func MapFromEntity(user entities.User) UserDTO {
	return UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		DateTime: models.DateTime{
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		},
	}
}
