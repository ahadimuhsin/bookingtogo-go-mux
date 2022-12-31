package main

import (
	"fmt"
	"go-backend-test/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Server dijalankan pada port 8080 ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}