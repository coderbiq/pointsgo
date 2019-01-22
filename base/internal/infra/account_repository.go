package infra

import (
	"errors"

	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
)

type inMemoryAccountRepo struct {
}

// NewInMemoryAccountRepo 创建积分账户资源库
func NewInMemoryAccountRepo() model.AccountRepository {
	return new(inMemoryAccountRepo)
}

func (repo inMemoryAccountRepo) Save(account model.Account) error {
	if _, has := db.get(accountKey(account.ID().String())); has {
		return repo.update(account)
	}
	return repo.insert(account)
}

func (repo inMemoryAccountRepo) update(account model.Account) error {
	data, _ := db.get(accountKey(account.ID().String()))
	po := data.(accountPO)
	if po.datas["version"].(uint) >= account.Version() {
		panic(errors.New("并发冲突：希望存储的积分账户已不是最新版本"))
	}

	repo.aggregateToPo(account, &po)
	db.set(accountKey(account.ID().String()), po)
	return nil
}

func (repo inMemoryAccountRepo) insert(account model.Account) error {
	ownerAccounts := []string{}
	if data, has := db.get(ownerKey(account.OwnerID().String())); has {
		ownerAccounts = data.([]string)
	}
	ownerAccounts = append(ownerAccounts, accountKey(account.ID().String()))

	po := accountPO{datas: make(map[string]interface{})}
	repo.aggregateToPo(account, &po)
	po.datas["id"] = account.ID().(vo.LongID).Int64()
	po.datas["ownerId"] = account.OwnerID().String()
	po.datas["created"] = account.CreatedAt().Unix()
	db.set(accountKey(account.ID().String()), po)
	db.set(ownerKey(account.OwnerID().String()), ownerAccounts)
	return nil
}

func (repo inMemoryAccountRepo) Get(accountID vo.LongID) (model.Account, error) {
	if data, has := db.get(accountKey(accountID.String())); has {
		po := data.(accountPO)
		return model.AccountFromDatas(po.datas), nil
	}
	return nil, errors.New("没有找到指定标识的积分账户")
}

func (repo inMemoryAccountRepo) aggregateToPo(account model.Account, po *accountPO) {
	po.datas["points"] = uint(account.Points())
	po.datas["deposited"] = uint(account.DepositedPoints())
	po.datas["consumed"] = uint(account.ConsumedPoints())
	po.datas["updated"] = account.UpdatedAt().Unix()
	po.datas["version"] = account.Version()
}
