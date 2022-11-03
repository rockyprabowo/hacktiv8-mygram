package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/comments"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	errorHelpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/errors"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type CommentHTTPHandler struct {
	UseCases *uc.CommentUseCases
}

func NewCommentHTTPHandler(useCases *uc.CommentUseCases) *CommentHTTPHandler {
	return &CommentHTTPHandler{UseCases: useCases}
}

func (h CommentHTTPHandler) PostComment(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.CommentInsertPayload)

	comment, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	comment.DateTime.OmitUpdatedAt()

	return ctx.JSON(http.StatusOK, comment)
}

func (h CommentHTTPHandler) GetComments(ctx echo.Context) error {
	payload := ctx.Get("payload").(*payloads.CommentGetAllPayload)

	comments, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), *payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	return ctx.JSON(http.StatusOK, comments)
}

func (h CommentHTTPHandler) GetOwnedComments(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedComments(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	return ctx.JSON(http.StatusOK, comments)
}

func (h CommentHTTPHandler) GetOwnedPhotosComments(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedPhotosComments(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	return ctx.JSON(http.StatusOK, comments)
}

func (h CommentHTTPHandler) UpdateComment(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.CommentUpdatePayload)

	comment, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	comment.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, comment)
}

func (h CommentHTTPHandler) DeleteComment(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.CommentDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return errorHelpers.ExtractError(err)
	}

	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
