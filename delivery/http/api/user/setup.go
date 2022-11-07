package user_http_delivery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	uc "rocky.my.id/git/mygram/application/users"
	commands "rocky.my.id/git/mygram/application/users/commands"
	queries "rocky.my.id/git/mygram/application/users/queries"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/delivery/http/api/user/handlers"
	"rocky.my.id/git/mygram/infrastructure/database/user"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type UserHTTPDeliveryDeps struct {
	http_api_contracts.APIWithJWTRouterDeps
	UseCases *uc.UserUseCases
}

func Setup(deps UserHTTPDeliveryDeps) {
	var (
		HTTPHandler = user_handlers.NewUserHTTPHandler(deps.UseCases)
		router      = NewUserHTTPRouter(deps.APIWithJWTRouterDeps, HTTPHandler)
	)
	router.Setup()
}

func SetupDefault(engine *echo.Echo, db *gorm.DB, service *jwt_user.UserJWTService) {
	var (
		repository = user_repository.NewUserRepository(db, service)
		useCases   = uc.NewUserUseCases(
			commands.NewUserCommands(repository),
			queries.NewUserQueries(repository),
		)
		apiJWTDeps = http_api_contracts.APIWithJWTRouterDeps{
			Engine:     engine,
			JWTService: service,
		}
		userDeliveryDeps = UserHTTPDeliveryDeps{
			APIWithJWTRouterDeps: apiJWTDeps,
			UseCases:             useCases,
		}
	)
	Setup(userDeliveryDeps)
}
