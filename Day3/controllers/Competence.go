package controllers

import (
	"SofwareGoDay3/ent"
	"SofwareGoDay3/ent/competence"
	"SofwareGoDay3/ent/developper"
	"context"
	"fmt"
)

// CreateCompetence Create a contact from a developper ID
func CreateCompetence(id, level int,  name string, ctx context.Context, client *ent.Client) (*ent.Competence, error) {
	c, err := client.Competence.
		Create().
		SetOwnerID(id).
		SetName(name).
		SetLevel(level).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Competence: %v", err)
	}
	fmt.Printf("competence was created\n")
	return c, nil
}

// GetCompetence Allows you to retrieve a contact from an id
func GetCompetence(id int, ctx context.Context, client *ent.Client) (*ent.Competence, error) {
	c, err := client.Developper.
		Query().
		Where(developper.ID(id)).
		QueryCompetence().
		Only(ctx)
	return c, err
}

// UpdateCompetence Update a contact from a developper ID
func UpdateCompetence(id, level int, name string, ctx context.Context, client *ent.Client) error {
	_, err := client.Competence.
		Update().
		Where(competence.HasOwnerWith(developper.ID(id))).
		SetName(name).
		SetLevel(level).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed update competence: %v", err)
	}
	fmt.Printf("competence was update\n")
	return nil
}

// DeleteCompetence Delete a contact from a developper ID
func DeleteCompetence(id int, ctx context.Context, client *ent.Client) error {
	_, err := client.Competence.
		Delete().
		Where(competence.HasOwnerWith(developper.ID(id))).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed delete competence: %v", err)
	}
	return nil
}