package handlers

import (
	"github.com/evsyukovmv/taskmanager/handlers/columns"
	"github.com/evsyukovmv/taskmanager/handlers/comments"
	appMiddleware "github.com/evsyukovmv/taskmanager/handlers/middleware"
	"github.com/evsyukovmv/taskmanager/handlers/projects"
	"github.com/evsyukovmv/taskmanager/handlers/tasks"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(appMiddleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/projects", projectsRouter())
	return r
}

func projectsRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", projects.GetList)
	r.Post("/", projects.Create)
	r.Get("/{projectId}", projects.GetById)
	r.Put("/{projectId}", projects.Update)
	r.Delete("/{projectId}", projects.Delete)

	r.Mount("/{projectId}/columns", columnsRouter())
	return r
}

func columnsRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", columns.GetList)
	r.Post("/", columns.Create)
	r.Get("/{columnId}", columns.GetById)
	r.Put("/{columnId}", columns.Update)
	r.Put("/{columnId}/move", columns.Move)
	r.Delete("/{columnId}", columns.Delete)

	r.Mount("/{columnId}/tasks", tasksRouter())
	return r
}

func tasksRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", tasks.GetList)
	r.Post("/", tasks.Create)
	r.Get("/{taskId}", tasks.GetById)
	r.Put("/{taskId}", tasks.Update)
	r.Put("/{taskId}/move", tasks.Move)
	r.Delete("/{taskId}", tasks.Delete)

	r.Mount("/{taskId}/comments", commentsRouter())
	return r
}

func commentsRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", comments.GetList)
	r.Post("/", comments.Create)
	r.Get("/{commentId}", comments.GetById)
	r.Put("/{commentId}", comments.Update)
	r.Delete("/{commentId}", comments.Delete)

	return r
}
