package interactors

import (
	"testing"
	"time"

	"github.com/adambombtastic/home-inventory/pkg/inventory"
	"github.com/adambombtastic/home-inventory/pkg/requirements"
)

func Test_InventoryMeetsRequirements(t *testing.T) {
	today := time.Now()
	yearAgo := today.AddDate(-1, 0, 0)
	inventory := []*inventory.Item{
		{Name: "Snacks", Stock: 5, LastRestock: yearAgo, ID: 0},
		{Name: "Toilet Paper", Stock: 0, LastRestock: yearAgo, ID: 1},
		{Name: "Paper Towels", Stock: 10, LastRestock: yearAgo, ID: 2},
	}
	reqs := []*requirements.Requirement{
		{Name: "Snacks", Quantity: 5, Units: "bags", MaxShelfTime: 90, ID: 0, ApplicableItems: []string{"Snacks"}},
		{Name: "Toilet Paper", Quantity: 1, Units: "rolls", ID: 1, ApplicableItems: []string{"Toilet Paper"}},
		{Name: "Paper Towels", Quantity: 5, Units: "rolls", ID: 2, ApplicableItems: []string{"Paper Towels"}},
	}

	missing := InventoryMeetsRequirements(inventory, reqs)
	if len(missing) != 2 {
		t.Error("Expected 2, got", len(missing))
	}
	if len(missing[0].Reasons) != 3 {
		t.Error("Expected 3, got", len(missing[0].Reasons))
	}
	if missing[0].Reasons[0].Reason != requirements.ReasonExpired {
		t.Error("Expected", requirements.ReasonExpired, "got", missing[0].Reasons[0])
	}
}

func Test_itemMeetsRequirement(t *testing.T) {

	today := time.Now()
	yearAgo := today.AddDate(-1, 0, 0)

	type testcase struct {
		name       string
		item       *inventory.Item
		req        *requirements.Requirement
		want       bool
		wantReason requirements.Reason
	}
	testcases := []testcase{
		{
			name:       "Snacks -- Expired",
			item:       &inventory.Item{Name: "Snacks", Stock: 5, LastRestock: yearAgo},
			req:        &requirements.Requirement{Name: "Snacks", Quantity: 5, Units: "bags", MaxShelfTime: 90, ApplicableItems: []string{"Snacks"}},
			want:       false,
			wantReason: requirements.ReasonExpired,
		},
		{
			name:       "Snacks -- Happy",
			item:       &inventory.Item{Name: "Snacks", Stock: 5, LastRestock: today},
			req:        &requirements.Requirement{Name: "Snacks", Quantity: 5, Units: "bags", MaxShelfTime: 90, ApplicableItems: []string{"Snacks"}},
			want:       true,
			wantReason: requirements.ReasonNone,
		},
		{
			name:       "Incompatible Units",
			item:       &inventory.Item{Name: "Snacks", Stock: 5, LastRestock: yearAgo},
			req:        &requirements.Requirement{Name: "Toilet Paper", Quantity: 5, Units: "rolls", ApplicableItems: []string{"Toilet Paper", "Paper Towels"}},
			want:       false,
			wantReason: requirements.ReasonNotApplicable,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, gotReason := itemMeetsRequirement(tc.req, tc.item)
			if got != tc.want {
				t.Errorf("case %s: want %t, got %t", tc.name, tc.want, got)
			}
			if gotReason != tc.wantReason {
				t.Errorf("case %s: want %s, got %s", tc.name, tc.wantReason, gotReason)
			}
		})
	}

}

func Test_ItemsMeetRequirement(t *testing.T) {
	today := time.Now()

	items := []*inventory.Item{
		{Name: "Snacks", Stock: 5, LastRestock: today},
		{Name: "Chips", Stock: 5, LastRestock: today},
	}
	req := &requirements.Requirement{Name: "Snacks", Quantity: 8, Units: "bags", MaxShelfTime: 90, ApplicableItems: []string{"Snacks", "Chips"}}

	got, gotReason := ItemsMeetRequirement(req, items)
	if got != true {
		t.Errorf("want %t, got %t", false, got)
	}

	if len(gotReason) != 0 {
		t.Errorf("want %d, got %d", 0, len(gotReason))
	}

}
