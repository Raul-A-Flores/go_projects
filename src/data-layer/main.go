package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Name string `json: "name, omitempty"`
	Age  int    `json: "age, omitempty"`
}

type Response struct {
	Message string `json: "message, omitempty"`
	Data    Person `json: "data, omitempty"`
	Error   string `json: "error, omitempty"`
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

func personHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	defer func() {
		json.NewEncoder(w).Encode(res)
	}()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	switch r.Method {
	case http.MethodPost:

		var p Person
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			res.Error = err.Error()
			return
		}

		id := createPerson(p)
		res.Message = fmt.Sprintf("New person created with id: %d", id)

	case http.MethodGet:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("invalid id")
			res.Error = "invalid id"
			w.WriteHeader(400)
			return
		}

		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		p, err := readPerson(id)
		if err != nil {
			log.Println("read error", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		res.Data = p

	case http.MethodPut:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("invalid id")
			res.Error = "invalid id"
			w.WriteHeader(400)
			return
		}

		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		var p Person
		json.NewDecoder(r.Body).Decode(&p)

		err = updatePerson(id, p)
		if err != nil {
			log.Println("update person error", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		res.Message = "person successfully updated"

	case http.MethodDelete:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("invalid id")
			res.Error = "invalid id"
			w.WriteHeader(400)
			return
		}

		id, err := strconv.Atoi(strID)
		if err != nil {
			log.Println("invalid id", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		err = deletePerson(id)
		if err != nil {
			log.Println("Delete Error", err)
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}

		res.Message = "person successfuly deleted"

	}
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
