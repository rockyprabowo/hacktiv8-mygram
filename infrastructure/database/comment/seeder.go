package comment_repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

func Seed(ctx context.Context, db *gorm.DB, count int, multiplier int) {
	var payloads []entities.Comment

	faker := gofakeit.New(0)
	repository := CommentRepository{DB: db}

	maxUserID := count
	maxPhotoID := count * multiplier
	maxCommentCount := maxPhotoID * multiplier

	for i := 1; i <= maxCommentCount; i++ {
		payload := comment_payloads.CommentInsertPayload{}
		if err := faker.Struct(&payload); err != nil {
			log.Println(err)
		}
		payload.UserID = faker.Number(1, maxUserID)
		payload.PhotoID = faker.Number(1, maxPhotoID)
		payloads = append(payloads, entities.Comment{
			UserID:  payload.UserID,
			PhotoID: payload.PhotoID,
			Message: payload.Message,
		})
	}
	commentsCount, err := repository.BatchSave(ctx, payloads)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Comments created: %d\n", commentsCount)
}
