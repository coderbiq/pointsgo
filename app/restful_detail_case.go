package app

import (
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
	fields    []string
	result    string
}

// NewDetailRestfulTestSuite 返回通过 rest 接口获取账户详情的测试
func NewDetailRestfulTestSuite(
	ws *restful.WebService,
	accountID int64,
	fields []string,
	result string) suite.TestingSuite {
	return &detailRestfulTestSuite{
		restfulTesting: restfulTesting{ws: ws},
		accountID:      accountID,
		fields:         fields,
		result:         result,
	}
}

func (suite detailRestfulTestSuite) TestDetail() {
	url := strings.Replace(
		DetailRoute,
		"{accountId}",
		strconv.FormatInt(suite.accountID, 10),
		1)
	url = url + "?fields=" + strings.Join(suite.fields, ",")
	resp := suite.request(http.MethodGet, url, nil)
	if !suite.Equal(http.StatusOK, resp.Code, resp.Body.String()) {
		suite.FailNow("请求充值接口报错")
	}

	suite.JSONEq(suite.result, resp.Body.String(), resp.Body.String())
}
