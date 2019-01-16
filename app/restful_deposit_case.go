package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/suite"
)

type depositRestfulTestSuite struct {
	suite.Suite
	restfulTesting
	input  DepositInput
	result DepositResult
}

// NewDepositRestfulTestSuite 返回通过 rest 接口充值积分账户的测试
func NewDepositRestfulTestSuite(ws *restful.WebService, input DepositInput, result DepositResult) suite.TestingSuite {
	return &depositRestfulTestSuite{
		restfulTesting: restfulTesting{ws: ws},
		input:          input,
		result:         result,
	}
}

func (suite *depositRestfulTestSuite) TestDeposit() {
	url := strings.Replace(
		DepositRoute,
		"{accountId}",
		strconv.FormatInt(suite.input.AccountId, 10),
		1)
	fmt.Println(url)
	resp := suite.request(http.MethodPost, url, suite.input)
	if !suite.Equal(http.StatusOK, resp.Code, resp.Body.String()) {
		suite.FailNow("请求充值接口报错")
	}

	result := DepositResult{}
	suite.Empty(json.Unmarshal(resp.Body.Bytes(), &result))
	suite.Equal(suite.result, result, resp.Body.String())
}
