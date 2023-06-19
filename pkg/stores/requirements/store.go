package requirements

import (
	"fmt"
	"os"

	"github.com/adambombtastic/home-inventory/pkg/core/entities"
	"gopkg.in/yaml.v2"
)

type Store interface {
	All() ([]*entities.Requirement, error)
	// Add(r *Requirement) error
	// Remove(id int) error
}

type yamlStore struct {
	items        []*entities.Requirement
	fileLocation string
}

func (y *yamlStore) load() error {
	inData, err := os.ReadFile(y.fileLocation)
	if err != nil {
		return fmt.Errorf("failed to load requirements %w", err)
	}
	return yaml.Unmarshal(inData, &y.items)
}

func New() (Store, error) {
	// This should be configurable
	fileLocation := `C:\Users\adama\Documents\Github\home-inventory\data\requirements\store.yaml`
	y := &yamlStore{fileLocation: fileLocation}
	err := y.load()
	if err != nil {
		return nil, fmt.Errorf("failed to create yaml store %w", err)
	}
	return y, nil
}

func (y *yamlStore) All() ([]*entities.Requirement, error) {
	return y.items, nil
}
