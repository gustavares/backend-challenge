package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tasks/datastore"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	db datastore.UserDatastore
}

func NewUserHandler(db datastore.UserDatastore) *UserHandler {
	return &UserHandler{
		db,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TODO: implement GET /tasks/:id
	response := "Getting task: " + id
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: implement GET /tasks filter by user
	json.NewEncoder(w).Encode("Getting all tasks")

}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: implement POST
	var user datastore.UserEntity

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorMessage := fmt.Sprintf("error decoding JSON: %v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
	}

	user.ID = uuid.New()
	createErr := h.db.CreateUser(user)
	if createErr != nil {
		errorMessage := fmt.Sprintf("error decoding JSON: %v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
	}
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: implement DELETE
}
