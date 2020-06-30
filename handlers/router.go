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
	r.Use(middleware.RequestID)
	r.Use(appMiddleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/projects", projectsRouter())
	r.Mount("/columns", columnsItemRouter())
	r.Mount("/tasks", tasksItemRouter())
	r.Mount("/comments", commentsItemRouter())
	return r
}

func projectsRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", projects.GetList)
	r.Post("/", projects.Create)
	r.Get("/{projectId}", projects.GetById)
	r.Put("/{projectId}", projects.Update)
	r.Delete("/{projectId}", projects.Delete)

	r.Mount("/{projectId}/columns", columnsCollectionRouter())
	return r
}

func columnsCollectionRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", columns.GetList)
	r.Post("/", columns.Create)
	return r
}

func columnsItemRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{columnId}", columns.GetById)
	r.Put("/{columnId}", columns.Update)
	r.Put("/{columnId}/move", columns.Move)
	r.Delete("/{columnId}", columns.Delete)

	r.Mount("/{columnId}/tasks", tasksCollectionRouter())
	return r
}

func tasksCollectionRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", tasks.GetList)
	r.Post("/", tasks.Create)
	return r
}

func tasksItemRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{taskId}", tasks.GetById)
	r.Put("/{taskId}", tasks.Update)
	r.Put("/{taskId}/move", tasks.Move)
	r.Put("/{taskId}/shift", tasks.Shift)
	r.Delete("/{taskId}", tasks.Delete)
	r.Mount("/{taskId}/comments", commentsCollectionRouter())
	return r
}

func commentsCollectionRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", comments.GetList)
	r.Post("/", comments.Create)
	r.Get("/{commentId}", comments.GetById)
	r.Put("/{commentId}", comments.Update)
	r.Delete("/{commentId}", comments.Delete)
	return r
}

func commentsItemRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{commentId}", comments.GetById)
	r.Put("/{commentId}", comments.Update)
	r.Delete("/{commentId}", comments.Delete)
	return r
}
