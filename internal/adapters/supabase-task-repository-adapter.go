package adapters

import (
	"encoding/json"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseTaskRepositoryAdapter struct {
	client *supabase.Client
}

func NewSupabaseTasksRepositoryAdapter(client *supabase.Client) *SupabaseTaskRepositoryAdapter {
	return &SupabaseTaskRepositoryAdapter{
		client: client,
	}
}

func serializeSupabaseData(data map[string]interface{}) (*entities.Task, error) {
	var task entities.Task
	taskData, err := json.Marshal(data)

	if (err != nil) {
		return nil, err
	}

	var creator entities.Creator
	creatorData, err := json.Marshal(data["creators"])

	if (err != nil) {
		return nil, err
	}

	err = json.Unmarshal(creatorData, &creator)

	if (err != nil) {
		return nil, err
	}

	err = json.Unmarshal(taskData, &task)

	if (err != nil) {
		return nil, err
	}

	task.Creator = &creator;

	return &task, nil
}

func (a *SupabaseTaskRepositoryAdapter) FindById(id int32) (*entities.Task, error) {
	var supabaseData map[string]interface{}
	err := a.client.DB.From("tasks").Select("*", "creators (*)").Single().Eq("id", iId).Execute(&supabaseData)

	if (err != nil) {
		return nil, err
	}

	response, err := serializeSupabaseData(supabaseData)

	if (err != nil) {
		return nil, err
	}

	return response, nil
}

// func (a *SupabaseTaskRepositoryAdapter) FindByTitle(title string, creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) FindAll(creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) Create(input *entities.Task) error