package service

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/fs"
	"net/http"
	"os"
	"refactoring/internal/errs"
	"refactoring/internal/requests"
	"refactoring/internal/service"
	"refactoring/internal/user"
	"strconv"
	"time"
)

type ServiceStruct struct {
	Router *chi.Mux
	Store  string
}

var (
	UserNotFound = errors.New("user_not_found")
)

func (service *ServiceStruct) SearchUsers(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(service.Store)
	s := user.UserStore{}
	_ = json.Unmarshal(f, &s)
	render.JSON(w, r, s.List)
}

func (service *ServiceStruct) GetUser(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(service.Store)
	s := user.UserStore{}
	_ = json.Unmarshal(f, &s)
	id := chi.URLParam(r, "id")
	render.JSON(w, r, s.List[id])
}

func (service *ServiceStruct) UpdateUser(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(service.Store)
	s := user.UserStore{}
	_ = json.Unmarshal(f, &s)
	request := requests.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, errs.ErrInvalidRequest(err))
		return
	}
	id := chi.URLParam(r, "id")
	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, errs.ErrInvalidRequest(UserNotFound))
		return
	}
	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u
	b, _ := json.Marshal(&s)
	_ = os.WriteFile(service.Store, b, fs.ModePerm)
	render.Status(r, http.StatusNoContent)
}

func (service *ServiceStruct) DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(service.Store)
	s := user.UserStore{}
	_ = json.Unmarshal(f, &s)
	id := chi.URLParam(r, "id")
	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, errs.ErrInvalidRequest(UserNotFound))
		return
	}
	delete(s.List, id)
	b, _ := json.Marshal(&s)
	_ = os.WriteFile(service.Store, b, fs.ModePerm)
	render.Status(r, http.StatusNoContent)
}

func (service *ServiceStruct) CreateUser(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(service.Store)
	s := user.UserStore{}
	_ = json.Unmarshal(f, &s)
	request := requests.CreateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, errs.ErrInvalidRequest(err))
		return
	}
	s.Increment++
	u := user.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}
	id := strconv.Itoa(s.Increment)
	s.List[id] = u
	b, _ := json.Marshal(&s)
	_ = os.WriteFile(service.Store, b, fs.ModePerm)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func NewService(r *chi.Mux, store string) service.Service {
	return &ServiceStruct{r, store}
}
