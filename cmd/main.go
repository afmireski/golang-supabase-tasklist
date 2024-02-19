package main

import (
	"fmt"
	"log"
	"os"

	"github.com/afmireski/golang-supabase-tasklist/internal/adapters"
	"github.com/joho/godotenv"
	supabase "github.com/nedpals/supabase-go"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")

	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

	creatorRepository := adapters.NewSupabaseCreatorRepositoryAdapter(supabaseClient)

	creator, err := creatorRepository.FindByEmail("teste@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Id: ", creator.Id)
	fmt.Println("Name: ", creator.Name)
	fmt.Println("Email: ", creator.Email)
}