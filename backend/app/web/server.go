package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

var FRONTEND_ORIGIN = os.Getenv("FRONTEND_ORIGIN")

func RunServer(db *sqlx.DB) func(port int) {
	return func(port int) {
		r := mux.NewRouter()

		routes(r, db)

		handler := cors.New(cors.Options{
			AllowedOrigins:   []string{FRONTEND_ORIGIN},
			AllowCredentials: true,
		}).Handler(r)

		srv := &http.Server{
			Handler:      handlers.LoggingHandler(os.Stdout, handler),
			Addr:         fmt.Sprintf(":%d", port),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())
	}
}
