package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yuchenOutlook/hotel-reservation/api"
	"github.com/yuchenOutlook/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	// connect to mongo db, first you need to docker run mongo db container
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{ 
		FirstName: "Oriana",
		LastName: "Grande",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var oriana types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&oriana); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Oriana: ", oriana)
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	// This parses the command line args from [1:]
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
