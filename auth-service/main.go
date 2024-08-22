package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type JwtResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func main() {
	kcClient := NewKeycloakAdminClientService()
	cardService := NewCardService()
	authHandler := AuthHandler{kcClient: kcClient}

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", authHandler.Login)
	mux.HandleFunc("/auth/refresh", authHandler.Refresh)
	mux.Handle("GET /cards", JWTMiddleware(kcClient)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		cards := cardService.GetCards()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cards)
	})))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
