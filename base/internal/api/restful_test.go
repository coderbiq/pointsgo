package api_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base/internal/api"
	"github.com/coderbiq/pointsgo/base/internal/mocks"
	"github.com/coderbiq/pointsgo/base/internal/model"
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

func TestRestfulDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountID := int64(123)
	logDatas := []map[string]interface{}{
		map[string]interface{}{
			"action":  "created",
			"desc":    "会员创建了账户",
			"created": time.Now(),
		},
	}
	datas := map[string]interface{}{
		"id":        accountID,
		"ownerId":   "testCustomerId",
		"points":    uint(100),
		"deposited": uint(200),
		"consumed":  uint(150),
		"logs":      logDatas,
		"created":   time.Now(),
	}
	reader := model.AccountReaderFromData(datas)
	result := app.FindResult{
		AccountId:  reader.ID(),
		CustomerId: reader.OwnerID(),
		Points:     uint32(reader.Points()),
		Deposited:  uint32(reader.Deposited()),
		Consumed:   uint32(reader.Consumed()),
		Logs: []*app.Log{
			&app.Log{
				Action:  reader.Logs()[0].Action(),
				Desc:    reader.Logs()[0].Desc(),
				Created: reader.Logs()[0].CreatedAt().Unix(),
			},
		},
		Created: reader.CreatedAt().Unix(),
	}

	finder := mocks.NewMockAccountFinder(ctrl)
	finder.EXPECT().Detail(accountID).Return(reader, nil).Times(1)
	service := mocks.NewMockAppServices(ctrl)
	service.EXPECT().Finder().Return(finder)

	suite.Run(t, app.NewDetailRestfulTestSuite(
		api.WebService(service), accountID, result))
}
