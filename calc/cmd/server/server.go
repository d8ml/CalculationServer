package main

import (
	"github.com/d8ml/calculation_server/calc/internal/app/server"
	"log"
)

func main() {
	err := server.NewApp().Start()
	if err != nil {
		log.Fatal(err)
	}
}
