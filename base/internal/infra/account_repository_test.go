package infra

import (
	"fmt"
	"testing"

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

	_, hasAccount := db.get(suite.repo.accountKey(account.ID()))
	suite.True(hasAccount)
	_, hasOwner := db.get(suite.repo.ownerKey(account.OwnerID()))
	suite.True(hasOwner)
}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(accountRepositoryTestSuite))
}
