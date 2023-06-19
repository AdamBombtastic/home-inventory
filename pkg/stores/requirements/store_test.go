package requirements

import (
	"testing"
)

func Test_load(t *testing.T) {
	y := &yamlStore{fileLocation: "./testdata/store.yaml"}
	err := y.load()
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if len(y.items) != 7 {
		t.Error("Expected 7, got", len(y.items))
	}
	if y.items[0].Name != "Paper Towels" {
		t.Error("Expected Paper Towels, got", y.items[0].Name)
	}
}
