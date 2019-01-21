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
}

// NewInMemoryAccountRepo 创建积分账户资源库
func NewInMemoryAccountRepo() model.AccountRepository {
	return new(inMemoryAccountRepo)
}

func (repo *inMemoryAccountRepo) Save(account model.Account) error {
	if _, has := db.get(repo.accountKey(account.ID())); has {
		return repo.update(account)
	}
	return repo.insert(account)
}

func (repo *inMemoryAccountRepo) update(account model.Account) error {
	data, _ := db.get(repo.accountKey(account.ID()))
	po := data.(accountPO)
	if po.version >= account.Version() {
		panic(errors.New("并发冲突：希望存储的积分账户已不是最新版本"))
	}

	accountData, err := json.Marshal(account)
	if err != nil {
		panic(errors.New("序列化积分账户异常：" + err.Error()))
	}
	po.data = accountData
	po.version = account.Version()
	db.set(repo.accountKey(account.ID()), po)
	return nil
}

func (repo *inMemoryAccountRepo) insert(account model.Account) error {
	ownerAccounts := []string{}
	if data, has := db.get(repo.ownerKey(account.OwnerID())); has {
		ownerAccounts = data.([]string)
	}
	ownerAccounts = append(ownerAccounts, repo.accountKey(account.ID()))
	data, err := json.Marshal(account)
	if err != nil {
		panic(errors.New("序列化积分账户异常：" + err.Error()))
	}
	db.set(repo.accountKey(account.ID()), accountPO{
		ownerID: account.OwnerID().String(),
		version: account.Version(),
		data:    data,
	})
	db.set(repo.ownerKey(account.OwnerID()), ownerAccounts)
	return nil
}

func (repo inMemoryAccountRepo) Get(accountID vo.LongID) (model.Account, error) {
	if data, has := db.get(repo.accountKey(accountID)); has {
		po := data.(accountPO)
		return model.AccountFromData(po.data, po.version), nil
	}
	return nil, errors.New("没有找到指定标识的积分账户")
}

func (repo inMemoryAccountRepo) FindByOwner(ownerID vo.LongID) ([]model.Account, error) {
	accounts := []model.Account{}

	if keys, has := db.get(repo.ownerKey(ownerID)); has {
		accountKeys := keys.([]string)
		for _, key := range accountKeys {
			data, _ := db.get(key)
			po := data.(accountPO)
			accounts = append(accounts, model.AccountFromData(po.data, po.version))
		}
	}
	return []model.Account{}, nil
}

func (repo inMemoryAccountRepo) accountKey(id vo.Identity) string {
	return "account." + id.String()
}

func (repo inMemoryAccountRepo) ownerKey(id vo.Identity) string {
	return "owner." + id.String() + ".accounts"
}
