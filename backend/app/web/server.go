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
)

func RunServer(db *sqlx.DB) func(port int) {
	return func(port int) {
		r := mux.NewRouter()

		routes(r, db)

		srv := &http.Server{
			Handler:      handlers.LoggingHandler(os.Stdout, r),
			Addr:         fmt.Sprintf(":%d", port),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())
	}
}
