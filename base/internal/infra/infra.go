package infra

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/base/internal/service"
)

type infra struct {
	repo     model.AccountRepository
	eventBus devent.EventBus
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

func (i *infra) EventBus() devent.EventBus {
	if i.eventBus == nil {
		i.eventBus = devent.SimpleEventBus(10)
	}
	return i.eventBus
}
