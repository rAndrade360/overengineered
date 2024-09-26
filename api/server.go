package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"terraform-mongodb-pratical-example/api/controllers"
	"terraform-mongodb-pratical-example/repositories"
	"terraform-mongodb-pratical-example/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func StartServer() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.PrimaryPreferred())
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

	http.ListenAndServe(":"+os.Getenv("API_PORT"), mux)
}
