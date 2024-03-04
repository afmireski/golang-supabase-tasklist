package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/afmireski/golang-supabase-tasklist/internal/services"
	"github.com/go-chi/chi/v5"
	myErrors "github.com/afmireski/golang-supabase-tasklist/internal/errors"
)

type TaskController struct {
	service *services.TasksService
}

func NewTaskController(service *services.TasksService) *TaskController {
	return &TaskController{service}
}

func (c *TaskController) GetById(w http.ResponseWriter, r *http.Request) {
	
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		err := myErrors.NewInternalError(err.Error(), http.StatusBadRequest)
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	output, serviceErr := c.service.FindById(int32(id))

	w.Header().Set("Content-Type", "application/json")
	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}