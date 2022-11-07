package social_media_http_delivery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	uc "rocky.my.id/git/mygram/application/social_medias"
	commands "rocky.my.id/git/mygram/application/social_medias/commands"
	queries "rocky.my.id/git/mygram/application/social_medias/queries"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/delivery/http/api/social_media/handlers"
	"rocky.my.id/git/mygram/infrastructure/database/social_media"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type SocialMediaHTTPDeliveryDeps struct {
	http_api_contracts.APIWithJWTRouterDeps
	UseCases *uc.SocialMediaUseCases
}

func Setup(deps SocialMediaHTTPDeliveryDeps) {
	var (
		HTTPHandler = social_media_handlers.NewSocialMediaHTTPHandler(deps.UseCases)
		router      = NewSocialMediaHTTPRouter(deps.APIWithJWTRouterDeps, HTTPHandler)
	)
	router.Setup()
}
func SetupDefault(engine *echo.Echo, db *gorm.DB, service *jwt_user.UserJWTService) {
	var (
		repository = social_media_repository.NewSocialMediaRepository(db)
		useCases   = uc.NewSocialMediaUseCases(
			commands.NewSocialMediaCommands(repository),
			queries.NewSocialMediaQueries(repository),
		)
		apiJWTDeps = http_api_contracts.APIWithJWTRouterDeps{
			Engine:     engine,
			JWTService: service,
		}
		socialMediaDeliveryDeps = SocialMediaHTTPDeliveryDeps{
			APIWithJWTRouterDeps: apiJWTDeps,
			UseCases:             useCases,
		}
	)
	Setup(socialMediaDeliveryDeps)
}
