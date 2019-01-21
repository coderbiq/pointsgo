package infra

import (
	"fmt"
	"testing"

	"github.com/coderbiq/dgo/base/vo"

	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/stretchr/testify/suite"
)

type accountRepositoryTestSuite struct {
	suite.Suite
	repo *inMemoryAccountRepo
}

func (suite *accountRepositoryTestSuite) SetupTest() {
	suite.repo = new(inMemoryAccountRepo)
}

func (suite *accountRepositoryTestSuite) TestSave() {
	account := model.RegisterAccount("testCustomerId")
	suite.repo.Save(account)
	fmt.Println(db)

	_, hasAccount := db.get(accountKey(account.ID().String()))
	suite.True(hasAccount)
	_, hasOwner := db.get(ownerKey(account.OwnerID().String()))
	suite.True(hasOwner)

	account2, _ := suite.repo.Get(account.ID().(vo.LongID))
	suite.Equal(account.ID(), account2.ID())
	suite.Equal(account.Version(), account2.Version())
}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(accountRepositoryTestSuite))
}
