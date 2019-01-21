package infra

import (
	"errors"
	"time"

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

func (storer accountLogStorer) Append(log common.AccountLog) {
	data, has := db.get(accountKey(log.AccountID().String()))
	if !has {
		panic(errors.New("添加账户日志异常：没有找到指定账户"))
	}
	po := data.(accountPO)
	po.logs = append(po.logs, map[string]interface{}{
		"id":        log.ID().(vo.LongID).Int64(),
		"accountId": log.AccountID().(vo.LongID).Int64(),
		"action":    log.Action(),
		"desc":      log.Desc(),
		"created":   log.CreatedAt().Unix(),
	})
	db.set(accountKey(log.AccountID().String()), po)
}

func (storer accountLogStorer) Get(accountID vo.Identity) []common.AccountLog {
	logs := []common.AccountLog{}

	data, has := db.get(accountKey(accountID.String()))
	if !has {
		return logs
	}
	po := data.(accountPO)
	for _, log := range po.logs {
		logs = append(logs, common.AccountActionLog{
			Identity:        vo.LongID(log["id"].(int64)),
			AccountIdentity: vo.LongID(log["accountId"].(int64)),
			ActionName:      log["action"].(string),
			Describe:        log["desc"].(string),
			Created:         time.Unix(log["created"].(int64), 0),
		})
	}
	return logs
}
