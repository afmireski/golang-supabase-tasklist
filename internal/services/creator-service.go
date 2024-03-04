package services

import (
	"regexp"

	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	myErrors "github.com/afmireski/golang-supabase-tasklist/internal/errors"
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
	"github.com/google/uuid"
)

type CreatorService struct {
	repository ports.CreatorRepository
}

func NewCreatorService(repository ports.CreatorRepository) *CreatorService {
	return &CreatorService{repository}
}

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}
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

func (s *CreatorService) Create(name string, email string) *myErrors.InternalError {
	if name == "" || email == "" {
		return myErrors.NewInternalError("missing name or email", 400)
	} else if !isValidEmail(email) {
		return myErrors.NewInternalError("invalid email", 400)
	} else if !isValidName(name) {
		return myErrors.NewInternalError("invalid name", 400)
	}

	err := s.repository.Create(types.NewCreatorInput{name, email})

	if err != nil {
		return myErrors.NewInternalError(err.Error(), 500)
	}

	return nil
}

func (s *CreatorService) FindById(id string) (*entities.Creator, *myErrors.InternalError) {

	if !isValidUuid(id) {
		return nil, myErrors.NewInternalError("invalid id", 400)
	}

	response, err := s.repository.FindById(id)

	if err != nil {
		return nil, myErrors.NewInternalError(err.Error(), 500)
	} else if response == nil {
		return nil, myErrors.NewInternalError("Creator not found", 404)
	}
	return response, nil
}

func (s *CreatorService) FindByEmail(email string) (*entities.Creator, *myErrors.InternalError) {
	if !isValidEmail(email) {
		return nil, myErrors.NewInternalError("invalid email", 400)
	}

	response, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, myErrors.NewInternalError(err.Error(), 500)
	} else if response == nil {
		return nil, myErrors.NewInternalError("Creator not found", 404)
	}
	return response, nil
}
