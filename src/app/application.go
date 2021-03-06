package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":9000",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
