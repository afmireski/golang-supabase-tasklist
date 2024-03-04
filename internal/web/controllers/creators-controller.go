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

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *CreatorsController) GetByEmail(w http.ResponseWriter, r *http.Request) {

	emailParam := chi.URLParam(r, "email")

	output, err := c.service.FindByEmail(emailParam)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *CreatorsController) NewCreator(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Fail on try read the request body"))
		return
	}

	serviceErr := c.service.Create(input.Name, input.Email)

	w.Header().Set("Content-Type", "application/json")
	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

