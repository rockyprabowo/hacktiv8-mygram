package user_handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	uc "rocky.my.id/git/mygram/application/users"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/domain/exceptions"
)

type UserHTTPHandler struct {
	UseCases *uc.UserUseCases
}

func NewUserHTTPHandler(useCases *uc.UserUseCases) *UserHTTPHandler {
	return &UserHTTPHandler{UseCases: useCases}
}

// GetUser godoc
// @Summary     Get current user
// @Description Get current user
// @Security 	ApiKeyAuth
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200   {object} user_dto.UserDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /users [get]
func (h UserHTTPHandler) GetUser(ctx echo.Context) error {
	claims := jwt_helpers.ExtractUserClaims(ctx)

	user, err := h.UseCases.Queries.GetUserProfile(ctx.Request().Context(), payloads.UserGetPayload{Key: "ID", Value: claims.UserID})
	if err != nil {
		if errors.Is(err, exceptions.UserNotFoundError) {
			return responses.EchoErrorResponse(http.StatusNotFound, err.Error())
		}
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

// Login godoc
// @Summary     Login user
// @Description Login user
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       login body	   user_payloads.UserLoginPayload  true "User Login Request"
// @Success     200   {object} authentication_dto.AuthTokenDTO
// @Failure     401   {object} responses.ErrorResult
// @Router      /users/login [post]
func (h UserHTTPHandler) Login(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.UserLoginPayload)

	_, token, err := h.UseCases.Queries.AuthenticateUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(http.StatusOK, token)
}

// Register godoc
// @Summary     Register new user
// @Description Registers new user
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       register body	   user_payloads.UserRegisterPayload  true "User Register Request"
// @Success     200   {object} authentication_dto.AuthTokenDTO
// @Failure     422   {object} responses.ErrorResult
// @Router      /users/register [post]
func (h UserHTTPHandler) Register(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.UserRegisterPayload)

	newUser, err := h.UseCases.Commands.RegisterUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	newUser.DateTime.Omit()

	return ctx.JSON(http.StatusOK, newUser)
}

// UpdateUser godoc
// @Summary     Update user
// @Description Updates current user.
// @Security 	ApiKeyAuth
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body   user_payloads.UserProfileUpdatePayload true "Update User Request"
// @Success     200   {object} user_dto.UserDTO
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /users [put]
func (h UserHTTPHandler) UpdateUser(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.UserProfileUpdatePayload)

	updatedUser, err := h.UseCases.Commands.UpdateUser(ctx.Request().Context(), payload)
	if err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}
	updatedUser.DateTime.OmitCreatedAt()

	return ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary     Delete user
// @Description Deletes current user.
// @Security 	ApiKeyAuth
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body   user_payloads.UserDeletePayload true "Delete User Request"
// @Success     200   {object} responses.InfoResult
// @Failure     401   {object} responses.ErrorResult
// @Failure     422   {object} responses.ErrorResult
// @Router      /users [delete]
func (h UserHTTPHandler) DeleteUser(ctx echo.Context) error {
	payload := *ctx.Get(constants.Payload).(*payloads.UserDeletePayload)

	deleted, err := h.UseCases.Commands.DeleteUser(ctx.Request().Context(), payload)
	if !deleted || err != nil {
		return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
	}

	return responses.WithDeleteSuccess(ctx, "user")
}
