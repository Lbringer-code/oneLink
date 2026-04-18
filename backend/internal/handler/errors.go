package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lbringer-code/oneLink/backend/internal/domain"
	"github.com/Lbringer-code/oneLink/backend/internal/service"
)

var errorStatusMap = map[error]int{
	service.ErrValidation: http.StatusBadRequest , 
	service.ErrNotFound: http.StatusNotFound ,
	service.ErrInternal: http.StatusInternalServerError ,
}

func (h *Handler) writeJSON( w http.ResponseWriter , status int ,  data any) {
	body , err := json.Marshal(data)
	if err != nil {
		h.logger.Error("failed to marshal response" , "error" , err)
		http.Error( w , "something went wrong" , http.StatusInternalServerError )
		return
	}

	w.Header().Set("Content-type" , "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func (h *Handler) writeError( w http.ResponseWriter , err error ) {
	status := http.StatusInternalServerError
	for sentinel , code := range errorStatusMap {
		if errors.Is(err , sentinel) {
			status = code
			break
		}
	}

	msg := err.Error()
	if status == http.StatusInternalServerError {
		h.logger.Error("internal server error" , "error" , err)
		msg = "something went wrong"
	}

	h.writeJSON(w , status , domain.ErrorResponse{Error: msg})
}