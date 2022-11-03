package application_commands

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"math/big"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
)

var GenerateSecretsCmd = &cobra.Command{
	Use:   "app:generate-secrets",
	Short: "Generate secrets for this application",
	Run: func(cmd *cobra.Command, args []string) {
		assertAvailablePRNG()

		randomString, err := GenerateRandomString(32)
		if err != nil {
			log.Fatal(err)
		}
		viper.Set(config_keys.JWTSecret, randomString)

		fmt.Println("New JWT secrets generated")

		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}
	},
}

func assertAvailablePRNG() {
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
