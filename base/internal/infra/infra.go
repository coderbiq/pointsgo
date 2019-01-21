package infra

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/pointsgo/base/internal/model"
)

type infra struct {
	eventBus devent.Bus
}

// NewInfra 返回基础设施服务容器
func NewInfra() model.Infra {
	return &infra{
		eventBus: devent.SimpleBus(5),
	}
}

func (i infra) AccountRepo() model.AccountRepository {
	return inMemoryAccountRepo{}
}

func (i infra) EventBus() devent.Bus {
	return i.eventBus
}

func (i infra) LogStorer() model.AccountLogStorer {
	return accountLogStorer{}
}

func (i infra) Finder() model.AccountFinder {
	return accountFinder{}
}
