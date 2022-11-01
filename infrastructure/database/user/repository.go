package user_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"rocky.my.id/git/mygram/application/photos/contracts"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/database/common/passwords"
)

type UserRepository struct {
	DB            *gorm.DB
	UserTokenizer photo_contracts.UserTokenAuthenticationContract
}

func NewUserRepository(DB *gorm.DB, userTokenizer photo_contracts.UserTokenAuthenticationContract) *UserRepository {
	return &UserRepository{DB: DB, UserTokenizer: userTokenizer}
}

func (r UserRepository) CheckIdentityUnique(ctx context.Context, email value_objects.Email, username value_objects.Username) error {
	var usernameRowCount, emailRowCount int64
	r.DB.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).Count(&emailRowCount)
	if emailRowCount > 0 {
		return exceptions.EmailAlreadyRegistered
	}

	r.DB.WithContext(ctx).Model(&entities.User{}).Where("username = ?", username).Count(&usernameRowCount)
	if usernameRowCount > 0 {
		return exceptions.UsernameAlreadyRegistered
	}

	return nil
}

func (r UserRepository) CreateAuthToken(payload payloads.AuthTokenPayload) (string, error) {
	token, err := r.UserTokenizer.GenerateUserToken(payload.ID, string(payload.Username), string(payload.Email))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r UserRepository) GetUser(ctx context.Context, payload payloads.UserGetPayload) (*entities.User, error) {
	var user *entities.User
	var keyQuery string

	switch payload.Key {
	case payloads.UserIDPayloadKey:
		keyQuery = "id = ?"
	case payloads.UserEmailPayloadKey:
		keyQuery = "email = ?"
	case payloads.UsernamePayloadKey:
		keyQuery = "username = ?"
	default:
		return nil, exceptions.UserNotFoundError
	}

	err := r.DB.WithContext(ctx).Where(keyQuery, payload.Value).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exceptions.UserNotFoundError
	}
	return user, err
}

func (r UserRepository) CreateUser(ctx context.Context, payload payloads.UserRegisterPayload) (*entities.User, error) {
	var user *entities.User

	if err := r.CheckIdentityUnique(ctx, payload.Email, payload.Username); err != nil {
		return nil, err
	}

	hashedPassword, err := passwords.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user = &entities.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
		Age:      payload.Age,
	}
	err = r.DB.WithContext(ctx).Save(user).Error
	return user, err
}

func (r UserRepository) BatchSave(ctx context.Context, payloads []entities.User) (int64, error) {
	results := r.DB.WithContext(ctx).CreateInBatches(payloads, 100)
	return results.RowsAffected, results.Error
}

func (r UserRepository) AuthenticateUser(ctx context.Context, payload payloads.UserLoginPayload) (*entities.User, error) {
	user, err := r.GetUser(ctx, payloads.UserGetPayload{
		Key:   "Email",
		Value: payload.Email,
	})
	if err != nil {
		if errors.Is(err, exceptions.UserNotFoundError) {
			return nil, exceptions.InvalidCredentials
		}
		return nil, err
	}

	passwordOk := passwords.ComparePasswordHash(user.Password, []byte(payload.Password))
	if !passwordOk {
		return nil, exceptions.InvalidCredentials
	}

	return user, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, payload payloads.UserProfileUpdatePayload) (*entities.User, error) {
	user, err := r.GetUser(ctx, payloads.UserGetPayload{
		Key:   "ID",
		Value: payload.ID,
	})

	if err != nil {
		return nil, err
	}

	user.Username = payload.Username
	user.Email = payload.Email

	err = r.DB.WithContext(ctx).Save(user).Error
	return user, err
}

func (r UserRepository) DeleteUser(ctx context.Context, payload payloads.UserDeletePayload) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.User{}, payload.ID)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
