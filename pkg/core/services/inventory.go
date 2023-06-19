package services

import "github.com/adambombtastic/home-inventory/pkg/core/entities"

type InventoryService interface {
	All() ([]*entities.Item, error)
}

type MockInventoryService struct {
	Items []*entities.Item
	err   error
}

func (m *MockInventoryService) All() ([]*entities.Item, error) {
	return m.Items, m.err
}
