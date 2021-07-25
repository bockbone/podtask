package main

import (
	"log"
	"net/http"

	"github.com/bockbone/podtask/routers"
)

func main() {

	router := routers.InitRoutes()

	log.Printf("Server is starting on port 5000 ...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
