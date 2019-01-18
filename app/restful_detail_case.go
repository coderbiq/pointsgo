package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/suite"
)

type detailRestfulTestSuite struct {
	suite.Suite
	restfulTesting
	accountID int64
	result    FindResult
}

// NewDetailRestfulTestSuite 返回通过 rest 接口获取账户详情的测试
func NewDetailRestfulTestSuite(ws *restful.WebService, accountID int64, result FindResult) suite.TestingSuite {
	return &detailRestfulTestSuite{
		restfulTesting: restfulTesting{ws: ws},
		accountID:      accountID,
		result:         result,
	}
}

func (suite detailRestfulTestSuite) TestDetail() {
	url := strings.Replace(
		DetailRoute,
		"{accountId}",
		strconv.FormatInt(suite.accountID, 10),
		1)
	resp := suite.request(http.MethodGet, url, nil)
	if !suite.Equal(http.StatusOK, resp.Code, resp.Body.String()) {
		suite.FailNow("请求充值接口报错")
	}

	result := FindResult{}
	suite.Empty(json.Unmarshal(resp.Body.Bytes(), &result))
	suite.Equal(suite.result, result, resp.Body.String())
}
