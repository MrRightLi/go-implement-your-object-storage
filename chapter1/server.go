package main

import (
	"./objects"
	"log"
	"net/http"
	"os"
)

func main() {
	// export LISTEN_ADDRESS=:1234
	// export STORAGE_ROOT=$GOPATH//src/go-implement-your-object-storage/storage
	log.Println("LISTEN_ADDRESS:", os.Getenv("LISTEN_ADDRESS"))
	log.Println("STORAGE_ROOT:", os.Getenv("STORAGE_ROOT"))

	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
