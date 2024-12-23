package test

import (
	"github.com/d8ml/calculation_server/calc/internal/app/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestHandlerOk(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculate",
		strings.NewReader(`{"expression": "2 + 3"}`))
	w := httptest.NewRecorder()
	server.Calculate(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}
}

func TestRequestHandlerInvalidExpression(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculate",
		strings.NewReader(`{"expression": "2 / 0"}`))
	w := httptest.NewRecorder()
	server.Calculate(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("wrong status code")
	}
}
