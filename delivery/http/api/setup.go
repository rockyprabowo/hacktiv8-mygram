package api

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"rocky.my.id/git/mygram/delivery/http/api/comment"
	"rocky.my.id/git/mygram/delivery/http/api/photo"
	"rocky.my.id/git/mygram/delivery/http/api/social_media"
	"rocky.my.id/git/mygram/delivery/http/api/user"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func SetupDefault(engine *echo.Echo, db *gorm.DB) {
	jwtSecret := viper.GetString(config_keys.JWTSecret)
	jwtService := jwt_user.NewUserJWTService([]byte(jwtSecret))

	user_http_delivery.SetupDefault(engine, db, jwtService)
	social_media_http_delivery.SetupDefault(engine, db, jwtService)
	photo_http_delivery.SetupDefault(engine, db, jwtService)
	comment_http_delivery.SetupDefault(engine, db, jwtService)
}