package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	jwt_helpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/delivery/http/api/photo/handlers"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type PhotoHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    photo_handlers.PhotoHTTPHandlerContract
}

func NewPhotoHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler photo_handlers.PhotoHTTPHandlerContract) *PhotoHTTPRouter {
	return &PhotoHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r PhotoHTTPRouter) Setup() {
	jwtMiddleware := jwt_helpers.BuildEchoJWTMiddleware(r.JWTService.ParseUserToken)

	r.Router.GET(
		"/me/photos",
		r.Handler.GetOwnedPhotos,
		middlewares.WithJWTValidation(
			jwtMiddleware,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.PhotosGetByOwnerPayload) {
					payload.UserID = claims.UserID
				},
			),
		)...,
	)

	routeGroup := r.Router.Group("/photos", middlewares.WithJWTValidation(jwtMiddleware)...)
	{
		routeGroup.GET(
			"",
			r.Handler.GetPhotos,
			middlewares.BindPayloadAndValidate(&payloads.PhotoGetAllPayload{}),
		)

		routeGroup.GET(
			"/:id",
			r.Handler.GetPhotoByID,
			middlewares.BindPayloadAndValidate(&payloads.PhotoGetByIDPayload{}),
		)

		routeGroup.POST(
			"",
			r.Handler.PostPhoto,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.PhotoInsertPayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.PUT(
			"/:id",
			r.Handler.UpdatePhoto,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.PhotoUpdatePayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.DELETE(
			"/:id",
			r.Handler.DeletePhoto,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.PhotoDeletePayload) {
					payload.UserID = claims.UserID
				},
			),
		)
	}
}
