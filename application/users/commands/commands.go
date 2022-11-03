package user_commands

import (
	contracts "rocky.my.id/git/mygram/application/users/contracts"
)

type UserCommands struct {
	Repository contracts.UserRepositoryContracts
}

func NewUserCommands(repository contracts.UserRepositoryContracts) *UserCommands {
	return &UserCommands{Repository: repository}
}
