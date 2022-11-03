package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type PhotoHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    PhotoHTTPHandlerContract
}

func NewPhotoHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler PhotoHTTPHandlerContract) *PhotoHTTPRouter {
	return &PhotoHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r PhotoHTTPRouter) Setup() {
	routeGroup := r.Router.Group("/photos")
	{
		jwtMiddlewares := middlewares.WithJWTValidation(r.JWTService.ParseUserToken)
		routeGroup.Use(jwtMiddlewares...)

		routeGroup.GET(
			"",
			r.Handler.GetPhotos,
			middlewares.BindPayloadAndValidate(&payloads.PhotoGetAllPayload{}),
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
