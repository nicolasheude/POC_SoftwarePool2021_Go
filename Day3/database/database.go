package database

import (
	"SofwareGoDay3/ent"
	"context"
	"fmt"
	"log"
	"os"
)

type Database struct {
	Def *ent.Client
	Ctx context.Context
}

var Db Database

func NewDatabase() Database {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_NAME"),
			os.Getenv("POSTGRES_PASS")))
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	} else {
		fmt.Printf("Database postgres is ready\n")
	}
	Db = Database{Def: client, Ctx: context.Background()}
	if err := Db.Def.Schema.Create(Db.Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return Db
}
