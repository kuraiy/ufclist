package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	gen "server/gen"
	"server/internal/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type FighterHandler struct {
	svc      domain.FighterService
	validate *validator.Validate
}

func New(svc domain.FighterService, v *validator.Validate) *FighterHandler {
	return &FighterHandler{
		svc:      svc,
		validate: v,
	}
}

func (h *FighterHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /fighters", wrap(h, decodeJSON[CreateFighterRequest], h.Create, http.StatusCreated))
	mux.HandleFunc("GET /fighters", wrap(h, decodeEmpty, h.List, http.StatusOK))
	mux.HandleFunc("GET /fighters/{id}", wrap(h, decodeID, h.GetByID, http.StatusOK))
	//mux.HandleFunc("DELETE /fighters/{id}", wrap(h, decodeID, h.Delete, http.StatusNoContent)) todo
}

func wrap[Req, Resp any](
	h *FighterHandler,
	decode func(*FighterHandler, *http.Request) (Req, error),
	fn func(context.Context, Req) (Resp, error),
	successStatus int,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decode(h, r)
		if err != nil {
			if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
				writeResponse(w, http.StatusBadRequest, formatValidationErrors(ve))
				return
			}
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := fn(r.Context(), req)

		if err != nil {
			writeError(w, statusFromError(err), err.Error())
			return
		}

		writeResponse(w, successStatus, resp)
	}
}

func decodeJSON[T any](h *FighterHandler, r *http.Request) (T, error) {
	var v T
	if err := readBody(r, &v); err != nil {
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

func statusFromError(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func (h *FighterHandler) Create(ctx context.Context, req CreateFighterRequest) (domain.Fighter, error) {
	return h.svc.Create(ctx, gen.CreateFighterParams{
		Name: req.Name,
		Age:  req.Age,
		Nickname: sql.NullString{
			String: req.Nickname,
			Valid:  req.Nickname != "",
		},
	})
}

func (h *FighterHandler) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	return h.svc.GetByID(ctx, id)
}

func (h *FighterHandler) List(ctx context.Context, _ struct{}) ([]domain.Fighter, error) {
	return h.svc.List(ctx)
}
