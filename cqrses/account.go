package cqrses

import (
	"github.com/coderbiq/dgo/eventsourcing"
	"github.com/coderbiq/dgo/model"
	"github.com/coderbiq/pointsgo/common"
)

type account struct {
	common.BaseAccount
	events *eventsourcing.EventRecorder
}

// RegisterAccount 为指定会员标识的会员注册一个新的积分账户
func RegisterAccount(ownerID model.StringID) common.Account {
	a := new(account)
	a.events = eventsourcing.EventRecorderFromSourced(a, 0)
	a.events.RecordThan(common.OccurAccountCreated(a.Identity, ownerID))
	return a
}

func (a *account) Deposit(points common.Points) {
	a.events.RecordThan(common.OccurDeposited(a.Identity, points))
}

func (a *account) Consume(points common.Points) error {
	if !a.CurPoints.GreaterThan(points) {
		return common.PointsInsufficientError{
			Current: a.CurPoints,
			Expect:  points,
		}
	}
	a.events.RecordThan(common.OccurConsumed(a.Identity, points))
	return nil
}

func (a *account) Apply(event model.DomainEvent) {
	switch e := event.(type) {
	case common.AccountCreated:
		a.Identity = e.AggregateID().(model.LongID)
		a.OwnerIdentity = e.OwnerID()
		break
	case common.AccountDeposited:
		points := e.Points()
		a.CurPoints = a.CurPoints.Inc(points)
		break
	case common.AccountConsumed:
		points := e.Points()
		a.CurPoints = a.CurPoints.Dec(points)
		break
	}
}
