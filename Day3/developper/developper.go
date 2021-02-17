package developper

import (
	"SofwareGoDay3/ent"
	"context"
	"fmt"
	"log"
)

// CreateDevelopper Create a developper
func CreateDevelopper(ctx context.Context, client *ent.Client) (*ent.Developper, error) {
	d, err := client.Developper.
		Create().
		SetName("Nico").
		SetAge(18).
		SetSchool("EPITECH").
		SetExperience(1).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating developper: %v", err)
	}
	log.Println("developper was created: ", d)
	return d, nil
}
