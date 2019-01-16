package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	restful "github.com/emicklei/go-restful"
)

type restfulTesting struct {
	container *restful.Container
	ws        *restful.WebService
}

func (rest *restfulTesting) SetupTest() {
	rest.container = restful.NewContainer()
	rest.container.Add(rest.ws)
}

func (rest *restfulTesting) request(method, url string, body interface{}) *httptest.ResponseRecorder {
	var b io.Reader
	if body != nil {
		bodyByte, _ := json.Marshal(body)
		b = bytes.NewReader(bodyByte)
	}
	req, _ := http.NewRequest(method, url, b)
	req.Header.Set("Content-Type", restful.MIME_JSON)
	resp := httptest.NewRecorder()
	rest.container.ServeHTTP(resp, req)
	return resp
}
