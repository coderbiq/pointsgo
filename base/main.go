package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coderbiq/dgo/base/devent"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base/internal/api"
	"github.com/coderbiq/pointsgo/base/internal/infra"
	"github.com/coderbiq/pointsgo/base/internal/service"
	"github.com/coderbiq/pointsgo/common"
	"github.com/emicklei/go-restful"
)

var (
	cancel context.CancelFunc
)

func main() {
	// 注册一个积分账户
	in := app.RegisterInput{CustomerId: "testCustomerId"}
	data := post(app.RegisterRoute, in)
	fmt.Println(string(data))

	cancel()
}

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c

	i := infra.NewInfra()

	go i.EventBus().(runner).Run(ctx)
	printLogs(i.EventBus())

	service := service.NewAppServices(i)
	ws := api.WebService(service)
	restful.Add(ws)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
}

type runner interface {
	Run(context.Context)
}

func printLogs(eventBus devent.EventBus) {
	eventBus.Listen(common.AccountCreatedEvent,
		devent.EventConsumerFunc(func(event devent.DomainEvent) {
			e := event.(common.AccountCreated)
			fmt.Printf("新建积分账户：所属会员 %s 账户标识 %s \n",
				e.OwnerID().String(), e.AggregateID().String())
		}))
}

func post(route string, d interface{}) []byte {
	data, err := json.Marshal(d)
	panicOrNil(err)
	resp, err := http.Post("http://localhost:8080"+route, "application/json", bytes.NewReader(data))
	panicOrNil(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	panicOrNil(err)
	return body
}

func panicOrNil(err error) {
	if err != nil {
		panic(err)
	}
}
