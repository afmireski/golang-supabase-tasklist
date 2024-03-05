package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"github.com/afmireski/golang-supabase-tasklist/internal/adapters"
	"github.com/afmireski/golang-supabase-tasklist/internal/services"
	"github.com/afmireski/golang-supabase-tasklist/internal/web/controllers"
	"github.com/afmireski/golang-supabase-tasklist/internal/web/routers"
	"github.com/go-chi/chi/v5"
	supabase "github.com/nedpals/supabase-go"
)

func main() {

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")

	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

	creatorsRepository := adapters.NewSupabaseCreatorRepositoryAdapter(supabaseClient)
	creatorsService := services.NewCreatorService(creatorsRepository)
	creatorsController := controllers.NewCreatorsController(creatorsService)

	taskRepository := adapters.NewSupabaseTasksRepositoryAdapter(supabaseClient)
	taskService := services.NewTasksService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	r := chi.NewRouter()
	routers.SetupCreatorsRouter(r, creatorsController)
	routers.SetupTasksRouter(r, taskController)

	fmt.Println("Running server...")
	http.ListenAndServe(":3000", r)
}