package services

import (
	"errors"
	"regexp"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
)

type CreatorService struct {
	repository ports.CreatorRepository
}

func NewCreatorService(repository ports.CreatorRepository) *CreatorService {
	return &CreatorService{repository}
}

func isValidEmail(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func isValidName(name string) bool {
	length := len(name);
	return 3 <= length && length <= 200
}

func (s *CreatorService) Create(name string, email string) error {
	if name == "" || email == "" {
		return errors.New("name and email are required")
	} else if !isValidEmail(email) {
		return errors.New("invalid email")
	} else if !isValidName(name) {
		return errors.New("the name should be between 3 and 200 characters")
	}

	creator := entities.NewCreator("", name, email)
	return s.repository.Create(creator)
}