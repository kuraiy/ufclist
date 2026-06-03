package handler

import (
	"encoding/json"
	"net/http"
	"server/internal/domain"
	"strconv"
)

type FighterHandler struct {
	svc domain.FighterService
}

func New(svc domain.FighterService) *FighterHandler {
	return &FighterHandler{svc: svc}
}

func (h *FighterHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /fighters", h.Create)
	mux.HandleFunc("GET /fighters", h.List)
	mux.HandleFunc("GET /fighters/{id}", h.GetByID)
	mux.HandleFunc("DELETE /fighters/{id}", h.Delete)
}

func (h *FighterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var f domain.Fighter
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.svc.Create(r.Context(), f)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *FighterHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fighter, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "applciation/json")
	json.NewEncoder(w).Encode(fighter)
}

func (h *FighterHandler) List(w http.ResponseWriter, r *http.Request) {
	fighters, err := h.svc.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fighters)
}

func (h *FighterHandler) Delete(w http.ResponseWriter, r *http.Request) {}
