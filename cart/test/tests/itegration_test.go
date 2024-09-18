package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
	ISsuite "gitlab.ozon.dev/1mikle1/homework/cart/test/suite"
)

func TestIntegrationSuite(t *testing.T) {
	//suite.Run(t, new(test_suite.ItemS))
	suite.Run(t, new(ISsuite.ItemServiceSuite))
}
