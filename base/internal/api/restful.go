package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base/internal/model"
	"github.com/emicklei/go-restful"
)

type restHandlerFunc func(*restful.Request, *restful.Response) error

// WebService 返回 restful 服务
func WebService(services model.AppServices) *restful.WebService {

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST(app.RegisterRoute).To(restHandler(
		func(req *restful.Request, resp *restful.Response) error {
			input := app.RegisterInput{}
			if err := req.ReadEntity(&input); err != nil {
				return errors.New("请求信息错误！")
			}
			accountID, err := services.RegisterApp().Register(input.CustomerId)
			if err != nil {
				return err
			}
			result := app.RegisterResult{AccountId: accountID}
			resp.WriteEntity(result)
			return nil
		})))

	ws.Route(ws.POST(app.DepositRoute).To(restHandler(
		func(req *restful.Request, resp *restful.Response) error {
			accountID, err := strconv.ParseInt(req.PathParameter("accountId"), 10, 0)
			if err != nil {
				return errors.New("请提供正确的积分账户标识")
			}
			input := new(app.DepositInput)
			if err := req.ReadEntity(input); err != nil {
				return errors.New("请求信息错误")
			}
			curPoints, deposited, err := services.DepositApp().Deposit(accountID, uint(input.Points))
			if err != nil {
				return err
			}
			resp.WriteEntity(app.DepositResult{
				CurPoints:       uint32(curPoints),
				DepositedPoints: uint32(deposited),
			})
			return nil
		})))

	ws.Route(ws.POST(app.ConsumeRoute).To(restHandler(
		func(req *restful.Request, resp *restful.Response) error {
			accountID, err := strconv.ParseInt(req.PathParameter("accountId"), 10, 0)
			if err != nil {
				return errors.New("请提供正确的积分账户标识")
			}
			input := new(app.ConsumeInput)
			if err := req.ReadEntity(input); err != nil {
				return errors.New("请求信息错误")
			}
			cur, consumed, err := services.ConsumeApp().Consume(accountID, uint(input.Points))
			if err != nil {
				return err
			}
			resp.WriteEntity(app.ConsumeResult{
				CurPoints:      uint32(cur),
				ConsumedPoints: uint32(consumed),
			})
			return nil
		})))

	ws.Route(ws.GET(app.DetailRoute).To(restHandler(
		func(req *restful.Request, resp *restful.Response) error {
			accountID, err := strconv.ParseInt(req.PathParameter("accountId"), 10, 0)
			if err != nil {
				return errors.New("请提供正确的积分账户标识")
			}
			reader, err := services.Finder().Detail(accountID)
			if err != nil {
				resp.WriteError(http.StatusNotFound, err)
				return nil
			}
			result := app.FindResult{
				AccountId:  reader.ID(),
				CustomerId: reader.OwnerID(),
				Points:     uint32(reader.Points()),
				Deposited:  uint32(reader.Deposited()),
				Consumed:   uint32(reader.Consumed()),
				Logs:       []*app.Log{},
				Created:    reader.CreatedAt().Unix(),
			}
			for _, log := range reader.Logs() {
				result.Logs = append(result.Logs, &app.Log{
					Action:  log.Action(),
					Desc:    log.Desc(),
					Created: log.CreatedAt().Unix(),
				})
			}
			resp.WriteEntity(result)
			return nil
		})))

	return ws
}

func restHandler(handler restHandlerFunc) restful.RouteFunction {
	return func(req *restful.Request, resp *restful.Response) {
		defer func() {
			if e := recover(); e != nil {
				resp.WriteErrorString(
					http.StatusInternalServerError,
					"服务器开了点小差，请稍候再试！")
			}
		}()
		if err := handler(req, resp); err != nil {
			resp.WriteError(http.StatusBadRequest, err)
		}
	}
}
