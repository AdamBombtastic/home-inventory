package main

import (
	"fmt"

	"github.com/adambombtastic/home-inventory/pkg/interactors"
	"github.com/adambombtastic/home-inventory/pkg/inventory"
	"github.com/adambombtastic/home-inventory/pkg/requirements"
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
	reqStore, err := requirements.NewStore()
	if err != nil {
		panic(err)
	}
	reqs, err := reqStore.All()
	if err != nil {
		panic(err)
	}
	println("Requirements")
	println("------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s | %-20s |\n", "Name", "Quanity", "Units")
	for _, req := range reqs {
		fmt.Printf("| %-20s | %-10d | %-20s |\n", req.Name, req.Quantity, req.Units)
	}
	inv, err := inventory.NewStore()
	if err != nil {
		panic(err)
	}
	items, err := inv.All()
	if err != nil {
		panic(err)
	}
	println("Inventory")
	println("------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s |\n", "Name", "Quanity")
	for _, item := range items {
		fmt.Printf("| %-20s | %-10d |\n", item.Name, item.Stock)
	}

	println("Checking Inventory")
	failures := interactors.InventoryMeetsRequirements(items, reqs)
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
