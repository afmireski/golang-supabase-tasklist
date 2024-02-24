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

	// creatorRepository := adapters.NewSupabaseCreatorRepositoryAdapter(supabaseClient)
	tasksRepository := adapters.NewSupabaseTasksRepositoryAdapter(supabaseClient)

	task, err := tasksRepository.FindById(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task)	
}