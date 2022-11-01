package social_media_http_delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	http_middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type SocialMediaHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    SocialMediaHTTPHandlerContract
}

func NewSocialMediaHTTPRouter(deps http_api_contracts.APIWithJWTRouterDeps, handler SocialMediaHTTPHandlerContract) *SocialMediaHTTPRouter {
	return &SocialMediaHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r SocialMediaHTTPRouter) Setup() {
	jwtMiddlewareConfig := jwt_helpers.BuildEchoJWTMiddlewareConfig(r.JWTService.ParseUserToken)
	jwtMiddleware := middleware.JWTWithConfig(jwtMiddlewareConfig)

	routeGroup := r.Router.Group("/socialmedias")
	{
		routeGroup.Use(jwtMiddleware, http_middlewares.MustHaveValidToken)
		routeGroup.GET("", r.Handler.GetUserSocialMedias)
		routeGroup.POST("", r.Handler.CreateUserSocialMedia)
		routeGroup.PUT("/:id", r.Handler.UpdateUserSocialMedia)
		routeGroup.DELETE("/:id", r.Handler.DeleteUserSocialMedia)
	}
}
