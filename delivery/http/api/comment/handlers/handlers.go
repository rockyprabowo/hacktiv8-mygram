package comment_handlers

import (
	"github.com/labstack/echo/v4"
	uc "rocky.my.id/git/mygram/application/comments"
	"rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type CommentHTTPHandler struct {
	UseCases *uc.CommentUseCases
}

func NewCommentHTTPHandler(useCases *uc.CommentUseCases) *CommentHTTPHandler {
	return &CommentHTTPHandler{UseCases: useCases}
}

// PostComment godoc
// @Summary     Post comment
// @Description Creates a new comment on a photo.
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       comment body     comment_payloads.CommentInsertPayload true "Create Comment Request"
// @Success     200   {object} comment_dto.CommentDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /comments [post]
func (h CommentHTTPHandler) PostComment(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*comment_payloads.CommentInsertPayload)

	comment, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	comment.DateTime.OmitUpdatedAt()

	return responses.WithData(ctx, comment)
}

// GetComments godoc
// @Summary     Get comments
// @Description Get all comments
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200   {object} comment_dto.PaginatedCommentWithRelationsDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /comments [get]
func (h CommentHTTPHandler) GetComments(ctx echo.Context) error {
	payload := ctx.Get(constants.Payload).(*comment_payloads.CommentGetAllPayload)

	comments, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), *payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

// GetOwnedComments godoc
// @Summary     Get comments owned by user.
// @Description Get all comments owned by user.
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200   {object} comment_dto.PaginatedCommentWithRelationsDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /me/comments [get]
func (h CommentHTTPHandler) GetOwnedComments(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*comment_payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedComments(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

// GetOwnedPhotosComments godoc
// @Summary     Get comments on photos owned by user.
// @Description Get all comments on photos owned by user.
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200   {object} comment_dto.PaginatedCommentWithRelationsDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /me/photos/comments [get]
func (h CommentHTTPHandler) GetOwnedPhotosComments(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*comment_payloads.CommentGetByOwnerPayload)

	comments, err := h.UseCases.Queries.GetOwnedPhotoComments(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	return responses.WithData(ctx, comments)
}

// UpdateComment godoc
// @Summary     Update comment
// @Description Updates a comment with the given ID.
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Comment ID"
// @Param       comment body   comment_payloads.CommentUpdatePayload true "Update Comment Request"
// @Success     200   {object} comment_dto.CommentDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /comments/{id} [put]
func (h CommentHTTPHandler) UpdateComment(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*comment_payloads.CommentUpdatePayload)

	comment, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	comment.DateTime.OmitCreatedAt()

	return responses.WithData(ctx, comment)
}

// DeleteComment godoc
// @Summary     Delete comment
// @Description Deletes a comment with the given ID.
// @Security 	ApiKeyAuth
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Comment ID"
// @Param       comment body     comment_payloads.CommentDeletePayload true "Delete Comment Request"
// @Success     200   {object} comment_dto.CommentDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /comments/{id} [delete]
func (h CommentHTTPHandler) DeleteComment(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*comment_payloads.CommentDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.WithError(err)
	}

	return responses.WithDeleteSuccess(ctx, "comment")
}
