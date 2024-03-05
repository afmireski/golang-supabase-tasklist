package routers

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupTasksRouter(router *chi.Mux, controller *controllers.TaskController) {
	router.Get("/tasks/{id}", controller.GetById)
}
