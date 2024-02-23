package adapters

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseCreatorRepositoryAdapter struct {
	client *supabase.Client
}

func NewSupabaseCreatorRepositoryAdapter(client *supabase.Client) *SupabaseCreatorRepositoryAdapter {
	return &SupabaseCreatorRepositoryAdapter{
		client: client,
	}
}

func (a *SupabaseCreatorRepositoryAdapter) FindByEmail(email string) (*entities.Creator, error) {

	var supabaseData map[string]interface{}
	err := a.client.DB.From("creators").Select("*").Single().Eq("email", email).Execute(&supabaseData)

	if err != nil {
		return nil, err
	}

	return entities.NewCreator(supabaseData["id"].(string), supabaseData["name"].(string), supabaseData["email"].(string)), nil
}

func (a *SupabaseCreatorRepositoryAdapter) FindById(id string) (*entities.Creator, error) {
	var supabaseData map[string]interface{}
	err := a.client.DB.From("creators").Select("*").Single().Eq("id", id).Execute(&supabaseData)

	if err != nil {
		return nil, err
	}

	return entities.NewCreator(supabaseData["id"].(string), supabaseData["name"].(string), supabaseData["email"].(string)), nil
}

func (a *SupabaseCreatorRepositoryAdapter) Create(input entities.Creator) error {

	var results []entities.Creator
	err := a.client.DB.From("creators").Insert(input).Execute(&results)

	return err
}
