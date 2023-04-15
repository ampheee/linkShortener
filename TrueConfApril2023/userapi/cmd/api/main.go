package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"refactoring/internal/client"
	service "refactoring/internal/service/repository"
)

const store = `../../users.json`

func main() {
	r := chi.NewRouter()
	s := service.NewService(r, store)
	client.InitMiddleWare(r)
	client.InitHandlers(r, s)
	http.ListenAndServe(":3333", r)
}
