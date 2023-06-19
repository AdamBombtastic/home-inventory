package inventory

import (
	"testing"
)

func Test_load(t *testing.T) {
	y := &yamlStore{fileLocation: "./testdata/store.yaml"}
	err := y.load()
	if err != nil {
		t.Error("Expected nil, got", err)
	}

}
