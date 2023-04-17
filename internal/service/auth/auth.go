package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	dtoauth "mindstore/internal/object/dto/auth"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"regexp"
	"time"
)

type Service struct {
	User
	emailRegex    *regexp.Regexp
	jwtKey        []byte
	tokenExpiring time.Duration
}

func New(user User) *Service {
	n := new(Service)

	n.User = user
	n.emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	n.jwtKey = []byte("jwt-secret-key")
	n.tokenExpiring = time.Hour << 7

	return n
}

func (s *Service) SignUp(c ctx.Ctx, input *user.UserCreate) error {
	_, err := s.Create(c, input)

	return err
}

func (s *Service) LogIn(c ctx.Ctx, data *dtoauth.LogIn) (*dtoauth.Token, error) {
	if data.Identifier == nil || data.Password == nil {
		return nil, errors.New("identifier and password is required")
	}
	var m *model.User
	var err error
	if s.IsValidEmail(*data.Identifier) {
		m, err = s.User.GetByEmail(c, *data.Identifier)
	} else {
		m, err = s.User.GetByUsername(c, *data.Identifier)
	}

	if err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(m.Id)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// specific functions
func (s *Service) IsValidEmail(email string) bool {
	return s.emailRegex.MatchString(email)
}

func (s *Service) GenerateToken(id hash.Int) (*dtoauth.Token, error) {
	claims := &Claims{
		ID: &id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenExpiring).Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(s.jwtKey)

	if err != nil {
		return nil, err
	}

	token := new(dtoauth.Token)
	token.Token = tokenString

	return token, nil
}

func (s *Service) UserIdFromToken(tokenStr string) (*hash.Int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims.ID, err
}
