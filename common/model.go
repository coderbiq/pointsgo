package common

import (
	"fmt"
	"time"

	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/dgo/base/vo"
)

type (
	// Points 定义积分点数据模型
	Points uint

	// Account 定义积分账户聚合基础的外观
	Account interface {
		ID() vo.LongID
		OwnerID() vo.StringID
		Points() Points
		Deposit(points Points)
		Consume(points Points) error
	}

	// AccountLog 定义积分账户的变更记录模型
	AccountLog interface {
		ID() vo.Identity
		AccountID() vo.Identity
		Action() string
		Desc() string
		CreatedAt() time.Time
	}
)

// AccountLogFromEvent 根据积分账户的领域模型生成账户变更记录
func AccountLogFromEvent(event devent.DomainEvent) AccountLog {
	switch e := event.(type) {
	case AccountCreated:
		return AccountActionLog{
			Identity:        vo.IDGenerator.LongID(),
			AccountIdentity: e.AggregateID().(vo.LongID),
			ActionName:      "创建账户",
			Describe:        "会员申请开通积分账户",
			Created:         e.CreatedAt(),
		}
	case AccountDeposited:
		return AccountActionLog{
			Identity:        vo.IDGenerator.LongID(),
			AccountIdentity: e.AggregateID().(vo.LongID),
			ActionName:      "积分充值",
			Describe:        fmt.Sprintf("会员为积分账户充值 %d 积分", e.Points()),
			Created:         e.CreatedAt(),
		}
	case AccountConsumed:
		return AccountActionLog{
			Identity:        vo.IDGenerator.LongID(),
			AccountIdentity: e.AggregateID().(vo.LongID),
			ActionName:      "积分消费",
			Describe:        fmt.Sprintf("会员消费积分账户 %d 积分", e.Points()),
			Created:         e.CreatedAt(),
		}
	}
	return nil
}

// AccountActionLog 对 AccountLog 的实现
type AccountActionLog struct {
	Identity        vo.LongID
	AccountIdentity vo.LongID
	ActionName      string
	Describe        string
	Created         time.Time
}

// ID 返回变更记录的唯一标识
func (log AccountActionLog) ID() vo.Identity {
	return log.Identity
}

// AccountID 返回变更记录所属积分账户标识
func (log AccountActionLog) AccountID() vo.Identity {
	return log.AccountIdentity
}

// Action 返回变更记录的操作名称
func (log AccountActionLog) Action() string {
	return log.ActionName
}

// Desc 返回变更记录的详细描述
func (log AccountActionLog) Desc() string {
	return log.Describe
}

// CreatedAt 返回变更记录的创建时间
func (log AccountActionLog) CreatedAt() time.Time {
	return log.Created
}

// BaseAccount 实现基本的积分账户模型
type BaseAccount struct {
	Identity      vo.LongID   `json:"id"`
	OwnerIdentity vo.StringID `json:"ownerId"`
	CurPoints     Points      `json:"points"`
}

// ID 返回积分账户标识
func (a BaseAccount) ID() vo.LongID {
	return a.Identity
}

// OwnerID 返回积分账户所属的会员标识
func (a BaseAccount) OwnerID() vo.StringID {
	return a.OwnerIdentity
}

// Points 返回积分账户中当前的可用积分
func (a BaseAccount) Points() Points {
	return a.CurPoints
}

// Inc 返回当前积分加上一个积分值后的新积分值
func (p Points) Inc(points Points) Points {
	return Points(uint(p) + uint(points))
}

// Dec 返回当前积分值减去一个积分值后的新积分值
func (p Points) Dec(points Points) Points {
	return Points(uint(p) - uint(points))
}

// GreaterThan 返回当前积分值是否比指定积分值大
func (p Points) GreaterThan(other Points) bool {
	return uint(p) > uint(other)
}

// Zero 返回当前积分值是否为零
func (p Points) Zero() bool {
	return p == 0
}
