package api

import (
	"tasks/datastore"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	TaskHandler *TaskHandler
	UserHandler *UserHandler
	router      chi.Router
}

func (h *Handlers) userRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.UserHandler.Create)

	return r
}

func (h *Handlers) taskRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.TaskHandler.GetAll)
	r.Get("/{id}", h.TaskHandler.Get)
	r.Post("/", h.TaskHandler.Create)
	r.Delete("/", h.TaskHandler.Delete)

	return r
}

func (h *Handlers) RegisterRoutes() chi.Router {
	h.router.Mount("/tasks", h.taskRouter())
	h.router.Mount("/users", h.userRouter())

	return h.router
}

func NewHandlers(db *datastore.Datastore) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(db.UserDatastore),
		TaskHandler: NewTaskHandler(db.TaskDatastore),
		router:      chi.NewRouter(),
	}
}
