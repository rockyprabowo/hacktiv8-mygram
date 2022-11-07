package user_queries

import (
	"context"
	"rocky.my.id/git/mygram/application/common/authentication/dto"
	"rocky.my.id/git/mygram/application/users/dto"
	"rocky.my.id/git/mygram/application/users/payloads"
)

func (q UserQueries) AuthenticateUser(ctx context.Context, payload user_payloads.UserLoginPayload) (*user_dto.UserDTO, *authentication_dto.AuthTokenDTO, error) {
	var (
		userVM      user_dto.UserDTO
		authTokenVM authentication_dto.AuthTokenDTO
		authToken   string
	)

	if err := payload.Validate(); err != nil {
		return nil, nil, err
	}

	data, err := q.Repository.AuthenticateUser(ctx, payload)
	if err != nil {
		return nil, nil, err
	}

	authPayload := user_payloads.AuthTokenPayload{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}

	authToken, err = q.Repository.CreateAuthToken(authPayload)
	if err != nil {
		return nil, nil, err
	}

	userVM = user_dto.MapFromEntity(*data)
	authTokenVM.Token = authToken
	return &userVM, &authTokenVM, nil
}
