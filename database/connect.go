package go 

import (
	"log" 
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/option"

)

func ConnectUsers ( ) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("mongodb+srv://nhatlinh:nhatlinh@cluster0.ajj2g.mongodb.net/test?retryWrites=true&w=majority"))
	client ,err := mongo.Connect(context.TODO(), clientOptions)

	if err != nill {
		log.Fatal(err)
	}
	err = client.Ping (context.TODO() , nil )

	if err != nil {
		log.Fatal(err )
	}

	collection := client .Database("go lang ") .Collection("user")
	return collection 
}