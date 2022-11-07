package user_payloads

import "rocky.my.id/git/mygram/domain/entities/value_objects"

type AuthTokenPayload struct {
	ID       int
	Username value_objects.Username
	Email    value_objects.Email
}
