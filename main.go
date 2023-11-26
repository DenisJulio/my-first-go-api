package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

const serverAddress = "localhost:8080"

func main() {
	r := mux.NewRouter()
	dbClient := connectToDB()
	h := &Handler{DBClient: dbClient}

	r.HandleFunc("/messages", h.getMessageHandler).Methods("GET")

	go func() {
		log.Println("Server started on http://" + serverAddress)
		log.Fatal(http.ListenAndServe(serverAddress, r))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server disconnected")	
}
