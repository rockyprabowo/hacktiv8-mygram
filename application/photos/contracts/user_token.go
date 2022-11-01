package photo_contracts

type UserTokenAuthenticationContract interface {
	GenerateUserToken(id int, username, email string) (string, error)
	ParseUserToken(token string) (any, error)
}
