package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/photos"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	errorHelpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/errors"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type PhotoHTTPHandler struct {
	UseCases *uc.PhotoUseCases
}

func NewPhotoHTTPHandler(useCases *uc.PhotoUseCases) *PhotoHTTPHandler {
	return &PhotoHTTPHandler{UseCases: useCases}
}

func (h PhotoHTTPHandler) PostPhoto(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.PhotoInsertPayload)

	photo, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	photo.DateTime.OmitUpdatedAt()

	return ctx.JSON(http.StatusOK, photo)
}
func (h PhotoHTTPHandler) GetPhotos(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.PhotoGetAllPayload)

	photos, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	return ctx.JSON(http.StatusOK, photos)
}

func (h PhotoHTTPHandler) UpdatePhoto(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.PhotoUpdatePayload)

	photo, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return errorHelpers.ExtractError(err)
	}
	photo.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, photo)
}

func (h PhotoHTTPHandler) DeletePhoto(ctx echo.Context) error {
	payload := *ctx.Get("payload").(*payloads.PhotoDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return errorHelpers.ExtractError(err)
	}
	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
