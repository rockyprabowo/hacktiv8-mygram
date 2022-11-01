package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	http_middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type PhotoHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    PhotoHTTPHandlerContract
}

func NewPhotoHTTPRouter(deps http_api_contracts.APIWithJWTRouterDeps, handler PhotoHTTPHandlerContract) *PhotoHTTPRouter {
	return &PhotoHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r PhotoHTTPRouter) Setup() {
	jwtMiddlewareConfig := jwt_helpers.BuildEchoJWTMiddlewareConfig(r.JWTService.ParseUserToken)
	jwtMiddleware := middleware.JWTWithConfig(jwtMiddlewareConfig)

	routeGroup := r.Router.Group("/photos")
	{
		routeGroup.Use(jwtMiddleware, http_middlewares.MustHaveValidToken)
		routeGroup.GET("", r.Handler.GetPhotos)
		routeGroup.POST("", r.Handler.PostPhoto)
		routeGroup.PUT("/:id", r.Handler.UpdatePhoto)
		routeGroup.DELETE("/:id", r.Handler.DeletePhoto)
	}
}
