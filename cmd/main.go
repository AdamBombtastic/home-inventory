package main

import (
	"fmt"

	"github.com/adambombtastic/home-inventory/pkg/core"
	"github.com/adambombtastic/home-inventory/pkg/stores/inventory"
	"github.com/adambombtastic/home-inventory/pkg/stores/requirements"
)

// First use case is to be able to answer the "requirements" of a "household".
// Each household has a list of items that they need to buy or restock.
// Each item has a name, a stock, and a last restock date.
// This first thing that the cli will answer is "What do I need to buy?"

// To answer this question, we need to be able to load a list of requirements.
// Then we need to be able to compare that list of requirements to a list of items.
// Requirements can be satisfied by multiple items.
// Items can satisfy multiple requirements.
func main() {

	reqService, err := requirements.New()
	if err != nil {
		panic(err)
	}

	inventoryService, err := inventory.New()
	if err != nil {
		panic(err)
	}

	app := core.NewApplication(reqService, inventoryService)

	// must
	reqs, _ := app.GetRequirements()

	// must
	items, _ := app.GetInventory()

	println("Requirements")
	println("------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s | %-20s |\n", "Name", "Quanity", "Units")
	for _, req := range reqs {
		fmt.Printf("| %-20s | %-10d | %-20s |\n", req.Name, req.Quantity, req.Units)
	}
	println("Inventory")
	println("------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s |\n", "Name", "Quanity")
	for _, item := range items {
		fmt.Printf("| %-20s | %-10d |\n", item.Name, item.Stock)
	}

	println("Checking Inventory")
	failures, err := app.InventoryMeetsRequirements()
	if err != nil {
		panic(err)
	}
	for _, failure := range failures {
		println("------------------------------------------------------------")
		fmt.Printf("Requirement: %s\n", failure.Requirement.Name)
		for _, reason := range failure.Reasons {
			fmt.Printf("Item: %s failed because : %s\n", reason.Item.Name, reason.Reason)
		}
	}
	if len(failures) == 0 {
		println("------------------------------------------------------------")
		println("All requirements met!")
		println("------------------------------------------------------------")
	}

}
