package api_test

// TODO: ADD A DUBLICATE PRODUCT CODE 409 TEST

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
	handler http.Handler
}

func (suite *ProductTestSuite) SetupTest() {
	suite.handler = api.NewRouter()
}

func (suite *ProductTestSuite) TestAddKiteSuccess() {
	data, err := json.Marshal(test_data.VorasTestKite())
	suite.NoError(err)

	request := httptest.NewRequest(http.MethodPost, api.ProductMountPoint, bytes.NewReader(data))
	w := httptest.NewRecorder()
	suite.handler.ServeHTTP(w, request)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *ProductTestSuite) TestAddKiteFailed() {
	request := httptest.NewRequest(http.MethodPost, api.ProductMountPoint, nil)
	w := httptest.NewRecorder()
	suite.handler.ServeHTTP(w, request)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductTestSuite) TestAddKiteMalformedJSON() {
	malformedJSON := "{{}"
	request := httptest.NewRequest(http.MethodPost, api.ProductMountPoint, bytes.NewReader([]byte(malformedJSON)))
	w := httptest.NewRecorder()
	suite.handler.ServeHTTP(w, request)

	suite.Equal(http.StatusBadRequest, w.Code)

}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
