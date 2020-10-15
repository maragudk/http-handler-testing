package storage

import (
	"fmt"
)

type Storer struct{}

func (s *Storer) StartParty(id string) error {
	fmt.Printf("🥳 %v\n", id)
	return nil
}
