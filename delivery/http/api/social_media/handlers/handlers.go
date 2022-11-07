package social_media_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/social_medias"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type SocialMediaHTTPHandler struct {
	UseCases *uc.SocialMediaUseCases
}

func NewSocialMediaHTTPHandler(useCases *uc.SocialMediaUseCases) *SocialMediaHTTPHandler {
	return &SocialMediaHTTPHandler{UseCases: useCases}
}

// GetUserSocialMedias godoc
// @Summary     Get user's social media
// @Description Get user's social media
// @Security 	ApiKeyAuth
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Success     200   {object} SocialMediaCollectionResponse
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /socialmedias [get]
func (h SocialMediaHTTPHandler) GetUserSocialMedias(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.SocialMediaGetAllByOwnerPayload)

	socialMedias, err := h.UseCases.Queries.GetOwnedSocialMedia(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)

	}

	return ctx.JSON(http.StatusOK, SocialMediaCollectionResponse{SocialMedias: socialMedias})
}

// CreateUserSocialMedia godoc
// @Summary     Create social media.
// @Description Creates a new social media of a user.
// @Security 	ApiKeyAuth
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       social_media body     social_media_payloads.SocialMediaInsertPayload true "Create SocialMedia Request"
// @Success     200   {object} social_media_dto.SocialMediaDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /socialmedias [post]
func (h SocialMediaHTTPHandler) CreateUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.SocialMediaInsertPayload)

	socialMedia, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)

	}
	socialMedia.DateTime.OmitUpdatedAt()

	return ctx.JSON(http.StatusOK, socialMedia)
}

// UpdateUserSocialMedia godoc
// @Summary     Update social media
// @Description Updates a social media with the given ID.
// @Security 	ApiKeyAuth
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Social Media ID"
// @Param       social_media body     social_media_payloads.SocialMediaUpdatePayload true "Update Social Media Request"
// @Success     200   {object} social_media_dto.SocialMediaDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /socialmedias/{id} [put]
func (h SocialMediaHTTPHandler) UpdateUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.SocialMediaUpdatePayload)

	socialMedia, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)

	}
	socialMedia.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, socialMedia)
}

// DeleteUserSocialMedia godoc
// @Summary     Delete social media
// @Description Deletes a social media with the given ID.
// @Security 	ApiKeyAuth
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Social Media ID"
// @Param       social_media body     social_media_payloads.SocialMediaDeletePayload true "Delete Social Media Request"
// @Success     200   {object} responses.InfoResult
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /socialmedias/{id} [delete]
func (h SocialMediaHTTPHandler) DeleteUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.SocialMediaDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.WithError(err)

	}

	return responses.WithDeleteSuccess(ctx, "social media")
}
