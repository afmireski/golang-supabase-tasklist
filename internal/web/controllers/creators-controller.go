package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/golang-supabase-tasklist/internal/services"
	"github.com/go-chi/chi/v5"
)

type CreatorsController struct {
	service *services.CreatorService
}

func NewCreatorsController(service *services.CreatorService) *CreatorsController {
	return &CreatorsController{service}
}

func (c *CreatorsController) GetById(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	output, err := c.service.FindById(idParam)

	if err != nil {
		w.WriteHeader(err.HttpCode())
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

