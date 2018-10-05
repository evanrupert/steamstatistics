package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func main() {
	RunDatabaseMigrations()

	http.HandleFunc("/api/data/", DataController)
	http.HandleFunc("/api/test", TestController)

	fmt.Printf("Web server started on port %s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		panic(err)
	}
}
