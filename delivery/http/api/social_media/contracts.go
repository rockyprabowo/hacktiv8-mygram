package social_media_http_delivery

import "github.com/labstack/echo/v4"

type SocialMediaHTTPHandlerContract interface {
	CreateUserSocialMedia(ctx echo.Context) error
	GetUserSocialMedias(ctx echo.Context) error
	UpdateUserSocialMedia(ctx echo.Context) error
	DeleteUserSocialMedia(ctx echo.Context) error
}
