package server

import (
	"encoding/json"
	"fmt"
	"github.com/d8ml/calculation_server/calc/internal/pkg/calculation"
	"io"
	"log"
	"net/http"
)

type CalcRequest struct {
	Exp string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, errorCode int, errorText string) {
	errorResponse := ErrorResponse{
		Error: errorText,
	}
	data, err := json.Marshal(errorResponse)
	if err != nil {
		_ = fmt.Errorf("failed marshal response: %w", err)
		return
	}
	w.WriteHeader(errorCode)
	_, err = w.Write(data)
	if err != nil {
		_ = fmt.Errorf("failed to write response: %w", err)
	}
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println("Failed to close request body")
		}
	}(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	var calcRequest CalcRequest
	err = json.Unmarshal(body, &calcRequest)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, "Expression is not valid")
		return
	}

	ans, err := calculation.Calculate(calcRequest.Exp)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, "Expression is not valid")
		return
	}

	calculateResponse := CalcResponse{
		Result: ans,
	}

	data, err := json.Marshal(calculateResponse)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Internal server error")
	} else {
		_, err = w.Write(data)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Internal server error")
		}
	}
}
