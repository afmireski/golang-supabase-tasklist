package ports

import "github.com/afmireski/golang-supabase-tasklist/internal/entities"

type TaskRepository interface {
	Create(input *entities.Task) error
	FindById(id int32) (*entities.Task, error)
	FindByTitle(title string, creatorId int32) ([]entities.Task, error)
	FindAll(creatorId int32) ([]entities.Task, error)
}