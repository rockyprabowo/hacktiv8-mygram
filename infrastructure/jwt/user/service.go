package jwt_user

import (
	"github.com/golang-jwt/jwt/v4"
	"rocky.my.id/git/mygram/domain/exceptions"
	"time"
)

type UserJWTService struct {
	SecretKey []byte
}

func NewUserJWTService(secretKey []byte) *UserJWTService {
	return &UserJWTService{SecretKey: secretKey}
}

func (s UserJWTService) GenerateUserToken(id int, username, email string) (string, error) {
	claims := &UserClaims{
		UserID:   id,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mygram",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}
	return encodedToken, err
}

func (s UserJWTService) ParseUserToken(tokenString string) (any, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, s.keyFunc)
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, exceptions.InvalidAuthToken
	}
	return parsedToken, nil
}

func (s UserJWTService) keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, exceptions.InvalidAuthToken
	}
	return s.SecretKey, nil
}
