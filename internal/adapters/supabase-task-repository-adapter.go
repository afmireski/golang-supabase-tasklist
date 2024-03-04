package adapters

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseTaskRepositoryAdapter struct {
	client *supabase.Client
}

type supabaseTaskResponse struct {
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

	var temp supabaseTaskResponse
	json.Unmarshal(jsonData, &temp)

	if err != nil {
		return nil, err
	}

	return entities.NewTask(temp.Id, temp.Title, temp.Description, temp.Date, temp.Finished, temp.Creators), nil
}

func serializeSupabaseDataArray(data interface{}) ([]*entities.Task, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	var temp []supabaseTaskResponse
	json.Unmarshal(jsonData, &temp)

	if err != nil {
		return nil, err
	}

	var tasks []*entities.Task
	for _, task := range temp {
		tasks = append(tasks, entities.NewTask(task.Id, task.Title, task.Description, task.Date, task.Finished, task.Creators))
	}

	return tasks, nil
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
	} else if response == nil {
		return nil, nil
	}

	return response, nil
}

func (a *SupabaseTaskRepositoryAdapter) FindByTitle(title string, creatorId string) ([]*entities.Task, error) {
	var supabaseData []map[string]interface{}

	err := a.client.DB.From("tasks").Select("*", "creators (*)").Like("title", "%" + title + "%").Eq("creator_id", creatorId).Execute(&supabaseData)

	if err != nil {
		return nil, err
	}

	return serializeSupabaseDataArray(supabaseData)
}

func (a *SupabaseTaskRepositoryAdapter) FindAll(creatorId string) ([]*entities.Task, error) {
	var supabaseData []map[string]interface{}

	err := a.client.DB.From("tasks").Select("*", "creators (*)").Eq("creator_id", creatorId).Execute(&supabaseData)

	if err != nil {
		return nil, err
	}

	return serializeSupabaseDataArray(supabaseData)
}

func (a *SupabaseTaskRepositoryAdapter) Create(input types.CreateTaskInput) error {
	var supabaseData []supabaseTaskResponse;

	err := a.client.DB.From("tasks").Insert(input).Execute(&supabaseData);

	return err
}
