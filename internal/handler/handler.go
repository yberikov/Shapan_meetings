package handler

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"hacknu/internal/gc"
	"hacknu/internal/match"
	"hacknu/internal/models"
	"net/http"
	"time"
)

type Handler struct {
	services *calendar.Service
	userPool *match.UserPool
}

func NewHandler(service *calendar.Service) *Handler {
	return &Handler{
		services: service,
		userPool: &match.UserPool{
			Users: make([]models.User, 0),
		},
	}
}

type SuccessResponse struct {
	Message string `json:"message"`
	// Add other fields as needed
}

func (h *Handler) DataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// Set CORS headers for the preflight request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.userPool.AddUser(user)

	user1, user2, err := h.userPool.FindMatch()
	if err != nil {
		fmt.Printf("Unable to find match: %v", err)
	}
	fmt.Println(user1)
	fmt.Println(user2)
	if user1 != (models.User{}) && user2 != (models.User{}) {
		createdEvent, err := gc.CreateEvent(h.services, user1, user2)
		if err != nil {
			fmt.Printf("Unable to create event: %v", err)
		}
		fmt.Printf("Event created: %s\n", createdEvent.HtmlLink)
	}
	successResponse := SuccessResponse{
		Message: "Success!",
	}
	jsonResponse, err := json.Marshal(successResponse)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	return
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method == "OPTIONS" {
		// Set CORS headers for the preflight request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   "test",
		Path:    "/",
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	return
}
