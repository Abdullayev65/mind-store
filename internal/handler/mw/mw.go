package mw

type MiddleWere struct {
	auth Auth
	user User
}

func New(auth Auth, user User) *MiddleWere {
	n := new(MiddleWere)

	n.auth = auth
	n.user = user

	return n
}
