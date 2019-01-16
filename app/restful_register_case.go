package app

import (
	"encoding/json"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/suite"
)

type registerRestfulTestSuite struct {
	suite.Suite
	ws         *restful.WebService
	customerID string
}

// NewRegisterRestfulTestSuite 创建一个通过 restful 接口注册积分账户的测试用例
func NewRegisterRestfulTestSuite(ws *restful.WebService, customerID string) suite.TestingSuite {
	return &registerRestfulTestSuite{ws: ws, customerID: customerID}
}

func (suite *registerRestfulTestSuite) SetupTest() {
	restful.Add(suite.ws)
}

func (suite *registerRestfulTestSuite) TestRegister() {
	input := RegisterInput{CustomerId: suite.customerID}
	resp := request(http.MethodPost, RegisterRoute, input)
	suite.Equal(http.StatusOK, resp.Code, resp.Body.String())

	result := new(RegisterResult)
	suite.Empty(json.Unmarshal(resp.Body.Bytes(), result))
	suite.NotEmpty(result.GetAccountId())
}
