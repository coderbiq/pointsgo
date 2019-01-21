package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/coderbiq/dgo/base/devent"

	"github.com/coderbiq/pointsgo/app"
	"github.com/coderbiq/pointsgo/base/internal/api"
	"github.com/coderbiq/pointsgo/base/internal/infra"
	"github.com/coderbiq/pointsgo/base/internal/model"
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

	regOut := new(app.RegisterResult)
	panicOrNil(json.Unmarshal(data, regOut))

	// 为账户充值 100 积分
	depURL := strings.Replace(app.DepositRoute, "{accountId}", strconv.FormatInt(regOut.AccountId, 10), 1)
	depIn := app.DepositInput{
		AccountId: regOut.AccountId,
		Points:    uint32(100),
	}
	post(depURL, depIn)

	conURL := strings.Replace(app.ConsumeRoute, "{accountId}", strconv.FormatInt(regOut.AccountId, 10), 1)
	conIn := app.ConsumeInput{AccountId: regOut.AccountId, Points: uint32(40)}
	post(conURL, conIn)

	detailURL := strings.Replace(app.DetailRoute, "{accountId}", strconv.FormatInt(regOut.AccountId, 10), 1)
	data = get(detailURL)
	detail := new(app.FindResult)
	panicOrNil(json.Unmarshal(data, detail))

	fmt.Print("\n\n-------------打印账户详情----------------------\n\n")
	fmt.Println("账户标识：", detail.AccountId)
	fmt.Println("所属会员：", detail.CustomerId)
	fmt.Println("可用积分：", detail.Points)
	fmt.Println("总充值积分：", detail.Deposited)
	fmt.Println("总消费积分：", detail.Consumed)
	fmt.Println("开通时间：", detail.Created)
	fmt.Println("操作记录：")
	for _, log := range detail.Logs {
		fmt.Printf("	- [%d] %s %s \n", log.Created, log.Action, log.Desc)
	}

	cancel()
}

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c

	i := infra.NewInfra()
	service := model.NewAppServices(i)
	go service.RunTasks(ctx)

	go i.EventBus().(runner).Run(ctx)
	printLogs(i.EventBus())

	ws := api.WebService(service)
	restful.Add(ws)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
}

type runner interface {
	Run(context.Context)
}

func printLogs(eventBus devent.Bus) {
	eventBus.AddRouter(devent.SimpleRouter(map[string][]devent.Consumer{
		common.AccountDepositedEvent: []devent.Consumer{
			devent.ConsumerFunc(func(event devent.Event) {
				e := event.(common.AccountDeposited)
				fmt.Printf("积分账户充值：账户标识 %s 充值额度 %d \n",
					e.AggregateID().String(), int(e.Points()))
			})},
		common.AccountConsumedEvent: []devent.Consumer{
			devent.ConsumerFunc(func(event devent.Event) {
				e := event.(common.AccountConsumed)
				fmt.Printf("积分消费：账户标识 %s 消费额度 %d \n",
					e.AggregateID().String(), int(e.Points()))
			})},
		common.AccountCreatedEvent: []devent.Consumer{
			devent.ConsumerFunc(func(event devent.Event) {
				e := event.(common.AccountCreated)
				fmt.Printf("新建积分账户：所属会员 %s 账户标识 %s \n",
					e.OwnerID().String(), e.AggregateID().String())
			})},
	}))
}

func get(route string) []byte {
	resp, err := http.Get("http://localhost:8080" + route)
	panicOrNil(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	panicOrNil(err)
	return body

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
