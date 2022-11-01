package user_http_delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	http_middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type UserHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    UserHTTPHandlerContract
}

func NewUserHTTPRouter(deps http_api_contracts.APIWithJWTRouterDeps, handler UserHTTPHandlerContract) *UserHTTPRouter {
	return &UserHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r UserHTTPRouter) Setup() {
	jwtMiddlewareConfig := jwt_helpers.BuildEchoJWTMiddlewareConfig(r.JWTService.ParseUserToken)
	jwtMiddleware := middleware.JWTWithConfig(jwtMiddlewareConfig)

	r.Router.POST("/users/register", r.Handler.Register)
	r.Router.POST("/users/login", r.Handler.Login)
	r.Router.GET("/me", r.Handler.GetUser, jwtMiddleware, http_middlewares.MustHaveValidToken)

	routeGroup := r.Router.Group("/users")
	{
		routeGroup.Use(jwtMiddleware, http_middlewares.MustHaveValidToken)
		routeGroup.PUT("", r.Handler.UpdateUser)
		routeGroup.DELETE("", r.Handler.DeleteUser)
	}

}
