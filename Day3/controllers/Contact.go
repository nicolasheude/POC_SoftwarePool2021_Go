package controllers

import (
	"SofwareGoDay3/ent"
	"SofwareGoDay3/ent/contact"
	"SofwareGoDay3/ent/developper"
	"context"
	"fmt"
)

// CreateContact Create a contact from a developper ID
func CreateContact(id int, email, phone, github, linkedin string, ctx context.Context, client *ent.Client) (*ent.Contact, error) {
	c, err := client.Contact.
		Create().
		SetOwnerID(id).
		SetEmail(email).
		SetPhone(phone).
		SetGithub(github).
		SetLinkedin(linkedin).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating contact: %v", err)
	}
	fmt.Printf("contact was created\n")
	return c, nil
}

// GetContact Allows you to retrieve a contact from an id
func GetContact(id int, ctx context.Context, client *ent.Client) (*ent.Contact, error) {
	c, err := client.Developper.
		Query().
		Where(developper.ID(id)).
		QueryContact().
		Only(ctx)
	return c, err
}

// UpdateContact Update a contact from a developper ID
func UpdateContact(id int, email, phone, github, linkedin string, ctx context.Context, client *ent.Client) error {
	_, err := client.Contact.
		Update().
		Where(contact.HasOwnerWith(developper.ID(id))).
		SetEmail(email).
		SetPhone(phone).
		SetGithub(github).
		SetLinkedin(linkedin).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed update contact: %v", err)
	}
	fmt.Printf("contact was update\n")
	return nil
}

// DeleteContact Delete a contact from a developper ID
func DeleteContact(id int, ctx context.Context, client *ent.Client) error {
	_, err := client.Contact.
		Delete().
		Where(contact.HasOwnerWith(developper.ID(id))).
		Exec(ctx)
	return err
}
