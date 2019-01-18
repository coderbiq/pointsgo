package infra

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/base/internal/service"
)

type infra struct {
	repo      model.AccountRepository
	eventBus  devent.Bus
	logStorer model.AccountLogStorer
}

// NewInfra 返回基础设施服务容器
func NewInfra() service.Infra {
	return &infra{
		repo:      NewInMemoryAccountRepo(),
		eventBus:  devent.SimpleBus(10),
		logStorer: NewAccountLogStorer(),
	}
}

func (i *infra) AccountRepo() model.AccountRepository {
	return i.repo
}

func (i *infra) EventBus() devent.Bus {
	return i.eventBus
}

func (i *infra) LogStorer() model.AccountLogStorer {
	return i.logStorer
}
