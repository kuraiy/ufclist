package handler

import (
	"net/http"
	"server/internal/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type FighterHandler struct {
	svc      domain.FighterService
	validate *validator.Validate
}

func New(svc domain.FighterService) *FighterHandler {
	return &FighterHandler{
		svc:      svc,
		validate: validator.New(),
	}
}

func (h *FighterHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /fighters", h.Create)
	mux.HandleFunc("GET /fighters", h.List)
	mux.HandleFunc("GET /fighters/{id}", h.GetByID)
	mux.HandleFunc("DELETE /fighters/{id}", h.Delete)
}

func (h *FighterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateFighterRequest

	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		writeJSON(w, http.StatusBadRequest, formatValidationErrors(err))
		return
	}

	created, err := h.svc.Create(r.Context(), domain.Fighter{
		Name:     req.Name,
		Age:      uint8(req.Age),
		Nickname: req.Nickname,
	})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *FighterHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	fighter, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, fighter)
}

func (h *FighterHandler) List(w http.ResponseWriter, r *http.Request) {
	fighters, err := h.svc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, fighters)
}

func (h *FighterHandler) Delete(w http.ResponseWriter, r *http.Request) {}
