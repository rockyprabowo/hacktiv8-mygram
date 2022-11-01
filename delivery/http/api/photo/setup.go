package photo_http_delivery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	uc "rocky.my.id/git/mygram/application/photos"
	commands "rocky.my.id/git/mygram/application/photos/commands"
	queries "rocky.my.id/git/mygram/application/photos/queries"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/infrastructure/database/photo"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type PhotoHTTPDeliveryDeps struct {
	http_api_contracts.APIWithJWTRouterDeps
	UseCases *uc.PhotoUseCases
}

func Setup(deps PhotoHTTPDeliveryDeps) {
	var (
		HTTPHandler = NewPhotoHTTPHandler(deps.UseCases)
		router      = NewPhotoHTTPRouter(deps.APIWithJWTRouterDeps, HTTPHandler)
	)
	router.Setup()
}
func SetupDefault(engine *echo.Echo, db *gorm.DB, service *jwt_user.UserJWTService) {
	var (
		repository = photo_repository.NewPhotoRepository(db)
		useCases   = uc.NewPhotoUseCases(
			commands.NewPhotoCommands(repository),
			queries.NewPhotoQueries(repository),
		)
		apiJWTDeps = http_api_contracts.APIWithJWTRouterDeps{
			Engine:     engine,
			JWTService: service,
		}
		socialMediaDeliveryDeps = PhotoHTTPDeliveryDeps{
			APIWithJWTRouterDeps: apiJWTDeps,
			UseCases:             useCases,
		}
	)
	Setup(socialMediaDeliveryDeps)
}
