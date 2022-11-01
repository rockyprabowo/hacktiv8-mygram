package social_media_http_delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/social_medias"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type SocialMediaHTTPHandler struct {
	UseCases *uc.SocialMediaUseCases
}

func NewSocialMediaHTTPHandler(useCases *uc.SocialMediaUseCases) *SocialMediaHTTPHandler {
	return &SocialMediaHTTPHandler{UseCases: useCases}
}

func (h SocialMediaHTTPHandler) CreateUserSocialMedia(ctx echo.Context) error {
	payload := payloads.SocialMediaInsertPayload{}
	claims := jwt_helpers.ExtractUserClaims(ctx)

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	socialMedia, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	socialMedia.DateTime.OmitUpdatedAt()
	return ctx.JSON(http.StatusOK, socialMedia)
}

func (h SocialMediaHTTPHandler) GetUserSocialMedias(ctx echo.Context) error {
	payload := payloads.SocialMediaGetAllByOwnerPayload{}
	claims := jwt_helpers.ExtractUserClaims(ctx)
	payload.UserID = claims.UserID

	socialMedias, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, SocialMediaCollectionResponse{socialMedias})
}

func (h SocialMediaHTTPHandler) UpdateUserSocialMedia(ctx echo.Context) error {
	payload := payloads.SocialMediaUpdatePayload{}
	claims := jwt_helpers.ExtractUserClaims(ctx)

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	socialMedia, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	socialMedia.DateTime.OmitCreatedAt()
	return ctx.JSON(http.StatusOK, socialMedia)
}

func (h SocialMediaHTTPHandler) DeleteUserSocialMedia(ctx echo.Context) error {
	payload := payloads.SocialMediaDeletePayload{}
	claims := jwt_helpers.ExtractUserClaims(ctx)

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
