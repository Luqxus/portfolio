package api

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"time"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

var nextFs embed.FS

type APIServer struct {
	mux     *http.ServeMux
	port    int
	address string
}

func New(address string, port int) *APIServer {
	return &APIServer{
		mux:     http.NewServeMux(),
		port:    port,
		address: address,
	}
}

func (api *APIServer) Run() error {

	//root endpoint | serve nextjs ui
	//distFs, err := fs.Sub(nextFs, "ui/dist")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// serve nextjs
	api.mux.Handle("/", http.FileServer(http.Dir("ui/dist")))

	return http.ListenAndServe(fmt.Sprintf("%s:%d", api.address, api.port), api.mux)
}

func (api *APIServer) handleFunc(fn APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(request.Context(), 30*time.Second)
		defer cancel()

		err := fn(writer, request.WithContext(ctx))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}
