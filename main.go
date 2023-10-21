package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	routes "go-api/app"

)

func main() {
	routes := routes.SetupRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(routes))

	port := ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
