package user_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type UserHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    UserHTTPHandlerContract
}

func NewUserHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler UserHTTPHandlerContract) *UserHTTPRouter {
	return &UserHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r UserHTTPRouter) Setup() {

	r.Router.POST(
		"/users/register",
		r.Handler.Register,
		middlewares.BindPayloadAndValidate(&payloads.UserRegisterPayload{}),
	)

	r.Router.POST(
		"/users/login",
		r.Handler.Login,
		middlewares.BindPayloadAndValidate(&payloads.UserLoginPayload{}),
	)

	r.Router.GET(
		"/me",
		r.Handler.GetUser,
		middlewares.WithJWTValidation(r.JWTService.ParseUserToken)...,
	)

	routeGroup := r.Router.Group("/users")
	{
		jwtMiddlewares := middlewares.WithJWTValidation(r.JWTService.ParseUserToken)
		routeGroup.Use(jwtMiddlewares...)

		routeGroup.PUT(
			"",
			r.Handler.UpdateUser,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.UserProfileUpdatePayload) {
					payload.ID = claims.UserID
				},
			),
		)

		routeGroup.DELETE(
			"",
			r.Handler.DeleteUser,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.UserDeletePayload) {
					payload.ID = claims.UserID
				},
			),
		)
	}

}
