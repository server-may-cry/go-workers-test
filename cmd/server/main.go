package main

import (
	"github.com/server-may-cry/go-workers-test/internal/di"
	"log"
	"net/http"
)

func main() {
	container := di.New()

	go container.DummyWorker.Go()

	http.HandleFunc("/strange", container.StrangeHandler)

	log.Println("started")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
