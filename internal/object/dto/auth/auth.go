package auth

type LogIn struct {
	Identifier *string
	Password   *string
	//identifier = email or username
}

type Token struct {
	Token string
}
