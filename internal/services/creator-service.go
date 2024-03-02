package services

import (
	"errors"
	"regexp"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
	"github.com/google/uuid"
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
	length := len(name)
	return 3 <= length && length <= 200
}

func isValidUuid(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
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

func (s *CreatorService) FindById(id string) (*entities.Creator, error) {

	if !isValidUuid(id) {
		return nil, errors.New("invalid id")
	}

	return s.repository.FindById(id)
}