package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/photos"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type PhotoHTTPHandler struct {
	UseCases *uc.PhotoUseCases
}

func NewPhotoHTTPHandler(useCases *uc.PhotoUseCases) *PhotoHTTPHandler {
	return &PhotoHTTPHandler{UseCases: useCases}
}

func (h PhotoHTTPHandler) PostPhoto(ctx echo.Context) error {
	claims := jwt_helpers.ExtractUserClaims(ctx)
	payload := payloads.PhotoInsertPayload{}

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	photo, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	photo.DateTime.OmitUpdatedAt()

	return ctx.JSON(http.StatusOK, photo)
}

func (h PhotoHTTPHandler) GetPhotos(ctx echo.Context) error {
	payload := payloads.PhotoGetAllPayload{}

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	photos, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, photos)
}

func (h PhotoHTTPHandler) UpdatePhoto(ctx echo.Context) error {
	claims := jwt_helpers.ExtractUserClaims(ctx)
	payload := payloads.PhotoUpdatePayload{}

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	photo, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	photo.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, photo)
}

func (h PhotoHTTPHandler) DeletePhoto(ctx echo.Context) error {
	claims := jwt_helpers.ExtractUserClaims(ctx)
	payload := payloads.PhotoDeletePayload{}

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	payload.UserID = claims.UserID

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
