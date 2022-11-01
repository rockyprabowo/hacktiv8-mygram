package user_http_delivery

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/users"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/domain/exceptions"
)

type UserHTTPHandler struct {
	UseCases *uc.UserUseCases
}

func NewUserHTTPHandler(useCases *uc.UserUseCases) *UserHTTPHandler {
	return &UserHTTPHandler{UseCases: useCases}
}

func (h UserHTTPHandler) GetUser(ctx echo.Context) error {
	claims := jwt_helpers.ExtractUserClaims(ctx)

	user, err := h.UseCases.Queries.GetUserProfile(ctx.Request().Context(), payloads.UserGetPayload{Key: "ID", Value: claims.UserID})
	if err != nil {
		if errors.Is(err, exceptions.UserNotFoundError) {
			return responses.EchoErrorResponse(http.StatusNotFound, err.Error())
		}
		return responses.EchoErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (h UserHTTPHandler) Login(ctx echo.Context) error {
	var payload payloads.UserLoginPayload

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	_, token, err := h.UseCases.Queries.AuthenticateUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(http.StatusOK, token)
}

func (h UserHTTPHandler) Register(ctx echo.Context) error {
	var payload payloads.UserRegisterPayload

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	newUser, err := h.UseCases.Commands.RegisterUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	newUser.DateTime.Omit()
	return ctx.JSON(http.StatusOK, newUser)
}

func (h UserHTTPHandler) UpdateUser(ctx echo.Context) error {
	var payload payloads.UserProfileUpdatePayload

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	updatedUser, err := h.UseCases.Commands.UpdateUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	updatedUser.DateTime.OmitCreatedAt()
	return ctx.JSON(http.StatusOK, updatedUser)
}

func (h UserHTTPHandler) DeleteUser(ctx echo.Context) error {
	var payload payloads.UserDeletePayload

	if bindErr := ctx.Bind(&payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}

	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}

	deleted, err := h.UseCases.Commands.DeleteUser(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, responses.InfoResult{Message: DeleteSuccessMessage})
}
