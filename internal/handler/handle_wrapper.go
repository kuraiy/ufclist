package handler

import (
	"context"
	"encoding/json"
	"errors"
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

func Decode[T any](r *http.Request) (T, error) {
	var req T

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func (h *FighterHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /fighters", wrap(h, decodeJSON[CreateFighterRequest], h.Create, http.StatusCreated))
	mux.HandleFunc("GET /fighters", wrap(h, decodeEmpty, h.List, http.StatusOK))
	mux.HandleFunc("GET /fighters/{id}", wrap(h, decodeID, h.GetByID, http.StatusOK))
	mux.HandleFunc("DELETE /fighters/{id}", wrap(h, decodeID, h.Delete, http.StatusOK))
	mux.HandleFunc("PUT /fighters/{id}", wrap(h, decodeUpdateFighter, h.Update, http.StatusOK))
}

func wrap[Req, Resp any](
	h *FighterHandler,
	decoder func(*FighterHandler, *http.Request) (Req, error),
	fn func(context.Context, Req) (Resp, error),
	successStatusCode int,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decoder(h, r)
		if err != nil {
			if val, ok := errors.AsType[validator.ValidationErrors](err); ok {
				writeJSON(w, http.StatusBadRequest, formatValidationErrors(val))
				return
			}
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := fn(r.Context(), req)

		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, successStatusCode, resp)
	}
}

func decodeJSON[T any](h *FighterHandler, r *http.Request) (T, error) {
	var v T

	if err := readJSON(r, &v); err != nil {
		return v, err
	}

	if err := h.validate.Struct(v); err != nil {
		return v, err
	}

	return v, nil
}

func decodeID(_ *FighterHandler, r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func decodeEmpty(_ *FighterHandler, _ *http.Request) (struct{}, error) {
	return struct{}{}, nil
}

func decodeUpdateFighter(h *FighterHandler, r *http.Request) (domain.UpdateFighterInput, error) {
	req, err := decodeJSON[domain.UpdateFighterInput](h, r)
	if err != nil {
		return domain.UpdateFighterInput{}, err
	}

	id, err := decodeID(h, r)

	if err != nil {
		return domain.UpdateFighterInput{}, err
	}

	req.ID = id
	return req, nil
}

func (h *FighterHandler) Create(ctx context.Context, req CreateFighterRequest) (domain.Fighter, error) {
	return h.svc.Create(ctx, domain.Fighter{
		Name:     req.Name,
		Age:      uint8(req.Age),
		Nickname: req.Nickname,
	})
}

func (h *FighterHandler) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	return h.svc.GetByID(ctx, id)
}

func (h *FighterHandler) List(ctx context.Context, _ struct{}) ([]domain.Fighter, error) {
	return h.svc.List(ctx)
}

func (h *FighterHandler) Delete(ctx context.Context, id int64) (struct{}, error) {
	return h.svc.Delete(ctx, id)
}

func (h *FighterHandler) Update(ctx context.Context, req domain.UpdateFighterInput) (domain.Fighter, error) {
	return h.svc.Update(ctx, req)
}
