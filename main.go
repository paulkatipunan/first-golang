package main

import (
	"log"
	"sample/routes"
	"net/http"
	"sample/config"
	"fmt"
)

func main() {
	port := config.GetEnv("PORT")

	fmt.Println(fmt.Sprintf("Application started on port :%v", port))

	routes := router.Routes()

	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port) , nil))
}
