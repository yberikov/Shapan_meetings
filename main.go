package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"hacknu/internal/gc"
	handler2 "hacknu/internal/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	//storage.CreateConn()
	ctx := context.Background()
	// Read credentials from environment variables
	credentialsJSON := os.Getenv("GOOGLE_CREDENTIALS_JSON")
	if credentialsJSON == "" {
		log.Fatal("GOOGLE_CREDENTIALS_JSON environment variable is not set")
	}
	fmt.Println(credentialsJSON)

	config, err := google.ConfigFromJSON([]byte(credentialsJSON), calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := gc.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	handler := handler2.NewHandler(srv)
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/searchSpeaking", handler.DataHandler).Methods("POST", "OPTIONS")
	//router.HandleFunc("/login", handler.SignIn).Methods("POST", "OPTIONS")
	fmt.Println("Server started:")
	log.Fatal(http.ListenAndServe("0.0", router))

}
