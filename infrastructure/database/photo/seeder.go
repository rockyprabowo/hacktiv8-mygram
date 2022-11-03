package photo_repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/mygram/application/photos/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

func Seed(ctx context.Context, db *gorm.DB, count int, multiplier int) {
	var payloads []entities.Photo
	maxUserID := count
	maxPhotoCount := count * multiplier
	faker := gofakeit.New(0)
	repository := PhotoRepository{DB: db}

	for i := 1; i <= maxPhotoCount; i++ {
		payload := photo_payloads.PhotoInsertPayload{}
		err := faker.Struct(&payload)
		if err != nil {
			log.Fatal(err)
		}
		payload.UserID = faker.Number(1, maxUserID)
		payloads = append(payloads, entities.Photo{
			Title:    payload.Title,
			Caption:  payload.Caption,
			PhotoURL: payload.PhotoURL,
			UserID:   payload.UserID,
		})
	}

	photoCount, err := repository.BatchSave(ctx, payloads)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Photo created: %d\n", photoCount)
}
