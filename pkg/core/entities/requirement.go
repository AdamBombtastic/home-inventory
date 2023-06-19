package entities

type Requirement struct {
	ID              int      `yaml:"ID"`
	Name            string   `yaml:"Name"`
	Quantity        int      `yaml:"Quantity"`
	Units           string   `yaml:"Units"`
	MaxShelfTime    int      `yaml:"MaxShelfTime"`
	ApplicableItems []string `yaml:"ApplicableItems"`
}

type Reason string

const (
	ReasonExpired       Reason = "Expired"
	ReasonNotApplicable Reason = "NA"
	ReasonTooFew        Reason = "Too Few"
	ReasonNone          Reason = ""
)
