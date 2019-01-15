package base

import (
	"github.com/coderbiq/pointsgo/app"
	"github.com/emicklei/go-restful"
)

// WebService 返回 restful 服务
func WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST(app.RegisterRoute).
		To(func(req *restful.Request, resp *restful.Response) {
			result := app.RegisterResult{AccountId: 123}
			resp.WriteEntity(result)
		}))

	return ws
}
