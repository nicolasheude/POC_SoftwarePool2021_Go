package main

import (
	"SofwareGoDay3/database"
	"SofwareGoDay3/server"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	client := database.NewDatabase()
	s := server.NewServer()
	defer client.Def.Close()
	s.Def.Run()
}
