package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
    "github.com/joho/godotenv"
	"github.com/conglt10/web-golang/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	
	router := httprouter.New()
	
	router.POST("/auth/login", routes.Login)
	router.POST("/auth/register", routes.Register)

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}