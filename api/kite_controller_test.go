package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/yakiza/Zephyros/api"
	"github.com/yakiza/Zephyros/internal/test_data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddKiteSuccess(t *testing.T) {
	x, err := json.Marshal(test_data.VorasTestKite())
	if err != nil {
		t.Errorf(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "/kite", bytes.NewReader(x))
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	handler := api.NewRouter()
	handler.ServeHTTP(w, request)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAddKiteFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/kite", nil)
	w := httptest.NewRecorder()
	handler := api.NewRouter()
	handler.ServeHTTP(w, request)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
