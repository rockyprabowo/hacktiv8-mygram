package social_media_http_delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/social_medias"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	errorHelpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/errors"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type SocialMediaHTTPHandler struct {
	UseCases *uc.SocialMediaUseCases
}

func NewSocialMediaHTTPHandler(useCases *uc.SocialMediaUseCases) *SocialMediaHTTPHandler {
	return &SocialMediaHTTPHandler{UseCases: useCases}
}

func (h SocialMediaHTTPHandler) GetUserSocialMedias(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.SocialMediaGetAllByOwnerPayload)

	socialMedias, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)

	}
	return ctx.JSON(http.StatusOK, SocialMediaCollectionResponse{socialMedias})
}

func (h SocialMediaHTTPHandler) CreateUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.SocialMediaInsertPayload)

	socialMedia, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)

	}
	socialMedia.DateTime.OmitUpdatedAt()
	return ctx.JSON(http.StatusOK, socialMedia)
}

func (h SocialMediaHTTPHandler) UpdateUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.SocialMediaUpdatePayload)

	socialMedia, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)

	}
	socialMedia.DateTime.OmitCreatedAt()
	return ctx.JSON(http.StatusOK, socialMedia)
}

func (h SocialMediaHTTPHandler) DeleteUserSocialMedia(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.SocialMediaDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return errorHelpers.ExtractError(err)

	}
	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
