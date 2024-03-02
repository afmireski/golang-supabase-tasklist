package ports

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
)

type TaskRepository interface {
	Create(input types.CreateTaskInput) error
	FindById(id int32) (*entities.Task, error)
	FindByTitle(title string, creatorId string) ([]*entities.Task, error)
	FindAll(creatorId string) ([]*entities.Task, error)
}
