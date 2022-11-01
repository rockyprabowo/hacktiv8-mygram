package comment_http_delivery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	uc "rocky.my.id/git/mygram/application/comments"
	commands "rocky.my.id/git/mygram/application/comments/commands"
	queries "rocky.my.id/git/mygram/application/comments/queries"
	"rocky.my.id/git/mygram/delivery/http/api/common/contracts"
	"rocky.my.id/git/mygram/infrastructure/database/comment"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type CommentHTTPDeliveryDeps struct {
	http_api_contracts.APIWithJWTRouterDeps
	UseCases *uc.CommentUseCases
}

func Setup(deps CommentHTTPDeliveryDeps) {
	var (
		HTTPHandler = NewCommentHTTPHandler(deps.UseCases)
		router      = NewCommentHTTPRouter(deps.APIWithJWTRouterDeps, HTTPHandler)
	)
	router.Setup()
}
func SetupDefault(engine *echo.Echo, db *gorm.DB, service *jwt_user.UserJWTService) {
	var (
		repository = comment_repository.NewCommentRepository(db)
		useCases   = uc.NewCommentUseCases(
			commands.NewCommentCommands(repository),
			queries.NewCommentQueries(repository),
		)
		apiJWTDeps = http_api_contracts.APIWithJWTRouterDeps{
			Engine:     engine,
			JWTService: service,
		}
		socialMediaDeliveryDeps = CommentHTTPDeliveryDeps{
			APIWithJWTRouterDeps: apiJWTDeps,
			UseCases:             useCases,
		}
	)
	Setup(socialMediaDeliveryDeps)
}
