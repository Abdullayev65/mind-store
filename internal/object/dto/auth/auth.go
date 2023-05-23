package auth

import "mindstore/internal/object/dto/user"

type LogIn struct {
	Identifier *string
	Password   *string
	//identifier = email or username
}

type LoginRes struct {
	Token string
	User  *user.UserDetail
}

type Available struct {
	Type  int
	Value string
	//type:
	//1 = username
	//2 = email
}
