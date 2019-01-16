package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/suite"
)

type consumeRestfulTestSutie struct {
	suite.Suite
	restfulTesting
	input  ConsumeInput
	result ConsumeResult
}

// NewConsumeRestfulTestSuite 返回通过 rest 接口消费积分的测试
func NewConsumeRestfulTestSuite(ws *restful.WebService, input ConsumeInput, result ConsumeResult) suite.TestingSuite {
	return &consumeRestfulTestSutie{
		restfulTesting: restfulTesting{ws: ws},
		input:          input,
		result:         result,
	}
}

func (suite consumeRestfulTestSutie) TestConsume() {
	url := strings.Replace(
		ConsumeRoute,
		"{accountId}",
		strconv.FormatInt(suite.input.AccountId, 10),
		1)
	resp := suite.request(http.MethodPost, url, suite.input)
	if !suite.Equal(http.StatusOK, resp.Code) {
		suite.FailNow("请求积分消费接口报错： %s", resp.Body.String())
	}
	result := ConsumeResult{}
	suite.Empty(json.Unmarshal(resp.Body.Bytes(), &result))
	suite.Equal(suite.result, result)
}
