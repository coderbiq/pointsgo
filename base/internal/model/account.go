package model

import (
	"time"

	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/common"
)

// Account 定义不使用 CQRS 的积分账户外观
type Account interface {
	common.Account

	DepositedPoints() common.Points
	ConsumedPoints() common.Points
	CreatedAt() time.Time
	UpdatedAt() time.Time
	Version() uint
}

// AccountRepository 定义积分账户资源库
type AccountRepository interface {
	Save(account Account) error
	Get(accountID vo.LongID) (Account, error)
}

type account struct {
	common.BaseAccount
	AccountReadModel
	events *devent.Recorder
}

// RegisterAccount 为指定会员标识的会员注册一个新的积分账户
func RegisterAccount(ownerID vo.StringID) Account {
	a := &account{events: devent.NewRecorder(0)}
	a.Identity = vo.IDGenerator.LongID()
	a.OwnerIdentity = ownerID
	a.DepPoints = common.Points(0)
	a.ConPoints = common.Points(0)
	a.Created = time.Now()
	a.Updated = time.Now()
	a.events.RecordThan(common.OccurAccountCreated(a.Identity, ownerID))
	return a
}

// AccountFromDatas 根据原始数据重建积分账户模型
// 资源库可以利用这个方法将从数据库获取到的数据还原为聚合模型
func AccountFromDatas(datas map[string]interface{}) Account {
	a := &account{events: devent.NewRecorder(datas["version"].(uint))}
	a.Identity = vo.LongID(datas["id"].(int64))
	a.OwnerIdentity = vo.StringID(datas["ownerId"].(string))
	a.CurPoints = common.Points(datas["points"].(uint))
	a.DepPoints = common.Points(datas["deposited"].(uint))
	a.ConPoints = common.Points(datas["consumed"].(uint))
	a.Created = time.Unix(datas["created"].(int64), 0)
	a.Updated = time.Unix(datas["updated"].(int64), 0)
	return a
}

func (a *account) Deposit(points common.Points) {
	a.CurPoints = a.CurPoints.Inc(points)
	a.DepPoints = a.DepPoints.Inc(points)
	a.Updated = time.Now()
	a.events.RecordThan(common.OccurDeposited(a.Identity, points))
}

func (a *account) Consume(points common.Points) error {
	if !a.CurPoints.GreaterThan(points) {
		return common.PointsInsufficientError{
			Current: a.CurPoints,
			Expect:  points,
		}
	}
	a.CurPoints = a.CurPoints.Dec(points)
	a.ConPoints = a.ConPoints.Inc(points)
	a.Updated = time.Now()
	a.events.RecordThan(common.OccurConsumed(a.Identity, points))
	return nil
}

// Version 返回积分账户聚合的当前版本
func (a *account) Version() uint {
	return a.events.LastVersion()
}

func (a *account) CommitEvents(publishers ...devent.Publisher) {
	a.events.CommitToPublisher(publishers...)
}

// AccountReadModel 积分展示读模型的实现
type AccountReadModel struct {
	DepPoints common.Points `json:"depositedPoints"`
	ConPoints common.Points `json:"consumedPoints"`
	Created   time.Time     `json:"createdAt"`
	Updated   time.Time     `json:"updatedAt"`
}

// DepositedPoints 返回积分账户历史充值总额
func (arm AccountReadModel) DepositedPoints() common.Points {
	return arm.DepPoints
}

// ConsumedPoints 返回积分账户历史消息总额
func (arm AccountReadModel) ConsumedPoints() common.Points {
	return arm.ConPoints
}

// CreatedAt 返回积分账户创建时间
func (arm AccountReadModel) CreatedAt() time.Time {
	return arm.Created
}

// UpdatedAt 返回积分账户最后变更时间
func (arm AccountReadModel) UpdatedAt() time.Time {
	return arm.Updated
}
