package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	. "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/service"
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
	mux.HandleFunc("POST /protected", ProtectedHandler(userController))
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
		w.Header().Set("Content-Type", "application/json")

		var u AuthDTO
		json.NewDecoder(r.Body).Decode(&u)
		log.Printf("The user request value %v", u)

		if u.Username == "Chek" && u.Password == "123456" {
			tokenString, err := uc.userService.CreateToken(u.Username)
			if err != nil {
				NewInternalErrorResponse(w, ErrNoUsernameFound)
			}

			NewOkResponse(w, tokenString)

			return
		} else {
			NewUnauthorizedErrorResponse(w, ErrInvalidCredentials)
		}
	}
}

func ProtectedHandler(uc *Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := uc.userService.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		fmt.Fprint(w, "Welcome to the the protected area")
	}
}
