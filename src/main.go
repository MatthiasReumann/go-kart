package main

import (
	"github.com/matthiasreumann/go-kart/src/config"
	"github.com/matthiasreumann/go-kart/src/router"
	"log"
	"net/http"
)

func main() {
	endpoints := []config.Endpoint{
		{"/test", "/Volumes/T/work/go-kart/test-files/Dockerfile"},
	}

	r := router.NewRouter(endpoints)
	err := http.ListenAndServe(":1337", r)
	if err != nil {
		log.Fatal(err)
	}
}
