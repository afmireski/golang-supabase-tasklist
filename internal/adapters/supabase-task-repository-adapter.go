package adapters

import (
	"time"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseTaskRepositoryAdapter struct {
	client *supabase.Client
}

func (a *SupabaseTaskRepositoryAdapter) FindById(id int32) (*entities.Task, error) {
	var supabaseData entities.Task

	err := a.client.DB.From("tasks").Select("*").Single().Eq("id", string(id)).Execute(&supabaseData)

	if (err != nil) {
		return nil, err
	}

	return &supabaseData, nil
}

// func (a *SupabaseTaskRepositoryAdapter) FindByTitle(title string, creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) FindAll(creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) Create(input *entities.Task) error