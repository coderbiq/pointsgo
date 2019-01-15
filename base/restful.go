package base

import (
	"errors"
	"net/http"

	"github.com/coderbiq/pointsgo/app"
	"github.com/emicklei/go-restful"
)

type restHandlerFunc func(*restful.Request, *restful.Response) error

// WebService 返回 restful 服务
func WebService(services AppServices) *restful.WebService {

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
