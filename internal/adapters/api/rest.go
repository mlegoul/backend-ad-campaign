package api

import (
	"backend-ad-campaign/internal/core"
	"backend-ad-campaign/internal/ports"
	"encoding/json"
	"log"
	"net/http"
)

type CampaignHandler struct {
	Service ports.Repository
}

func NewCampaignHandler(service ports.Repository) *CampaignHandler {
	return &CampaignHandler{Service: service}
}

func (h *CampaignHandler) HandleCreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign core.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err) // Log l'erreur ici
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Received campaign: %+v", campaign)

	createdCampaign, err := h.Service.CreateCampaign(&campaign)
	if err != nil {
		http.Error(w, "Error creating campaign", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCampaign)
}
