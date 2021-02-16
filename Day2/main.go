package main

import (
	"SofwareGoDay2/server"
)

func main() {
	s := server.NewServer()
	s.Def.Run()
}
