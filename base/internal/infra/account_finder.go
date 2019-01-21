package infra

import (
	"strconv"

	"github.com/coderbiq/pointsgo/base/internal/model"
)

type accountFinder struct {
}

func (finder accountFinder) Detail(accountID int64) (model.AccountReader, error) {
	if accountData, has := db.get(accountKey(strconv.FormatInt(accountID, 10))); has {
		po := accountData.(accountPO)
		datas := po.datas
		datas["logs"] = po.logs
		return model.AccountReaderFromData(datas), nil
	}

	return model.AccountReader{}, nil
}
