package routers

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupCreatorsRouter(router *chi.Mux, controller *controllers.CreatorsController) {
	router.Get("/creators/{id}", controller.GetById)

	router.Get("/creators/{email}/email", controller.GetByEmail)

	router.Post("/creators/new", controller.NewCreator)
}