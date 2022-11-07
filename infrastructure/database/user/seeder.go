package user_repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/mygram/application/users/payloads"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/infrastructure/database/common/passwords"
)

func Seed(ctx context.Context, db *gorm.DB, count int) {
	var payloads []entities.User
	faker := gofakeit.New(0)
	repository := UserRepository{DB: db}

	defaultPassword := "12345678"
	hashedPassword, err := passwords.HashPassword(defaultPassword)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= count; i++ {
		payload := user_payloads.UserRegisterPayload{}
		if fakerErr := faker.Struct(&payload); fakerErr != nil {
			log.Fatal(fakerErr)
		}

		payloads = append(payloads, entities.User{
			Username: payload.Username,
			Email:    payload.Email,
			Password: hashedPassword,
			Age:      payload.Age,
		})
	}
	userCount, err := repository.BatchSave(ctx, payloads)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User created: %d\n", userCount)
}
