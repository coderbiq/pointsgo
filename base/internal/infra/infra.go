package infra

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/base/internal/service"
)

type infra struct {
	repo     model.AccountRepository
	eventBus devent.Bus
}

// NewInfra 返回基础设施服务容器
func NewInfra() service.Infra {
	return new(infra)
}

func (i *infra) AccountRepo() model.AccountRepository {
	if i.repo == nil {
		i.repo = NewInMemoryAccountRepo()
	}
	return i.repo
}

func (i *infra) EventBus() devent.Bus {
	if i.eventBus == nil {
		i.eventBus = devent.SimpleBus(10)
	}
	return i.eventBus
}
