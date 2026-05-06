package handlers

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/service"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	DB          *sqlx.DB
	UserService *service.UserService
}

func NewUserHandler(db *sqlx.DB, us *service.UserService) *UserHandler {
	return &UserHandler{DB: db, UserService: us}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var p models.Register

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserService.Register(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var p models.Login
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.UserService.Login(&p, p.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
