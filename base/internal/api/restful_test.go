package api_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base/internal/api"
	"github.com/coderbiq/pointsgo/base/internal/mocks"
	"github.com/stretchr/testify/suite"
)

func TestRestfulRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	register := mocks.NewMockRegisterService(ctrl)
	register.EXPECT().Register(gomock.Any()).Return(int64(123), nil).Times(1)
	services := mocks.NewMockAppServices(ctrl)
	services.EXPECT().RegisterApp().Return(register)

	suite.Run(t, app.NewRegisterRestfulTestSuite(api.WebService(services)))
}
