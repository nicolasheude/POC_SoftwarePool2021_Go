package main

import (
	"SofwareGoDay4/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.NewServer()
	server.Def.Run()
}