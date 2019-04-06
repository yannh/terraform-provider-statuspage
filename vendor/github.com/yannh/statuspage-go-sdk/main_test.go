package statuspagesdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	req     *http.Request
	errCode int
}

func NewMockHTTPClient(errCode int) *mockHTTPClient {
	return &mockHTTPClient{
		errCode: errCode,
	}
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return &http.Response{StatusCode: c.errCode}, nil
}

type mockClient struct {
	resp *http.Response
	err  error
}

func NewMockClient(resp *http.Response, err error) *mockClient {
	return &mockClient{
		resp: resp,
		err:  err,
	}
}

func (client *mockClient) doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error) {
	return client.resp, client.err
}

func TestDoHTTPRequest(t *testing.T) {
	token := "token"
	client := NewClient(token)
	mc := NewMockHTTPClient(200)
	client.httpClient = mc

	r1, _ := http.NewRequest(
		"POST",
		"https://api.statuspage.io/v1/endpoint",
		strings.NewReader("{\"Name\":\"resource_name\"}"),
	)
	r2, _ := http.NewRequest(
		"GET",
		"https://api.statuspage.io/v1/components",
		strings.NewReader("{\"Name\":\"resource_name\",\"Description\":\"description\"}"),
	)

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

func TestCreateResource(t *testing.T) {
	nameShouldBe := "resource name"
	r := &http.Response{StatusCode: http.StatusCreated, Body: ioutil.NopCloser(bytes.NewReader([]byte("{\"Name\":\"" + nameShouldBe + "\"}")))}
	client := NewMockClient(r, nil)

	target := struct{ Name string }{}
	err := createResource(client, "somePageID", "component", struct{ Name string }{nameShouldBe}, &target)

	if err != nil {
		t.Errorf("Error while creating resource: %+v", err)
	}

	if target.Name != nameShouldBe {
		t.Errorf("Failed to read resources: %s, %s", nameShouldBe, target.Name)
	}
}
