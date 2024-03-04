package services

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
	"github.com/google/uuid"
)

type TasksService struct {
	repository ports.TaskRepository
}

func NewTasksService(repository ports.TaskRepository) *TasksService {
	return &TasksService{repository}
}

func isValidUuid(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}



