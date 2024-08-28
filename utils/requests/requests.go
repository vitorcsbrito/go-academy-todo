package requests

import (
	"encoding/json"
	"github.com/vitorcsbrito/utils/errors"
	"net/http"
)

func SetOkRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SetBadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
}

func SetNotFoundRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func SetInternalErrorRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}

func NewNotFoundResponse(w http.ResponseWriter, err error) {
	SetNotFoundRequest(w)
	json.NewEncoder(w).Encode(errors.NewErrResponse(err))
}

func NewOkResponse(w http.ResponseWriter, err any) {
	SetOkRequest(w)
	json.NewEncoder(w).Encode(err)
}

func NewBadRequestResponse(w http.ResponseWriter, err error) {
	SetBadRequest(w)
	json.NewEncoder(w).Encode(errors.NewErrResponse(err))
}

func NewInternalErrorResponse(w http.ResponseWriter, err error) {
	SetInternalErrorRequest(w)
	json.NewEncoder(w).Encode(errors.NewErrResponse(err))
}
