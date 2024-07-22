package api

import (
	"context"
	"log"
	"net/http"
	"terraform-mongodb-pratical-example/api/controllers"
	"terraform-mongodb-pratical-example/repositories"
	"terraform-mongodb-pratical-example/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartServer() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("football")
	playersCollection := db.Collection("players")

	playerrepo := repositories.PlayerMongoDBRepository{Collection: playersCollection}

	playerservice := services.PlayerService{Repository: &playerrepo}

	playercontroller := controllers.PlayerController{Service: &playerservice}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /players", playercontroller.Save)

	http.ListenAndServe(":8090", mux)
}
