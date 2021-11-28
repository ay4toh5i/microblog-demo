package web

import (
	"errors"
	"log"
	"net/http"
	"microblog/microblog/domain"
)

func handleError(w http.ResponseWriter, err error) {
	switch true {
	case errors.Is(err, domain.ErrInvalidRequest):
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	case errors.Is(err, domain.ErrAuthentication):
		http.Error(w, err.Error(), http.StatusUnauthorized)

	case errors.Is(err, domain.ErrNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	default:
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
