package ports

import "github.com/afmireski/golang-supabase-tasklist/internal/entities"

type CreatorRepository interface {
	Create(input entities.Creator) error
	FindById(id string) (*entities.Creator, error)
	FindByEmail(email string) (*entities.Creator, error)
}
