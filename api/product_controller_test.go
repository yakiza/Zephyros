package api_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/yakiza/Zephyros/api/test"
	"testing"
)

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, test.NewTestProductSuite())
}
