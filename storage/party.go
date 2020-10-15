package storage

import (
	"errors"
	"fmt"
)

type Storer struct{}

func (s *Storer) StartParty(id string) error {
	fmt.Printf("🥳 %v\n", id)
	return nil
}

func (s *Storer) StopParty(id string) error {
	return errors.New("the party cannot be stopped")
}
