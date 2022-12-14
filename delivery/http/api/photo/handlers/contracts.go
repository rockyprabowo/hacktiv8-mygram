package photo_handlers

import "github.com/labstack/echo/v4"

type PhotoHTTPHandlerContract interface {
	PostPhoto(ctx echo.Context) error
	GetPhotos(ctx echo.Context) error
	GetPhotoByID(ctx echo.Context) error
	GetOwnedPhotos(ctx echo.Context) error
	UpdatePhoto(ctx echo.Context) error
	DeletePhoto(ctx echo.Context) error
}
