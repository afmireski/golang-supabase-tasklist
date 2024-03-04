package services

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	myErrors "github.com/afmireski/golang-supabase-tasklist/internal/errors"
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
	"github.com/afmireski/golang-supabase-tasklist/internal/validators"
)

type TasksService struct {
	repository ports.TaskRepository
}

func NewTasksService(repository ports.TaskRepository) *TasksService {
	return &TasksService{repository}
}

func (s *TasksService) FindById(id int32) (*entities.Task, *myErrors.InternalError) {
	if !validators.IsValidNumericId(id) {
		return nil, myErrors.NewInternalError("invalid id", 400)
	}

	repositoryResponse, err := s.repository.FindById(id)

	if err != nil {
		return nil, myErrors.NewInternalError(err.Error(), 500)
	} else if repositoryResponse == nil {
		return nil, myErrors.NewInternalError("Task not found", 404)
	}
	return repositoryResponse, nil
}
