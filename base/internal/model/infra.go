package model

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/common"
)

type (
	// Infra 定义基础设施服务容器
	Infra interface {
		AccountRepo() AccountRepository
		EventBus() devent.Bus
		LogStorer() AccountLogStorer
	}

	// AccountLogStorer 定义积分账户日志存储器
	AccountLogStorer interface {
		Append(log common.AccountLog)
		Get(accountID vo.Identity) []common.AccountLog
	}
)
