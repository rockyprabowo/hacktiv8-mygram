package comment_http_delivery

import "github.com/labstack/echo/v4"

type CommentHTTPHandlerContract interface {
	PostComment(ctx echo.Context) error
	GetComments(ctx echo.Context) error
	GetOwnedComments(ctx echo.Context) error
	GetOwnedPhotosComments(ctx echo.Context) error
	UpdateComment(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
}
