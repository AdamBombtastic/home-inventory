package core

import (
	"github.com/adambombtastic/home-inventory/pkg/core/entities"
	"github.com/adambombtastic/home-inventory/pkg/core/interactors"
	"github.com/adambombtastic/home-inventory/pkg/core/services"
)

type Application interface {
	GetRequirements() ([]*entities.Requirement, error)
	GetInventory() ([]*entities.Item, error)
	// RequirementsWithReasons should be moved to entities
	InventoryMeetsRequirements() ([]*interactors.RequirementWithReasons, error)
}

type app struct {
	reqService services.RequirementService
	invService services.InventoryService
}

func NewApplication(reqService services.RequirementService, invService services.InventoryService) Application {
	return &app{
		reqService: reqService,
		invService: invService,
	}
}

func (a *app) GetRequirements() ([]*entities.Requirement, error) {
	return a.reqService.All()
}
func (a *app) GetInventory() ([]*entities.Item, error) {
	return a.invService.All()
}
func (a *app) InventoryMeetsRequirements() ([]*interactors.RequirementWithReasons, error) {
	reqs, err := a.reqService.All()
	if err != nil {
		return nil, err
	}
	inventory, _ := a.invService.All()
	if err != nil {
		return nil, err
	}
	return interactors.InventoryMeetsRequirements(inventory, reqs), nil
}
