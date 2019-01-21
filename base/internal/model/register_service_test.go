package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coderbiq/pointsgo/base/internal/mocks"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/common"
	"github.com/golang/mock/gomock"
)

func TestRegisterService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ownerId := "testCustomerId"

	assert := assert.New(t)
	repo := mocks.NewMockAccountRepository(ctrl)
	repo.EXPECT().Save(gomock.Any()).Times(1).Do(func(account model.Account) {
		assert.Equal(ownerId, account.OwnerID().String())
	})
	eventBus := mocks.NewMockBus(ctrl)
	eventBus.EXPECT().Publish(gomock.Any()).Times(1).Do(func(event common.AccountCreated) {
		assert.Equal(ownerId, event.OwnerID().String())
		assert.NotEmpty(event.AggregateID())
	})
	infra := mocks.NewMockInfra(ctrl)
	infra.EXPECT().AccountRepo().Return(repo)
	infra.EXPECT().EventBus().Return(eventBus)

	services := model.NewAppServices(infra)
	services.RegisterApp().Register(ownerId)
}
