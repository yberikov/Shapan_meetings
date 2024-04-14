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
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
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
	log.Fatal(http.ListenAndServe(":8080", router))

}
