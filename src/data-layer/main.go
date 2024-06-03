package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json: "name, omitempty"`
	Age  int    `json: "age, omitempty"`
}

var datalayer map[int]Person
var personID int

func init() {
	datalayer = make(map[int]Person)
	personID = 0

	log.Println("datalayer initiated")
}

func main() {
	http.HandleFunc("/", healthHandler)

	log.Println("server starting at :4444")

	err := http.ListenAndServe(":4444", nil)
	if err != nil {

		log.Fatal("server error", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	w.Write([]byte("Server is running..."))
}

// Datalayer

func createPerson(p Person) int {
	pID := personID + 1
	datalayer[pID] = p

	// update
	personID = pID

	return pID
}

func readPerson(id int) (Person, error) {
	if person, ok := datalayer[id]; ok {
		return person, nil
	}

	return Person{}, fmt.Errorf("person id not found. id: %d", id)

}

func updatePerson(id int, p Person) error {
	if _, ok := datalayer[id]; ok {
		datalayer[id] = p
		return nil
	}

	return fmt.Errorf("person id not found to update id: %d", id)
}

func deletePerson(id int) error {
	if _, ok := datalayer[id]; ok {
		delete(datalayer, id)
		return nil
	}

	return fmt.Errorf("Unable to delete person id: %d", id)
}
