package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"refactoring/internal/service"
	"time"
)

func InitMiddleWare(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func InitHandlers(r *chi.Mux, service service.Service) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", service.SearchUsers)
		r.Post("/", service.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", service.GetUser)
			r.Patch("/", service.UpdateUser)
			r.Delete("/", service.DeleteUser)
		})
	})
}
