package infra

import (
	"encoding/json"
	"errors"

	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/common"
)

type accountLogStorer struct {
	logs map[string][][]byte
}

// NewAccountLogStorer 返回积分账户日志存储器
func NewAccountLogStorer() model.AccountLogStorer {
	return &accountLogStorer{
		logs: map[string][][]byte{},
	}
}

func (storer *accountLogStorer) Append(log common.AccountLog) {
	logs, has := storer.logs[log.AccountID().String()]
	if !has {
		logs = [][]byte{}
	}
	data, err := json.Marshal(log)
	if err != nil {
		panic(errors.New("序列化积分日志异常： " + err.Error()))
	}
	logs = append(logs, data)
	storer.logs[log.AccountID().String()] = logs
}

func (storer *accountLogStorer) Get(accountID vo.LongID) []common.AccountLog {
	logs := []common.AccountLog{}
	if datas, has := storer.logs[accountID.String()]; has {
		for _, data := range datas {
			log := new(common.AccountActionLog)
			if err := json.Unmarshal(data, log); err != nil {
				panic(errors.New("解析积分日志异常：" + err.Error()))
			}
			logs = append(logs, log)
		}
	}
	return logs
}
