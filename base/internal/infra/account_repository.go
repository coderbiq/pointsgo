package infra

import (
	"errors"

	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
)

type inMemoryAccountRepo struct {
	accounts map[int64]model.Account
}

// NewInMemoryAccountRepo 创建积分账户资源库
func NewInMemoryAccountRepo() model.AccountRepository {
	return &inMemoryAccountRepo{
		accounts: map[int64]model.Account{},
	}
}

func (repo *inMemoryAccountRepo) Save(account model.Account) error {
	repo.accounts[account.ID().Int64()] = account
	return nil
}

func (repo inMemoryAccountRepo) Get(accountID vo.LongID) (model.Account, error) {
	account, has := repo.accounts[accountID.Int64()]
	if has {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func (repo inMemoryAccountRepo) FindByOwner(ownerID vo.LongID) ([]model.Account, error) {
	accounts := []model.Account{}
	for _, account := range repo.accounts {
		if account.OwnerID().Equal(ownerID) {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}
