package infra

import (
	"encoding/json"
	"errors"

	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
)

// accountPO 积分账户的持久化模型，持久化模型用于解决模型与持久化存储基础设施的关系。
//
// 持久化模型与领域模型并不一定是一对一映射，如果一个领域模型对应的是关系数据库里的多张表，
// 并且持久化模型使用的是 ORM 实现，那么一个领域模型就可能对应多个持久化模型。在这种情况下资源
// 库在进行持久化模型的读写的时候需要维护数据的事务一致性。
type accountPO struct {
	ownerID string
	version uint
	data    []byte
}

type inMemoryAccountRepo struct {
	accounts map[int64]accountPO
}

// NewInMemoryAccountRepo 创建积分账户资源库
func NewInMemoryAccountRepo() model.AccountRepository {
	return &inMemoryAccountRepo{
		accounts: map[int64]accountPO{},
	}
}

func (repo *inMemoryAccountRepo) Save(account model.Account) error {
	if po, has := repo.accounts[account.ID().Int64()]; has {
		if po.version < account.Version() {
			panic(errors.New("并发冲突：希望存储的积分账户已不是最新版本"))
		}
	}
	data, err := json.Marshal(account)
	if err != nil {
		panic(errors.New("序列化积分账户异常：" + err.Error()))
	}
	repo.accounts[account.ID().Int64()] = accountPO{
		ownerID: account.OwnerID().String(),
		version: account.Version(),
		data:    data,
	}
	return nil
}

func (repo inMemoryAccountRepo) Get(accountID vo.LongID) (model.Account, error) {
	po, has := repo.accounts[accountID.Int64()]
	if !has {
		return nil, errors.New("没有找到指定标识的积分账户")
	}
	return model.AccountFromData(po.data, po.version), nil
}

func (repo inMemoryAccountRepo) FindByOwner(ownerID vo.LongID) ([]model.Account, error) {
	accounts := []model.Account{}
	for _, po := range repo.accounts {
		if ownerID.Equal(vo.StringID(po.ownerID)) {
			accounts = append(accounts, model.AccountFromData(po.data, po.version))
		}
	}
	return accounts, nil
}
