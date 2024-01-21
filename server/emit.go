package server

import (
	"encoding/json"
	"net/http"
)

type event struct {
	Method  string `json:"method"`
	Message string `json:"message"`
}

func (s *Server) EmitEvent(w http.ResponseWriter, req *http.Request) {
	var e event
	err := json.NewDecoder(req.Body).Decode(&e)
	if err != nil {
		s.returnResponse(w, http.StatusInternalServerError, err)
	}

	err = s.client.Emit(e.Method, e.Message)
	if err != nil {
		s.returnResponse(w, http.StatusInternalServerError, err)
	}

	s.returnResponse(w, http.StatusOK, nil)
}
