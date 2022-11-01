package comment_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/application/common/pagination"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/database/common/authorizer"
	"rocky.my.id/git/mygram/infrastructure/database/common/scopes"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(DB *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: DB}
}

func (r CommentRepository) Authorize(ctx context.Context, ownerID, resourceID any) bool {
	return db_authorize.NewGormResourceOwnerAuthorizer(r.DB, &entities.Comment{}, ownerID, resourceID).Execute(ctx)
}

func (r CommentRepository) GetByID(ctx context.Context, payload payloads.CommentGetByIDPayload) (*entities.Comment, error) {
	comment := entities.Comment{}
	err := r.DB.WithContext(ctx).First(&comment, payload.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exceptions.CommentNotFoundError
	}
	return &comment, err
}

func (r CommentRepository) GetAll(ctx context.Context, payload payloads.CommentGetAllPayload) ([]entities.Comment, pagination.State, error) {
	var comments []entities.Comment
	var totalCommentCount int64
	paginationState := payload.PaginationState

	results := r.DB.WithContext(ctx).
		Preload("User").
		Preload("Photo").
		Scopes(scopes.Paginate(&paginationState), scopes.SortDescByID()).
		Find(&comments)
	if results.Error != nil {
		return nil, paginationState, results.Error
	}

	r.DB.WithContext(ctx).
		Model(&entities.Comment{}).
		Count(&totalCommentCount)
	paginationState.SetPaginateTotalCount(totalCommentCount)

	return comments, paginationState, nil
}

func (r CommentRepository) GetOwnedPhotosComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) ([]entities.Comment, pagination.State, error) {
	var comments []entities.Comment
	var userPhotos []entities.Photo
	var userPhotoIDs []int
	var totalCommentCount int64

	paginationState := payload.PaginationState
	userPhotosResult := r.DB.WithContext(ctx).
		Scopes(scopes.OwnedBy(payload.UserID)).
		Find(&userPhotos)
	if userPhotosResult.Error != nil || userPhotosResult.RowsAffected == 0 {
		return nil, paginationState, userPhotosResult.Error
	}

	for _, v := range userPhotos {
		userPhotoIDs = append(userPhotoIDs, v.ID)
	}

	results := r.DB.WithContext(ctx).
		Scopes(scopes.Paginate(&paginationState), scopes.SortDescByID()).
		Preload(clause.Associations).
		Where("photo_id IN ?", userPhotoIDs).
		Find(&comments)
	if results.Error != nil {
		return nil, paginationState, results.Error
	}

	r.DB.WithContext(ctx).
		Model(&entities.Comment{}).
		Where("photo_id IN ?", userPhotoIDs).
		Count(&totalCommentCount)
	paginationState.SetPaginateTotalCount(totalCommentCount)

	return comments, paginationState, nil
}

func (r CommentRepository) GetOwnedComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) ([]entities.Comment, pagination.State, error) {
	var comments []entities.Comment
	var totalCommentCount int64
	paginationState := payload.PaginationState

	results := r.DB.WithContext(ctx).
		Preload(clause.Associations).
		Scopes(
			scopes.Paginate(&paginationState),
			scopes.OwnedBy(payload.UserID),
			scopes.SortDescByID(),
		).
		//Where("user_id = ?", payload.UserID).
		Find(&comments)
	if results.Error != nil {
		return nil, paginationState, results.Error
	}

	r.DB.WithContext(ctx).
		Model(&entities.Comment{}).
		Scopes(scopes.OwnedBy(payload.UserID)).
		//Where("user_id = ?", payload.UserID).
		Count(&totalCommentCount)
	paginationState.SetPaginateTotalCount(totalCommentCount)

	return comments, paginationState, nil
}

func (r CommentRepository) Save(ctx context.Context, payload payloads.CommentInsertPayload) (*entities.Comment, error) {
	comment := &entities.Comment{
		UserID:  payload.UserID,
		PhotoID: payload.PhotoID,
		Message: payload.Message,
	}
	err := r.DB.WithContext(ctx).Create(&comment).Error
	return comment, err
}

func (r CommentRepository) BatchSave(ctx context.Context, payloads []entities.Comment) (int64, error) {
	results := r.DB.WithContext(ctx).CreateInBatches(payloads, 100)
	return results.RowsAffected, results.Error
}

func (r CommentRepository) Update(ctx context.Context, payload payloads.CommentUpdatePayload) (*entities.Comment, error) {
	comment, err := r.GetByID(ctx, payloads.CommentGetByIDPayload{ID: payload.ID})
	if err != nil {
		return nil, err
	}

	comment.Message = payload.Message

	err = r.DB.WithContext(ctx).Save(comment).Error
	return comment, err
}

func (r CommentRepository) Delete(ctx context.Context, payload payloads.CommentDeletePayload) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.Comment{}, payload.ID)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
