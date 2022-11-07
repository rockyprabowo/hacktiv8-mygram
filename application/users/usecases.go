package user_use_cases

import contracts "rocky.my.id/git/mygram/application/users/contracts"

type UserUseCases struct {
	Queries  contracts.UserQueriesContract
	Commands contracts.UserCommandsContract
}

func NewUserUseCases(commands contracts.UserCommandsContract, queries contracts.UserQueriesContract) *UserUseCases {
	return &UserUseCases{Commands: commands, Queries: queries}
}
