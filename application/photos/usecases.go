package photo_use_cases

import contracts "rocky.my.id/git/mygram/application/photos/contracts"

type PhotoUseCases struct {
	Commands contracts.PhotoCommandsContract
	Queries  contracts.PhotoQueriesContract
}

func NewPhotoUseCases(commands contracts.PhotoCommandsContract, queries contracts.PhotoQueriesContract) *PhotoUseCases {
	return &PhotoUseCases{Commands: commands, Queries: queries}
}
