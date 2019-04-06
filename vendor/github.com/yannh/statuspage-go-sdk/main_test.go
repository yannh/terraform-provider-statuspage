package statuspagesdk

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	req *http.Request
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return nil, nil
}

func TestDoHTTPRequest(t *testing.T) {
	token := "token"
	client := NewClient(token)
	mc := &mockHTTPClient{}
	client.httpClient = mc

	r1, _ := http.NewRequest("POST", "https://api.statuspage.io/v1/endpoint", strings.NewReader("{\"Name\":\"resource_name\"}"))
	r2, _ := http.NewRequest("GET", "https://api.statuspage.io/v1/components", strings.NewReader("{\"Name\":\"resource_name\",\"Description\":\"description\"}"))

	testCases := []struct {
		method    string
		endpoint  string
		resource  interface{}
		reqShould *http.Request
	}{
		{
			"POST",
			"/endpoint",
			struct{ Name string }{"resource_name"},
			r1,
		},
		{
			"GET",
			"/components",
			struct {
				Name        string
				Description string
			}{
				"resource_name",
				"description",
			},
			r2,
		},
	}

	for _, testCase := range testCases {
		client.doHTTPRequest(testCase.method, testCase.endpoint, testCase.resource)
		if mc.req.Method != testCase.reqShould.Method {
			t.Errorf("Request method should be %s, got %s", testCase.method, mc.req.Method)
		}
		if mc.req.URL.String() != testCase.reqShould.URL.String() {
			t.Errorf("Request URL should be %s, got %s", testCase.reqShould.URL.String(), mc.req.URL.String())
		}

		if mc.req.Header.Get("Authorization") != "OAuth "+token {
			t.Errorf("Request should have Authorization header set to OAuth %s, got %s", token, mc.req.Header.Get("Authorization"))
		}

		bodyWant := new(bytes.Buffer)
		bodyWant.ReadFrom(mc.req.Body)
		bodyWantS := bodyWant.String()

		bodyHave := new(bytes.Buffer)
		bodyHave.ReadFrom(testCase.reqShould.Body)
		bodyHaveS := bodyHave.String()

		if bodyWantS != bodyHaveS {
			t.Errorf("Request Body should be %s, got %s", bodyWantS, bodyHaveS)
		}
	}
}
