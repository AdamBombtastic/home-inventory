package services

import "github.com/adambombtastic/home-inventory/pkg/core/entities"

type RequirementService interface {
	All() ([]*entities.Requirement, error)
}

type MockRequirementService struct {
	Reqs []*entities.Requirement
	err  error
}

func (m *MockRequirementService) All() ([]*entities.Requirement, error) {
	return m.Reqs, m.err
}
