package photo_commands

import (
	contracts "rocky.my.id/git/mygram/application/photos/contracts"
)

type PhotoCommands struct {
	Repository contracts.PhotoRepositoryContract
}

func NewPhotoCommands(repository contracts.PhotoRepositoryContract) *PhotoCommands {
	return &PhotoCommands{Repository: repository}
}
