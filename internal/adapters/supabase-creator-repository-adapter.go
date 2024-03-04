package adapters

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
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
	err := a.client.DB.From("creators").Select("*").Single().Like("email", email).Execute(&supabaseData)

	if err != nil {
		return nil, err
	} else if len(supabaseData) == 0 {
		return nil, nil
	}

	
	return entities.NewCreator(supabaseData["id"].(string), supabaseData["name"].(string), supabaseData["email"].(string)), nil
}

func (a *SupabaseCreatorRepositoryAdapter) FindById(id string) (*entities.Creator, error) {
	var supabaseData map[string]interface{}
	err := a.client.DB.From("creators").Select("*").Single().Eq("id", id).Execute(&supabaseData)

	if err != nil {
		return nil, err
	} else if len(supabaseData) == 0 {
		return nil, nil
	}

	return entities.NewCreator(supabaseData["id"].(string), supabaseData["name"].(string), supabaseData["email"].(string)), nil
}

func (a *SupabaseCreatorRepositoryAdapter) Create(input types.NewCreatorInput) error {

	var results []entities.Creator
	err := a.client.DB.From("creators").Insert(input).Execute(&results)

	return err
}
