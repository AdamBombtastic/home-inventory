package interactors

import (
	"time"

	"github.com/adambombtastic/home-inventory/pkg/inventory"
	"github.com/adambombtastic/home-inventory/pkg/requirements"
)

/*
	TODO: Currently, requirements can only be evaluated against a single item, in the future, we'll want to be able to evaluate multiple items.
	For example: I need 5 'Snack' items. That could be 3 bags of chips and 2 boxes of cookies. Currently, we can only evaluate the group of chips or the group of cookies.
*/

// RequirementWithReasons is a struct that contains a requirement and the reasons why it failed to meet the requirement
// Need to add a reference to the item that failed to meet the requirement so that we can display the item in the UI.
type RequirementWithReasons struct {
	Requirement *requirements.Requirement
	Reasons     []*CheckInfo
}
type CheckInfo struct {
	Item   *inventory.Item
	Reason requirements.Reason
}

// InventoryMeetsRequirements is a function that checks the inventory against the requirements
// In the event that the inventory is missing items, it returns a list of requirements that are missing with the reasons why each
// item failed to meet the requirement
func InventoryMeetsRequirements(inventory []*inventory.Item, reqs []*requirements.Requirement) []*RequirementWithReasons {
	// map of requirement ID to whether or not it's met
	runMap := make(map[int][]*CheckInfo)
	metMap := make(map[int]bool)

	for _, req := range reqs {
		met, reasons := ItemsMeetRequirement(req, inventory)
		metMap[req.ID] = met
		runMap[req.ID] = reasons
	}
	unmetRequirements := []*RequirementWithReasons{}
	for _, req := range reqs {
		if !metMap[req.ID] {
			unmetRequirements = append(unmetRequirements, &RequirementWithReasons{req, runMap[req.ID]})
		}
	}

	return unmetRequirements
}

// ItemsMeetsRequirement is a function that checks a list of items against a requirement, returning whether or not the items meet the requirement
func ItemsMeetRequirement(req *requirements.Requirement, items []*inventory.Item) (bool, []*CheckInfo) {
	reasons := []*CheckInfo{}

	// No requirement, no problem
	if req == nil {
		return true, reasons
	}
	// inStock is the total number of items that meet the requirement, not the individual item quantity
	inStock := 0

	for _, item := range items {

		meets, reason := itemMeetsRequirement(req, item)
		if meets {
			inStock += item.Stock
			continue
		}
		reasons = append(reasons, &CheckInfo{item, reason})
	}

	return inStock >= req.Quantity, reasons
}

// itemMeetsRequirement is a function that checks a single item against a requirement, returning whether or not the item meets the requirement
func itemMeetsRequirement(req *requirements.Requirement, item *inventory.Item) (bool, requirements.Reason) {
	// No requirement, no problem
	if req == nil {
		return true, requirements.ReasonNone
	}

	// Doesn't apply to this item, so the requirement is not met
	applies := false
	for _, tag := range req.ApplicableItems {
		if tag == item.Name {
			applies = true
			break
		}
	}
	if !applies {
		return false, requirements.ReasonNotApplicable
	}

	// If the item is expired, it's not good
	if req.MaxShelfTime > 0 {
		minDate := time.Now().AddDate(0, 0, -req.MaxShelfTime)
		if item.LastRestock.Before(minDate) {
			return false, requirements.ReasonExpired
		}
	}
	return true, requirements.ReasonNone
}
