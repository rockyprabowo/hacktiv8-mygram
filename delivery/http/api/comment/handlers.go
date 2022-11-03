package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	uc "rocky.my.id/git/mygram/application/comments"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/consts"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type CommentHTTPHandler struct {
	UseCases *uc.CommentUseCases
}

func NewCommentHTTPHandler(useCases *uc.CommentUseCases) *CommentHTTPHandler {
	return &CommentHTTPHandler{UseCases: useCases}
}

func (h CommentHTTPHandler) PostComment(ctx echo.Context) error {
	payload := *ctx.Get(consts.Payload).(*payloads.CommentInsertPayload)

	comment, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	comment.DateTime.OmitUpdatedAt()

	return responses.WithData(ctx, comment)
}

func (h CommentHTTPHandler) GetComments(ctx echo.Context) error {
	payload := ctx.Get(consts.Payload).(*payloads.CommentGetAllPayload)

	comments, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), *payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

func (h CommentHTTPHandler) GetOwnedComments(ctx echo.Context) error {
	payload := *ctx.Get(consts.Payload).(*payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedComments(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

func (h CommentHTTPHandler) GetOwnedPhotosComments(ctx echo.Context) error {
	payload := *ctx.Get(consts.Payload).(*payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedPhotoComments(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

func (h CommentHTTPHandler) UpdateComment(ctx echo.Context) error {
	payload := *ctx.Get(consts.Payload).(*payloads.CommentUpdatePayload)

	comment, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	comment.DateTime.OmitCreatedAt()

	return responses.WithData(ctx, comment)
}

func (h CommentHTTPHandler) DeleteComment(ctx echo.Context) error {
	payload := *ctx.Get(consts.Payload).(*payloads.CommentDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.WithError(err)
	}

	return responses.WithDeleteSuccess(ctx, "comment")
}
