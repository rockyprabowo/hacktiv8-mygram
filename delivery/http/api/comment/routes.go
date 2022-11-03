package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type CommentHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    CommentHTTPHandlerContract
}

func NewCommentHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler CommentHTTPHandlerContract) *CommentHTTPRouter {
	return &CommentHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r CommentHTTPRouter) Setup() {
	routeGroup := r.Router.Group("/comments")
	{
		jwtMiddlewares := middlewares.WithJWTValidation(r.JWTService.ParseUserToken)
		routeGroup.Use(jwtMiddlewares...)

		routeGroup.GET(
			"",
			r.Handler.GetComments,
			middlewares.BindPayloadAndValidate(&payloads.CommentGetAllPayload{}),
		)

		routeGroup.GET(
			"/owned",
			r.Handler.GetOwnedComments,
			middlewares.BindPayloadAndValidate(&payloads.CommentGetByOwnerPayload{}),
		)

		routeGroup.GET(
			"/ownedPhotos",
			r.Handler.GetOwnedPhotosComments,
			middlewares.BindPayloadAndValidate(&payloads.CommentGetByOwnerPayload{}),
		)

		routeGroup.POST(
			"",
			r.Handler.PostComment,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.CommentInsertPayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.PUT(
			"/:id",
			r.Handler.UpdateComment,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.CommentUpdatePayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.DELETE(
			"/:id",
			r.Handler.DeleteComment,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.CommentDeletePayload) {
					payload.UserID = claims.UserID
				},
			),
		)
	}
}
