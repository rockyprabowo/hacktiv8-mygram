package user_queries

import (
	contracts "rocky.my.id/git/mygram/application/users/contracts"
)

type UserQueries struct {
	Repository contracts.UserRepositoryContracts
}

func NewUserQueries(repository contracts.UserRepositoryContracts) *UserQueries {
	return &UserQueries{Repository: repository}
}
