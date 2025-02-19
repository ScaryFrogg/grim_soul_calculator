package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /health", s.healthHandler)
	mux.HandleFunc("GET /weapons", s.WeaponDataHandler)
	mux.HandleFunc("GET /designs", s.DesingnsHandler)
	mux.HandleFunc("GET /design/{id}", s.GetBlueprintHandler)
	mux.HandleFunc("GET /item/{id}", s.GetItemHandler)
	mux.HandleFunc("GET /requirement/{id}", s.RequirementHandler)
	mux.HandleFunc("GET /trades", s.TradesHandler)
	mux.HandleFunc("GET /trades/{id}", s.TradesForIdHandler)
	mux.HandleFunc("GET /cook", s.CookHandler)
	mux.HandleFunc("GET /armor", s.ArmorSetsHandler)
	mux.HandleFunc("GET /armor/set/{id}", s.SetPiecesHandler)
	mux.HandleFunc("GET /shields", s.GetShieldsHandler)
	mux.HandleFunc("GET /armorPerSlot", s.GetArmorPerSlotHandler)
	mux.HandleFunc("GET /enemies", s.GetEnemiesHandler)
	mux.HandleFunc("GET /enemy/{id}", s.GetEnemyByIdHandler)

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
