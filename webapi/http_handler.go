package webapi

import (
	"encoding/json"
	"log"
	"ncbn/infrastucture"
	"ncbn/types"
	"net/http"
)

type WebAPI struct {
	DBService *infrastucture.DatabaseService
}

func (api *WebAPI) HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody types.RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, _ := api.DBService.CallStoredProcedure(requestBody.Procedure)
	if err != nil {
		http.Error(w, "Failed to execute the stored procedure", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal the results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api *WebAPI) executeStoredProcedure(procedure string, params ...interface{}) *string {
	queryResult, err := api.DBService.CallStoredProcedure(procedure)
	if err != nil {
		log.Fatal(err)
	}

	if queryResult == nil {
		log.Printf("Result is empty.")
		return nil
	}

	log.Printf(*queryResult)
	return queryResult
}
