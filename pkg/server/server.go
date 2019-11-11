package server

import (
	"log"
	"net/http"
	"net/http/cgi"

	"CASystem/pkg/DB"
	"CASystem/pkg/di"
	"CASystem/pkg/server/handler"

	"github.com/go-chi/chi"
)

func Serve() {
	/* ===== Initalize ===== */
	log.Println("Initialize...")
	DB.Init()
	di.Init()

	/* ===== Routing ===== */
	log.Println("Routing...")

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Route("/user", func(r chi.Router) {
		r.Get("/", get(handler.UserGet()))
	})

	/* ===== Run CGI Server ===== */
	log.Println("Server running...")
	err := cgi.Serve(r)
	if err != nil {
		log.Fatal(err)
	}
}
