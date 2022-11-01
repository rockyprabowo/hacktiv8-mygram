package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	jwt_helpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers"
	http_middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type CommentHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    CommentHTTPHandlerContract
}

func NewCommentHTTPRouter(deps http_api_contracts.APIWithJWTRouterDeps, handler CommentHTTPHandlerContract) *CommentHTTPRouter {
	return &CommentHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r CommentHTTPRouter) Setup() {
	jwtMiddlewareConfig := jwt_helpers.BuildEchoJWTMiddlewareConfig(r.JWTService.ParseUserToken)
	jwtMiddleware := middleware.JWTWithConfig(jwtMiddlewareConfig)

	routeGroup := r.Router.Group("/comments")
	{
		routeGroup.Use(jwtMiddleware, http_middlewares.MustHaveValidToken)
		routeGroup.GET("", r.Handler.GetComments)
		//routeGroup.GET("", r.Handler.GetOwnedComments)
		//routeGroup.GET("", r.Handler.GetOwnedPhotosComments)
		routeGroup.POST("", r.Handler.PostComment)
		routeGroup.PUT("/:id", r.Handler.UpdateComment)
		routeGroup.DELETE("/:id", r.Handler.DeleteComment)
	}
}
