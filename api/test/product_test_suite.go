package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/yakiza/Zephyros/api"
	"github.com/yakiza/Zephyros/internal/test_data"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ProductTestSuite struct {
	suite.Suite
}

func (suite *ProductTestSuite) TestAddKiteSuccess(t *testing.T) {
	x, err := json.Marshal(test_data.VorasTestKite())
	suite.NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/kite", bytes.NewReader(x))
	suite.NoError(err)

	w := httptest.NewRecorder()
	handler := api.NewRouter()
	handler.ServeHTTP(w, request)
	suite.Equal(w.Code, http.StatusOK)

}

func (suite *ProductTestSuite) TestAddKiteFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/kite", nil)
	w := httptest.NewRecorder()
	handler := api.NewRouter()
	handler.ServeHTTP(w, request)

	suite.Equal(w.Code, http.StatusBadRequest)
}

func (suite *ProductTestSuite) TestAddKiteMalformedJSON(t *testing.T) {
	malformedJSON := "{{}"
	request := httptest.NewRequest(http.MethodPost, "/kite", bytes.NewReader([]byte(malformedJSON)))
	w := httptest.NewRecorder()
	handler := api.NewRouter()
	handler.ServeHTTP(w, request)

	suite.Equal(w.Code, http.StatusBadRequest)

}

func NewTestProductSuite() *ProductTestSuite {
	return &ProductTestSuite{}
}
