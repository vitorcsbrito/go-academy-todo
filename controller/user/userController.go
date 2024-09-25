package user

import (
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	. "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/service"
	. "github.com/vitorcsbrito/middleware"
	. "github.com/vitorcsbrito/utils/errors"
	. "github.com/vitorcsbrito/utils/requests"
	"log"
	"net/http"
)

type Controller struct {
	userService *service.UserService
}

func (userController *Controller) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", CreateUser(userController))
	mux.HandleFunc("GET /users", GetAllUsers(userController))
	mux.HandleFunc("POST /auth", LoginHandler(userController))
	mux.Handle("POST /protected", Auth(ProtectedHandler(userController)))
}

func NewUserController(userService *service.UserService) *Controller {
	u := &Controller{
		userService,
	}
	return u
}

func CreateUser(userController *Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			NewBadRequestResponse(w, ErrMissingErrorDetails)
			return
		}

		var user CreateUserDTO
		err := json.NewDecoder(body).Decode(&user)

		newUser, createUserErr := userController.userService.CreateUser(user)

		var mySqlError *mysql.MySQLError
		if err != nil {
			NewBadRequestResponse(w, err)
		} else if errors.As(createUserErr, &mySqlError) {
			NewBadRequestResponse(w, ErrEmailTaken)
		} else {
			NewOkResponse(w, newUser)
		}
	}
}

func GetAllUsers(userController *Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userController.userService.GetAllUsers()

		if err != nil {
			NewInternalErrorResponse(w, err)
		} else {
			NewOkResponse(w, users)
		}
	}
}

func LoginHandler(uc *Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u AuthDTO
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			NewBadRequestResponse(w, err)
			return
		}
		log.Printf("The user request value %v", u)

		token, _, err := uc.userService.AuthenticateUser(u)

		if errors.Is(err, ErrUserNotFound) {
			NewNotFoundResponse(w, err)
			return
		} else if errors.Is(err, ErrInvalidCredentials) {
			NewUnauthorizedErrorResponse(w, err)
			return
		} else if errors.Is(err, ErrUserNotFound) {
			NewBadRequestResponse(w, err)
			return
		} else if err != nil {
			NewInternalErrorResponse(w, ErrNoUsernameFound)
			return
		}

		NewOkResponse(w, token)
		return
	}
}

func ProtectedHandler(uc *Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			NewUnauthorizedErrorResponse(w, ErrMissingAuthHeader)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := uc.userService.VerifyToken(tokenString)
		if err != nil {
			NewUnauthorizedErrorResponse(w, ErrInvalidToken)
			return
		}

		NewOkResponse(w, "Welcome to the the protected area")
	}
}
