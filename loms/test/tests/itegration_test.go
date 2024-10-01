package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
	LOMSsuite "gitlab.ozon.dev/1mikle1/homework/loms/test/suite"
)

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(LOMSsuite.LOMSServiceSuite))
}
