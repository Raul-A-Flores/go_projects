package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":4000", store)
	server.Run()

	fmt.Printf("%+v\n", store)
}
