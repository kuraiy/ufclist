package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func readBody(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func writeResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeResponse(w, status, map[string]string{"error": msg})
}

func formatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		switch field + "." + e.Tag() {
		case "age.gte":
			errors[field] = "age must be at least 18"
		case "age.lte":
			errors[field] = "age must be at most 100"
		case "name.required":
			errors[field] = "name is required"
		default:
			errors[field] = e.Tag()
		}
	}
	return errors
}
