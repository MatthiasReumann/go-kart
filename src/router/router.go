package router

import (
	"github.com/matthiasreumann/go-kart/src/config"
	"github.com/matthiasreumann/go-kart/src/docker"
	"log"
	"net/http"
)

func NewRouter(endpoints []config.Endpoint) *http.ServeMux {
	mux := http.NewServeMux()
	for _, endpoint := range endpoints {
		mux.HandleFunc(endpoint.Path, handleRequest(endpoint))
	}
	return mux
}

func handleRequest(endpoint config.Endpoint) http.HandlerFunc {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		log.Fatal(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			{
				err := dockerClient.Build(endpoint.DockerfileTar, endpoint.Tags)
				if err != nil {
					log.Println(err)
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				err = dockerClient.Run(endpoint.Tags[0])
				if err != nil {
					log.Println(err)
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				writer.WriteHeader(http.StatusOK)
				return
			}
		case http.MethodGet:
			fallthrough
		case http.MethodPut:
			fallthrough
		case http.MethodDelete:
			writer.WriteHeader(http.StatusForbidden)
			_, err := writer.Write([]byte("Method not allowed"))
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}
