package main

import (
	"log"
	"github.com/joho/godotenv"
	"sample/routes"
	"net/http"
	"fmt"
	"os"
)

func main() {
	initEnv()
	port := os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("Application started on port :%v", port))

	routes := router.Routes()

	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port) , nil))
}

func initEnv() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
}
