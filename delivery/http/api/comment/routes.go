package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/delivery/http/api/comment/handlers"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	jwt_helpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type CommentHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    comment_handlers.CommentHTTPHandlerContract
}

func NewCommentHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler comment_handlers.CommentHTTPHandlerContract) *CommentHTTPRouter {
	return &CommentHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r CommentHTTPRouter) Setup() {
	jwtMiddleware := jwt_helpers.BuildEchoJWTMiddleware(r.JWTService.ParseUserToken)

	r.Router.GET(
		"/me/comments",
		r.Handler.GetOwnedComments,
		middlewares.WithJWTValidation(
			jwtMiddleware,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.CommentGetByOwnerPayload) {
					payload.UserID = claims.UserID
				},
			),
		)...,
	)

	r.Router.GET(
		"/me/photos/comments",
		r.Handler.GetOwnedPhotosComments,
		middlewares.WithJWTValidation(
			jwtMiddleware,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.CommentGetByOwnerPayload) {
					payload.UserID = claims.UserID
				},
			),
		)...,
	)

	routeGroup := r.Router.Group("/comments", middlewares.WithJWTValidation(jwtMiddleware)...)
	{
		routeGroup.GET(
			"",
			r.Handler.GetComments,
			middlewares.BindPayloadAndValidate(&payloads.CommentGetAllPayload{}),
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
