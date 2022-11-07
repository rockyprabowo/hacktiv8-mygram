package user_queries

import (
	"context"
	"rocky.my.id/git/mygram/application/users/dto"
	"rocky.my.id/git/mygram/application/users/payloads"
)

func (q UserQueries) GetUserProfile(ctx context.Context, payload user_payloads.UserGetPayload) (*user_dto.UserDTO, error) {
	var userVM user_dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	userVM = user_dto.MapFromEntity(*data)

	return &userVM, nil
}
