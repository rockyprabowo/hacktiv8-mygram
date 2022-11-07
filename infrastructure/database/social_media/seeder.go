package social_media_repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/mygram/application/social_medias/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

func Seed(ctx context.Context, db *gorm.DB, count int, multiplier int) {
	var payloads []entities.SocialMedia
	maxUserID := count
	maxSocialMediaCount := maxUserID * multiplier
	faker := gofakeit.New(0)
	repository := SocialMediaRepository{DB: db}

	for i := 1; i <= maxSocialMediaCount; i++ {
		payload := social_media_payloads.SocialMediaInsertPayload{}

		if err := faker.Struct(&payload); err != nil {
			log.Fatal(err)
		}
		payload.UserID = faker.Number(1, count)
		payloads = append(payloads, entities.SocialMedia{
			UserID:         payload.UserID,
			Name:           payload.Name,
			SocialMediaURL: payload.SocialMediaURL,
		})
	}
	socialMediaCount, err := repository.BatchSave(ctx, payloads)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SocialMedia created: %d\n", socialMediaCount)

}
