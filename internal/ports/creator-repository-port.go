package ports

import "github.com/afmireski/golang-supabase-tasklist/internal/entities"

type CreateCreatorInput struct {
	Name  string
	Email string
}

type CreatorRepository interface {
	Create(input CreateCreatorInput) error
	FindById(id int32) (*entities.Creator, error)
	FindByEmail(email string) (*entities.Creator, error)
}