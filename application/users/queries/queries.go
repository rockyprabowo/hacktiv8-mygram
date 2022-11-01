package user_queries

import (
	"context"
	"rocky.my.id/git/mygram/application/common/authentication/dto"
	contracts "rocky.my.id/git/mygram/application/users/contracts"
	dto "rocky.my.id/git/mygram/application/users/dto"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
)

type UserQueries struct {
	Repository contracts.UserRepositoryContracts
}

func NewUserQueries(repository contracts.UserRepositoryContracts) *UserQueries {
	return &UserQueries{Repository: repository}
}

func (q UserQueries) GetUserProfile(ctx context.Context, payload payloads.UserGetPayload) (*dto.UserDTO, error) {
	var userVM dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	userVM = dto.MapFromEntity(*data)

	return &userVM, nil
}

func (q UserQueries) AuthenticateUser(ctx context.Context, payload payloads.UserLoginPayload) (*dto.UserDTO, *authentication_dto.AuthTokenDTO, error) {
	var (
		userVM      dto.UserDTO
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

	authPayload := payloads.AuthTokenPayload{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}

	authToken, err = q.Repository.CreateAuthToken(authPayload)
	if err != nil {
		return nil, nil, err
	}

	userVM = dto.MapFromEntity(*data)
	authTokenVM.Token = authToken
	return &userVM, &authTokenVM, nil
}
