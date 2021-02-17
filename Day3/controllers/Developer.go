package controllers

import (
	"SofwareGoDay3/ent"
	"SofwareGoDay3/ent/developper"
	"context"
	"fmt"
	"log"
)

// CreateDeveloper Create a developer from its ID
func CreateDeveloper(age, experience int, name, school string, ctx context.Context, client *ent.Client) (*ent.Developper, error) {
	d, err := client.Developper.
		Create().
		SetName(name).
		SetAge(age).
		SetSchool(school).
		SetExperience(experience).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating developper: %v", err)
	}
	log.Println("developper was created: ", d)
	return d, nil
}

// GetDevelopers Get a developer from its ID
func GetDevelopers(id int, ctx context.Context, client *ent.Client) (*ent.Developper, error) {
	d, err := client.Developper.
		Query().
		Where(developper.ID(id)).
		Only(ctx)
	return d, err
}

// UpdateDeveloper Update a developer from its ID
func UpdateDeveloper(id, age, experience int, name, school string, ctx context.Context, client *ent.Client) (*ent.Developper, error) {
	d, err := client.Developper.
		UpdateOneID(id).
		SetName(name).
		SetAge(age).
		SetSchool(school).
		SetExperience(experience).
		Save(ctx)
	return d, err
}

// DeleteDeveloper Removes a developer from its ID
func DeleteDeveloper(id int, ctx context.Context, client *ent.Client) error {
	err := client.Developper.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// TheoBonus Adds a year to the age of a developer
func TheoBonus(id int, ctx context.Context, client *ent.Client) (*ent.Developper, error) {
	d, err := GetDevelopers(id, ctx, client)
	newAge := d.Age + 1
	d, err = UpdateDeveloper(id, newAge, d.Experience, d.Name, d.School, ctx, client)
	return d, err
}
