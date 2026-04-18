package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lbringer-code/oneLink/backend/internal/domain"
	"github.com/Lbringer-code/oneLink/backend/internal/service"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateBundle( w http.ResponseWriter , r *http.Request ) {
	var req domain.CreateBundleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.writeError(w , errors.Join(service.ErrValidation , errors.New("invalid JSON body")))
		return
	}
	defer r.Body.Close()

	resp , err := h.svc.CreateBundle(req)
	if err != nil {
		h.writeError(w , err)
		return
	}

	h.writeJSON(w , http.StatusCreated , resp)
}

func (h *Handler) GetBundle( w http.ResponseWriter , r *http.Request ) {
	slug := chi.URLParam(r , "slug")

	resp , err := h.svc.GetBundle(slug)
	if err != nil {
		h.writeError(w , err)
		return
	}

	h.writeJSON(w , http.StatusOK , resp)
}