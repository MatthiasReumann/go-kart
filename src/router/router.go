package router

import (
	"github.com/matthiasreumann/go-kart/src/config"
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
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			{

			}
		case http.MethodGet:
		case http.MethodPut:
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
