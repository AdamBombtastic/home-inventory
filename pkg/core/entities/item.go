package entities

import "time"

type Item struct {
	ID          int       `yaml:"ID"`
	Name        string    `yaml:"Name"`
	Stock       int       `yaml:"Stock"`
	LastRestock time.Time `yaml:"LastRestock"`
}
