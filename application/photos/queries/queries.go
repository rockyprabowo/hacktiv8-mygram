package photo_queries

import (
	contracts "rocky.my.id/git/mygram/application/photos/contracts"
)

type PhotoQueries struct {
	Repository contracts.PhotoRepositoryContract
}

func NewPhotoQueries(repository contracts.PhotoRepositoryContract) *PhotoQueries {
	return &PhotoQueries{Repository: repository}
}
