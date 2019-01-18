package service

import (
	"context"

	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/coderbiq/pointsgo/common"
)

type (
	// AppServices 定义应用服务容器
	AppServices interface {
		RegisterApp() RegisterService
		DepositApp() DepositService
		ConsumeApp() ConsumeService
		Finder() AccountFinder
		RunTasks(context.Context)
	}
	// Infra 定义基础设施服务容器
	Infra interface {
		AccountRepo() model.AccountRepository
		EventBus() devent.Bus
		LogStorer() model.AccountLogStorer
	}
	// RegisterService 定义积分账户注册服务
	RegisterService interface {
		// Register 为指定会员注册一个积分账户，成功返回新注册的账户标识
		Register(customerID string) (int64, error)
	}
	// DepositService 提供积分账户充值服务
	DepositService interface {
		// Deposit 为指定积分账户充值指定积分，成功返回充值后的账户可用积分以及总充值积分
		Deposit(accountID int64, points uint) (curPoints, deposited uint, err error)
	}
	// ConsumeService 提供积分账户消费积分服务
	ConsumeService interface {
		// Consume 消费指定积分账户下的积分，成功返回消费后的可用积分和总消费积分
		Consume(accountID int64, points uint) (curPoints, consumed uint, err error)
	}
	// AccountFinder 定义积分账户的查询服务
	AccountFinder interface {
		Detail(accountID int64) (common.AccountReader, error)
	}
)

type services struct {
	infra Infra
}

// NewAppServices 创建应用服务容器
func NewAppServices(infra Infra) AppServices {
	return &services{infra: infra}
}

func (ss *services) RegisterApp() RegisterService {
	return &registerService{
		repo:     ss.infra.AccountRepo(),
		eventBus: ss.infra.EventBus()}
}

func (ss *services) DepositApp() DepositService {
	return &depositService{
		repo:     ss.infra.AccountRepo(),
		eventBus: ss.infra.EventBus(),
	}
}

func (ss *services) ConsumeApp() ConsumeService {
	return &consumeService{
		repo:     ss.infra.AccountRepo(),
		eventBus: ss.infra.EventBus(),
	}
}

func (ss *services) Finder() AccountFinder {
	return &finder{
		repo:      ss.infra.AccountRepo(),
		logStorer: ss.infra.LogStorer(),
	}
}

func (ss *services) RunTasks(ctx context.Context) {
	runAccountLogRecorder(ss.infra.EventBus(), ss.infra.LogStorer())
}

type registerService struct {
	repo     model.AccountRepository
	eventBus devent.Publisher
}

func (service registerService) Register(customerID string) (int64, error) {
	account := model.RegisterAccount(vo.StringID(customerID))
	service.repo.Save(account)
	account.(devent.Producer).CommitEvents(service.eventBus)
	return account.ID().(vo.LongID).Int64(), nil
}

type depositService struct {
	repo     model.AccountRepository
	eventBus devent.Publisher
}

func (service depositService) Deposit(accountID int64, points uint) (uint, uint, error) {
	account, err := service.repo.Get(vo.LongID(accountID))
	if err != nil {
		return 0, 0, err
	}
	account.Deposit(common.Points(points))
	service.repo.Save(account)
	account.(devent.Producer).CommitEvents(service.eventBus)
	return uint(account.Points()), uint(account.DepositedPoints()), nil
}

type consumeService struct {
	repo     model.AccountRepository
	eventBus devent.Publisher
}

func (service consumeService) Consume(accountID int64, points uint) (uint, uint, error) {
	account, err := service.repo.Get(vo.LongID(accountID))
	if err != nil {
		return 0, 0, err
	}
	if err := account.Consume(common.Points(points)); err != nil {
		return 0, 0, err
	}
	service.repo.Save(account)
	account.(devent.Producer).CommitEvents(service.eventBus)
	return uint(account.Points()), uint(account.ConsumedPoints()), nil
}

type finder struct {
	repo      model.AccountRepository
	logStorer model.AccountLogStorer
}

func (f finder) Detail(accountID int64) (common.AccountReader, error) {
	account, err := f.repo.Get(vo.LongID(accountID))
	if err != nil {
		return common.AccountReader{}, err
	}
	logs := f.logStorer.Get(account.ID())
	logData := []map[string]interface{}{}
	for _, log := range logs {
		logData = append(logData, map[string]interface{}{
			"action":  log.Action(),
			"desc":    log.Desc(),
			"created": log.CreatedAt(),
		})
	}
	reader := common.AccountReaderFromData(map[string]interface{}{
		"id":        account.ID().(vo.LongID).Int64(),
		"ownerId":   account.OwnerID().String(),
		"points":    uint(account.Points()),
		"deposited": uint(account.DepositedPoints()),
		"consumed":  uint(account.ConsumedPoints()),
		"logs":      logData,
		"created":   account.CreatedAt(),
	})
	return reader, nil
}

// runAccountLogRecorder 启动积分账户日志记录器
func runAccountLogRecorder(bus devent.Bus, storer model.AccountLogStorer) {
	bus.AddRouter(devent.RegexRouter(map[string][]devent.Consumer{
		"account.*": []devent.Consumer{
			devent.ConsumerFunc(func(e devent.Event) {
				log := common.AccountLogFromEvent(e)
				storer.Append(log)
			}),
		},
	}))
}
