package ports

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
)

type CreatorRepository interface {
	Create(input types.NewCreatorInput) error
	FindById(id string) (*entities.Creator, error)
	FindByEmail(email string) (*entities.Creator, error)
}
