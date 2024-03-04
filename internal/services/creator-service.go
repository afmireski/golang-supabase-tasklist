package services

import (
	"github.com/afmireski/golang-supabase-tasklist/internal/entities"
	myErrors "github.com/afmireski/golang-supabase-tasklist/internal/errors"
	"github.com/afmireski/golang-supabase-tasklist/internal/ports"
	"github.com/afmireski/golang-supabase-tasklist/internal/types"
	"github.com/afmireski/golang-supabase-tasklist/internal/validators"
)

type CreatorService struct {
	repository ports.CreatorRepository
}

func NewCreatorService(repository ports.CreatorRepository) *CreatorService {
	return &CreatorService{repository}
}

func (s *CreatorService) Create(name string, email string) *myErrors.InternalError {
	if name == "" || email == "" {
		return myErrors.NewInternalError("missing name or email", 400)
	} else if !validators.IsValidEmail(email) {
		return myErrors.NewInternalError("invalid email", 400)
	} else if !validators.IsValidName(name) {
		return myErrors.NewInternalError("invalid name", 400)
	}

	err := s.repository.Create(types.NewCreatorInput{name, email})

	if err != nil {
		return myErrors.NewInternalError(err.Error(), 500)
	}

	return nil
}

func (s *CreatorService) FindById(id string) (*entities.Creator, *myErrors.InternalError) {

	if !validators.IsValidUuid(id) {
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
	if !validators.IsValidEmail(email) {
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
