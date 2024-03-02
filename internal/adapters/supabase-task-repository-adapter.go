package adapters

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseTaskRepositoryAdapter struct {
	client *supabase.Client
}

type supabaseSelectAllTaskResponse struct {
	Id          int32             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Date        time.Time         `json:"date"`
	Finished    bool              `json:"completed"`
	Creators    *entities.Creator `json:"creators"`
}

func NewSupabaseTasksRepositoryAdapter(client *supabase.Client) *SupabaseTaskRepositoryAdapter {
	return &SupabaseTaskRepositoryAdapter{
		client: client,
	}
}

func serializeSupabaseData(data map[string]interface{}) (*entities.Task, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	var temp supabaseSelectAllTaskResponse
	json.Unmarshal(jsonData, &temp)

	if err != nil {
		return nil, err
	}

	return entities.NewTask(temp.Id, temp.Title, temp.Description, temp.Date, temp.Finished, temp.Creators), nil
}

func (a *SupabaseTaskRepositoryAdapter) FindById(id int32) (*entities.Task, error) {
	var supabaseData map[string]interface{}

	parsedId := strconv.Itoa(int(id))
	err := a.client.DB.From("tasks").Select("*", "creators (*)").Single().Eq("id", parsedId).Execute(&supabaseData)

	if err != nil {
		return nil, err
	}

	response, err := serializeSupabaseData(supabaseData)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// func (a *SupabaseTaskRepositoryAdapter) FindByTitle(title string, creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) FindAll(creatorId int32) ([]entities.Task, error)

// func (a *SupabaseTaskRepositoryAdapter) Create(input *entities.Task) error
