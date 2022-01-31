package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start(host string, port int) {
	
	router := mux.NewRouter()

	router.HandleFunc(`/api/time`, now)
	router.HandleFunc(`/api/time`, now).Queries("tz", "{tz}")
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port),router))
}