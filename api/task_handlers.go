package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tasks/datastore"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TaskHandler struct {
	db datastore.TaskDatastore
}

func NewTaskHandler(db datastore.TaskDatastore) *TaskHandler {
	return &TaskHandler{
		db,
	}
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TODO: implement GET /tasks/:id
	response := "Getting task: " + id
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	// TODO: validate if request is from an Admin or the same user technician requesting the tasks
	userId := chi.URLParam(r, "id")

	tasks, err := h.db.GetTasksFromUser(userId)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to fetch tasks from user %s, reason: %s", userId, err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: implement GET /tasks filter by user
	json.NewEncoder(w).Encode("Getting all tasks")

}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task datastore.TaskEntity

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		errorMessage := fmt.Sprintf("error decoding JSON: %v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
	}

	task.ID = uuid.New()
	createErr := h.db.CreateTask(task)
	if createErr != nil {
		errorMessage := fmt.Sprintf("error decoding JSON: %v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
	}
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: implement DELETE
}
