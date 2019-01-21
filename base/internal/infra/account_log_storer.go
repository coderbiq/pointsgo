package infra

import (
	"encoding/json"
	"errors"

	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/common"
)

type accountLogStorer struct {
}

// NewAccountLogStorer 返回积分账户日志存储器
func NewAccountLogStorer() model.AccountLogStorer {
	return new(accountLogStorer)
}

func (storer *accountLogStorer) Append(log common.AccountLog) {
	logs := [][]byte{}
	if data, has := db.get(storer.key(log.AccountID())); has {
		logs = data.([][]byte)
	}
	data, err := json.Marshal(log)
	if err != nil {
		panic(errors.New("序列化积分日志异常： " + err.Error()))
	}
	logs = append(logs, data)
	db.set(storer.key(log.AccountID()), logs)
}

func (storer *accountLogStorer) Get(accountID vo.Identity) []common.AccountLog {
	logs := []common.AccountLog{}
	if datas, has := db.get(storer.key(accountID)); has {
		bytes := datas.([][]byte)
		for _, data := range bytes {
			log := new(common.AccountActionLog)
			if err := json.Unmarshal(data, log); err != nil {
				panic(errors.New("解析积分日志异常：" + err.Error()))
			}
			logs = append(logs, log)
		}
	}
	return logs
}

func (storer *accountLogStorer) key(accountID vo.Identity) string {
	return "account." + accountID.String() + ".logs"
}
