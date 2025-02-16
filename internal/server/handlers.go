package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) TradesForIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	writeJSONResponse(w, s.db.GetTradesForItem(id), http.StatusOK)
}

func (s *Server) ArmorSetsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetSets(), http.StatusOK)
}

func (s *Server) GetShieldsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.ShieldsData(), http.StatusOK)
}
func (s *Server) GetArmorPerSlotHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.ArmorPerSlot(), http.StatusOK)
}

func (s *Server) TradesHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetTrades(), http.StatusOK)
}

func (s *Server) DesingnsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetDesigns(), http.StatusOK)
}

func (s *Server) SetPiecesHandler(w http.ResponseWriter, r *http.Request) {
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		http.Error(w, "Id not valid", http.StatusNotFound)
	}

	writeJSONResponse(w, s.db.GetPiecesForSet(id), http.StatusOK)
}

func (s *Server) RequirementHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	writeJSONResponse(w, s.db.GetRequirements(id), http.StatusOK)
}

func (s *Server) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" || id == "" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	//get design
	material, err := s.db.GetItem(id)
	if err != nil {
		fmt.Printf("Failed to get material with Id [%v] design data-> error %v", id, err)
		return
	}

	//get requirements
	designs := s.db.GetDesignsForItem(id)
	w.Header().Set("Content-Type", "application/json")
	material.Designs = designs
	writeJSONResponse(w, material, http.StatusOK)
}

func (s *Server) GetBlueprintHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	//get design
	design := s.db.GetDesign(id)

	//get requirements
	requirements := s.db.GetRequirements(id)
	design.Requirements = requirements
	writeJSONResponse(w, design, http.StatusOK)
}

func (s *Server) WeaponDataHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.WeaponData(), http.StatusOK)
}

func (s *Server) GetEnemiesHandeler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetEnemies(), http.StatusOK)
}

func (s *Server) CookHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetCookData(), http.StatusOK)
}
