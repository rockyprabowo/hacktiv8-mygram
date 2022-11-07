package social_media_http_delivery

import (
	"github.com/labstack/echo/v4"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	contracts "rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	jwt_helpers "rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
	middlewares "rocky.my.id/git/mygram/delivery/http/api/common/middlewares"
	"rocky.my.id/git/mygram/delivery/http/api/social_media/handlers"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type SocialMediaHTTPRouter struct {
	Router     *echo.Echo
	JWTService *jwt_user.UserJWTService
	Handler    social_media_handlers.SocialMediaHTTPHandlerContract
}

func NewSocialMediaHTTPRouter(deps contracts.APIWithJWTRouterDeps, handler social_media_handlers.SocialMediaHTTPHandlerContract) *SocialMediaHTTPRouter {
	return &SocialMediaHTTPRouter{Router: deps.Engine, JWTService: deps.JWTService, Handler: handler}
}

func (r SocialMediaHTTPRouter) Setup() {
	jwtMiddleware := jwt_helpers.BuildEchoJWTMiddleware(r.JWTService.ParseUserToken)

	routeGroup := r.Router.Group("/socialmedias", middlewares.WithJWTValidation(jwtMiddleware)...)
	{
		routeGroup.GET(
			"",
			r.Handler.GetUserSocialMedias,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.SocialMediaGetAllByOwnerPayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.POST(
			"",
			r.Handler.CreateUserSocialMedia,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.SocialMediaInsertPayload) {
					payload.UserID = claims.UserID
				},
			),
		)

		routeGroup.PUT(
			"/:id",
			r.Handler.UpdateUserSocialMedia,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.SocialMediaUpdatePayload) {
					payload.UserID = claims.UserID
				},
			),
		)
		routeGroup.DELETE(
			"/:id",
			r.Handler.DeleteUserSocialMedia,
			middlewares.BindPayloadWithUserClaimsAndValidate(
				func(claims *jwt_user.UserClaims, payload *payloads.SocialMediaDeletePayload) {
					payload.UserID = claims.UserID
				},
			),
		)
	}
}
