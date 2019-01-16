package basecqrs

import (
	"github.com/coderbiq/dgo/base/devent"
	"github.com/coderbiq/dgo/base/vo"
	"github.com/coderbiq/pointsgo/common"
)

type account struct {
	common.BaseAccount
	events *devent.EventRecorder
}

// RegisterAccount 为指定会员标识的会员注册一个新的积分账户
func RegisterAccount(ownerID vo.StringID) common.Account {
	a := &account{events: devent.NewEventRecorder(0)}
	a.Identity = vo.IDGenerator.LongID()
	a.OwnerIdentity = ownerID
	a.events.RecordThan(common.OccurAccountCreated(a.Identity, ownerID))
	return a
}

func (a *account) Deposit(points common.Points) {
	a.CurPoints = a.CurPoints.Inc(points)
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
	a.events.RecordThan(common.OccurConsumed(a.Identity, points))
	return nil
}
