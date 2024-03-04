package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afmireski/golang-supabase-tasklist/internal/adapters"
	"github.com/afmireski/golang-supabase-tasklist/internal/services"
	"github.com/afmireski/golang-supabase-tasklist/internal/web/controllers"
	"github.com/afmireski/golang-supabase-tasklist/internal/web/routers"
	"github.com/go-chi/chi/v5"
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

	creatorsRepository := adapters.NewSupabaseCreatorRepositoryAdapter(supabaseClient)
	creatorsService := services.NewCreatorService(creatorsRepository)
	creatorsController := controllers.NewCreatorsController(creatorsService)

	r := chi.NewRouter()
	routers.SetupCreatorsRouter(r, creatorsController)

	fmt.Println("Running server...")
	http.ListenAndServe(":3000", r)
}