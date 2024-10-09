package service

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/repository/user"
	"github.com/vitorcsbrito/mapper"
	"github.com/vitorcsbrito/utils"
	"github.com/vitorcsbrito/utils/errors"
	"net/http"
	"os"
	"time"
)

type UserService struct {
	userRepository user.Repository
	pubKey         *rsa.PublicKey
	privKey        *rsa.PrivateKey
}

type UserServiceInterface interface {
	CreateUser(user CreateUserDTO) (uuid.UUID, error)
	GetUser(id uuid.UUID) (model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateToken(email model.User) (string, error)
	VerifyToken(tokenString string) error
	AuthenticateUser(authDetails AuthDTO) (*Token, error)
}

func NewUserService(repo user.Repository) *UserService {
	publicKey, _ := os.ReadFile("./private/pub.key")
	keyPub, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	privateKey, _ := os.ReadFile("./private/priv.key")
	keyPriv, _ := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	return &UserService{
		userRepository: repo,
		privKey:        keyPriv,
		pubKey:         keyPub,
	}
}

func (userService *UserService) CreateUser(userDto CreateUserDTO) (uuid.UUID, error) {

	newUser := mapper.DtoToEntityNewUser(userDto)

	id, err := userService.userRepository.Save(newUser)

	return id, err
}

func (userService *UserService) GetUser(id uuid.UUID) (model.User, error) {
	user, err := userService.userRepository.Get(id)

	return user, err
}

func (userService *UserService) GetAllUsers() ([]model.User, error) {
	users, err := userService.userRepository.GetAll()

	return users, err
}

func (userService *UserService) CreateToken(user model.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(utils.CookieAge)
	claims := &model.Claims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(userService.privKey)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}

func (userService *UserService) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return userService.pubKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.ErrInvalidToken
	}

	return nil
}

func (userService *UserService) AuthenticateUser(authDetails AuthDTO) (*Token, time.Time, error) {
	usr, err := userService.userRepository.GetByUsername(authDetails.Username)
	if err != nil {
		return &Token{}, time.Now(), err
	}

	if usr.Password == authDetails.Password {
		token, expTime, err := userService.CreateToken(usr)
		if err != nil {
			return &Token{}, expTime, err
		}
		return NewToken(token, utils.CookieAge, expTime), time.Now(), err
	}

	return &Token{}, time.Now(), errors.ErrInvalidCredentials
}

func (userService *UserService) GetToken(token Token) http.Cookie {
	return http.Cookie{
		Name:     "todo_auth_cookie",
		Value:    token.Token,
		Path:     "/",
		MaxAge:   int(token.Age.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
		Expires:  token.Expires,
	}
}
