package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imrushi/restapi/helper"
	"github.com/imrushi/restapi/models"
	"github.com/imrushi/restapi/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

func gettodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var todos []models.Todo

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var todo models.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(todos)
}

func gettodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func createtodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo

	_ = json.NewDecoder(r.Body).Decode(&todo)

	result, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updatetodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var todo models.Todo

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&todo)

	update := bson.D{
		{"$set", bson.D{
			{"body", todo.Body},
			{"completed", todo.Completed},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&todo)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	todo.ID = id

	json.NewEncoder(w).Encode(todo)
}

func deletetodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api/todos", gettodos).Methods("GET")
	r.HandleFunc("/api/todos/{id}", gettodo).Methods("GET")
	r.HandleFunc("/api/todos", createtodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", updatetodo).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", deletetodo).Methods("DELETE")

	fmt.Print("Server is running on port: " + config.API_PORT)
	log.Fatal(http.ListenAndServe(":"+config.API_PORT,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"}),
			handlers.AllowedHeaders([]string{"X-requested-With", "Content-Type", "Authorization"}),
		)(r)))
}
