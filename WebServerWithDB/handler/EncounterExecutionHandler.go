package handler

import (
	"database-example/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EncounterExecutionHandler struct {
	EncounterExecutionService *service.EncounterExecutionService
}

func (handler *EncounterExecutionHandler) GetExecutionByUser(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr, ok := vars["userId"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encounter, err := handler.EncounterExecutionService.GetExecutionByUser(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
