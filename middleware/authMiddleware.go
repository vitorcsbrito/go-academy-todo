package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vitorcsbrito/utils/errors"
	"log"
	"net/http"
)

var secretKey = []byte("c2VjcmV0")

func verifyToken(tokenString string) error {
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

func Auth(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing AuthMiddleware")

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)

			fmt.Fprint(w, errors.ErrMissingAuthHeader)
			log.Print(errors.ErrMissingAuthHeader)

			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		log.Print("Finished executing AuthMiddleware")

		handler.ServeHTTP(w, r)
	})
}
