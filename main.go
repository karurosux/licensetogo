package main

import (
	"licensetogo/internal/app"
	_ "licensetogo/migrations"
	"log"
)

func main() {
	ls, err := app.NewLicenseToGoServer()
	if err != nil {
		log.Fatalf("Failed to create license to go server: %v", err)
		return
	}

	if err := ls.Start(); err != nil {
		log.Fatal(err)
	}
}
