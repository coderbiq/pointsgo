package base_test

import (
	"testing"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base"
	"github.com/stretchr/testify/suite"
)

func TestRestfulRegister(t *testing.T) {
	suite.Run(t, app.NewRegisterRestfulTestSuite(base.WebService()))
}
