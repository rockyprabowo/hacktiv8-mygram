package passwords

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"rocky.my.id/git/mygram/configurations/config/keys"
)

func HashPassword(password string) ([]byte, error) {
	var cost = viper.GetInt(config_keys.BcryptCost)

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return hashBytes, err
}

func ComparePasswordHash(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
