package middleware

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vitorcsbrito/go-academy-todo/model"
	errors2 "github.com/vitorcsbrito/utils/errors"
	"github.com/vitorcsbrito/utils/requests"
	"log"
	"net/http"
	"os"
	"time"
)

func verifyToken(tokenString string) error {

	publicKey, err := os.ReadFile("./private/pub.key") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	b1, _ := pem.Decode(publicKey)
	keyPub, _ := x509.ParsePKIXPublicKey(b1.Bytes)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return keyPub, nil
	})

	if token != nil && !token.Valid {
		return errors2.ErrInvalidToken
	}

	exp, _ := token.Claims.GetExpirationTime()

	if !exp.After(time.Now()) {
		if !token.Valid {
			return errors2.ErrExpiredToken
		}
	}

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors2.ErrInvalidToken
	}

	return nil
}

func Auth(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		publicKey, err1 := os.ReadFile("./private/pub.key") // just pass the file name
		if err1 != nil {
			fmt.Print(err1)
		}

		b1, _ := pem.Decode(publicKey)
		keyPub, _ := x509.ParsePKIXPublicKey(b1.Bytes)

		log.Print("Executing AuthMiddleware")

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			requests.NewUnauthorizedErrorResponse(w, errors2.ErrMissingAuthHeader)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			requests.NewUnauthorizedErrorResponse(w, err)
			return
		}

		// Initialize a new instance of `Claims`
		claims := &model.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return keyPub, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Print("Finished executing AuthMiddleware")

		handler.ServeHTTP(w, r)
	})
}
