package main

import (
	"log"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("initialize app: %v", err)
	}

	defer app.Close()

	err = app.Run()
	if err != nil {
		log.Fatalf("application run: %v", err)
	}
}
