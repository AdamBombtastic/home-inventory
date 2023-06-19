package core

import (
	"testing"

	"github.com/adambombtastic/home-inventory/pkg/core/entities"
	"github.com/adambombtastic/home-inventory/pkg/core/services"
)

func TestNewApplication(t *testing.T) {
	got := NewApplication(&services.MockRequirementService{}, &services.MockInventoryService{})
	if got == nil {
		t.Errorf("NewApplication() = %v, want %v", got, nil)
	}
}

func Test_AllReqs_Happy(t *testing.T) {
	app := NewApplication(&services.MockRequirementService{
		Reqs: []*entities.Requirement{
			{
				Name: "Test Requirement",
			},
		},
	}, &services.MockInventoryService{})
	got, err := app.GetRequirements()
	if err != nil {
		t.Errorf("GetRequirements() = %v, want %v", got, nil)
	}
	if len(got) != 1 {
		t.Errorf("GetRequirements() = %v, want %v", len(got), 1)
	}
	if got[0].Name != "Test Requirement" {
		t.Errorf("GetRequirements() = %v, want %v", got[0].Name, "Test Requirement")
	}
}

func Test_AllItems_Happy(t *testing.T) {
	app := NewApplication(&services.MockRequirementService{}, &services.MockInventoryService{
		Items: []*entities.Item{
			{
				Name: "Test Item",
			},
		},
	})
	got, err := app.GetInventory()
	if err != nil {
		t.Errorf("GetInventory() = %v, want %v", got, nil)
	}
	if len(got) != 1 {
		t.Errorf("GetInventory() = %v, want %v", len(got), 1)
	}
	if got[0].Name != "Test Item" {
		t.Errorf("GetInventory() = %v, want %v", got[0].Name, "Test Item")
	}
}
