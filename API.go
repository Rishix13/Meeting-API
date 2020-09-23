package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Meeting struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title	    string             `json:"Title,omitempty" bson:"Title,omitempty"`
	Participant string 			   `json:"Participant,omitempty" bson:"Participant,omitempty"`
	Starttime   string             `json:"starttime,omitempty" bson:"startitme",omitempty`
	endtime     string 			   `json:"endtime,omitempty" bson:"endtime",omitempty`
	timestamp	string 	     	   `json:"timestamp,omitempty" bson:"timestamp",omitempty`

}

func CreateParticipantEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var Meeting Participant
	_ = json.NewDecoder(request.Body).Decode(&Meeting)
	collection := client.Database("thepolyglotdeveloper").Collection("Meeting")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, Meeting)
	json.NewEncoder(response).Encode(result)
}
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) { }

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/Meeting", CreateParticipantEndpoint).Methods("POST")
	router.HandleFunc("/Meeting/{id}", GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)