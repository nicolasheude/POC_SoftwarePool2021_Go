package main

import (
	"SofwareGoDay2/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := server.NewServer()
	s.Def.Run()
}
