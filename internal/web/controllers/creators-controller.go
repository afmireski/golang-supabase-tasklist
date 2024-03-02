package controllers

import "github.com/afmireski/golang-supabase-tasklist/internal/services"

type CreatorsController struct {
	service *services.CreatorService
}

func NewCreatorsController(service *services.CreatorService) *CreatorsController {
	return &CreatorsController{service}
}

