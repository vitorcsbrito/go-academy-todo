package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vitorcsbrito/go-academy-todo/model"
	userDto "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/repository/user"
	"github.com/vitorcsbrito/mapper"
	"github.com/vitorcsbrito/utils/errors"
	"time"
)

type UserService struct {
	userRepository user.Repository
}

type UserServiceInterface interface {
	CreateUser(user userDto.CreateUserDTO) (uuid.UUID, error)
	GetUser(id uuid.UUID) (model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateToken(email string) (string, error)
	VerifyToken(tokenString string) error
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo}
}

func (service *UserService) CreateUser(userDto userDto.CreateUserDTO) (uuid.UUID, error) {

	newUser := mapper.DtoToEntityNewUser(userDto)

	id, err := service.userRepository.Save(newUser)

	return id, err
}

func (service *UserService) GetUser(id uuid.UUID) (model.User, error) {
	user, err := service.userRepository.Get(id)

	return user, err
}

func (service *UserService) GetAllUsers() ([]model.User, error) {
	users, err := service.userRepository.GetAll()

	return users, err
}

var secretKey = []byte("c2VjcmV0")

func (service *UserService) CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *UserService) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.ErrInvalidToken
	}

	return nil
}
