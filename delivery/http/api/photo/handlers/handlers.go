package photo_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/photos"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

type PhotoHTTPHandler struct {
	UseCases *uc.PhotoUseCases
}

func NewPhotoHTTPHandler(useCases *uc.PhotoUseCases) *PhotoHTTPHandler {
	return &PhotoHTTPHandler{UseCases: useCases}
}

// PostPhoto godoc
// @Summary     Post photo
// @Description Creates a new photo on a photo.
// @Security 	ApiKeyAuth
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       photo body     photo_payloads.PhotoInsertPayload true "Create Photo Request"
// @Success     200   {object} photo_dto.PhotoDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /photos [post]
func (h PhotoHTTPHandler) PostPhoto(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotoInsertPayload)

	photo, err := h.UseCases.Commands.Save(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	photo.DateTime.OmitUpdatedAt()

	return ctx.JSON(http.StatusOK, photo)
}

// GetPhotos godoc
// @Summary     Get photos
// @Description Get all photos
// @Security 	ApiKeyAuth
// @Tags        photos
// @Accept      json
// @Produce     json
// @Success     200   {object} photo_dto.PaginatedPhotoWithUserDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /photos [get]
func (h PhotoHTTPHandler) GetPhotos(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotoGetAllPayload)

	photos, err := h.UseCases.Queries.GetAll(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}

	return ctx.JSON(http.StatusOK, photos)
}

// GetPhotoByID godoc
// @Summary     Get single photo
// @Description Get a photo by its ID from the database.
// @Tags        photos
// @Produce     json
// @Param       id    path     int									 true "Photo ID"
// @Success     200   {object} photo_dto.PhotoDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /photos/{id} [get]
func (h PhotoHTTPHandler) GetPhotoByID(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotoGetByIDPayload)

	photos, err := h.UseCases.Queries.GetByID(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}

	return ctx.JSON(http.StatusOK, photos)
}

// GetOwnedPhotos godoc
// @Summary     Get photos owned by user.
// @Description Get all photos
// @Security 	ApiKeyAuth
// @Tags        photos
// @Accept      json
// @Produce     json
// @Success     200   {object} photo_dto.PaginatedPhotoWithUserDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /me/photos [get]
func (h PhotoHTTPHandler) GetOwnedPhotos(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotosGetByOwnerPayload)

	photos, err := h.UseCases.Queries.GetOwnedPhotos(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}

	return ctx.JSON(http.StatusOK, photos)
}

// UpdatePhoto godoc
// @Summary     Update photo
// @Description Updates a photo with the given ID.
// @Security 	ApiKeyAuth
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Photo ID"
// @Param       photo body     photo_payloads.PhotoUpdatePayload true "Update Photo Request"
// @Success     200   {object} photo_dto.PhotoDTO
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /photos/{id} [put]
func (h PhotoHTTPHandler) UpdatePhoto(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotoUpdatePayload)

	photo, err := h.UseCases.Commands.Update(ctx.Request().Context(), payload)
	if err != nil {
		return responses.WithError(err)
	}
	photo.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary     Delete photo
// @Description Deletes a photo with the given ID.
// @Security 	ApiKeyAuth
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id    path     int									 true "Photo ID"
// @Param       photo body     photo_payloads.PhotoDeletePayload true "Delete Photo Request"
// @Success     200   {object} responses.InfoResult
// @Failure     404   {object} responses.ErrorResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /photos/{id} [delete]
func (h PhotoHTTPHandler) DeletePhoto(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.PhotoDeletePayload)

	deleted, err := h.UseCases.Commands.Delete(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.WithError(err)
	}

	return responses.WithDeleteSuccess(ctx, "photo")
}
