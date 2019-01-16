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

	customerID := "testCustomerId"

	register := mocks.NewMockRegisterService(ctrl)
	register.EXPECT().Register(customerID).Return(int64(123), nil).Times(1)
	services := mocks.NewMockAppServices(ctrl)
	services.EXPECT().RegisterApp().Return(register)

	suite.Run(t, app.NewRegisterRestfulTestSuite(api.WebService(services), customerID))
}

func TestRestfulDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := app.DepositInput{
		AccountId: int64(123),
		Points:    uint32(321),
	}
	result := app.DepositResult{
		CurPoints:       uint32(400),
		DepositedPoints: uint32(600),
	}

	deposit := mocks.NewMockDepositService(ctrl)
	deposit.EXPECT().
		Deposit(input.AccountId, uint(input.Points)).
		Return(uint(result.CurPoints), uint(result.DepositedPoints), nil).
		Times(1)
	services := mocks.NewMockAppServices(ctrl)
	services.EXPECT().DepositApp().Return(deposit)

	suite.Run(t, app.NewDepositRestfulTestSuite(
		api.WebService(services),
		input,
		result,
	))
}

func TestRestfulConsume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := app.ConsumeInput{AccountId: int64(123), Points: uint32(100)}
	result := app.ConsumeResult{CurPoints: uint32(300), ConsumedPoints: uint32(1000)}

	consue := mocks.NewMockConsumeService(ctrl)
	consue.EXPECT().
		Consume(input.AccountId, uint(input.Points)).
		Return(uint(result.CurPoints), uint(result.ConsumedPoints), nil).
		Times(1)
	service := mocks.NewMockAppServices(ctrl)
	service.EXPECT().ConsumeApp().Return(consue)

	suite.Run(t, app.NewConsumeRestfulTestSuite(
		api.WebService(service),
		input,
		result,
	))
}
