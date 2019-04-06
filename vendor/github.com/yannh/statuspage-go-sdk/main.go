package statuspagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiRoot = "https://api.statuspage.io/v1"

// HTTPClient is the http wrapper for the application
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	token      string
	httpClient HTTPClient
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

// Allows overriding the HTTP Client, leaving the choice
// of using a retry library to the user
func (client *Client) UseHTTPClient(httpClient HTTPClient) {
	client.httpClient = httpClient
}

func (client *Client) doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error) {
	componentURL := apiRoot + endpoint

	var body io.Reader

	if item != nil {
		data, err := json.Marshal(item)
		log.Printf("Statuspage Data %s\n", data)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method, componentURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "OAuth "+client.token)

	maxRetries := 10
	retryInterval := 5 * time.Second

	// Basic Retry logic around rate limiting
	resp, err = client.httpClient.Do(req)
	for retries := 1; resp.Status == "429" && retries <= maxRetries; retries = retries + 1 {
		time.Sleep(retryInterval)
		resp, err = client.httpClient.Do(req)
	}

	return resp, err
}

func createResource(client *Client, pageID, resourceType string, resource, result interface{}) error {
	resp, err := client.doHTTPRequest(
		"POST",
		"/pages/"+pageID+"/"+resourceType+"s",
		resource,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusCreated {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("Failed creating %s, request returned %d", resourceType, resp.StatusCode)
}

func readResource(client *Client, pageID, ID, resourceType string, target interface{}) error {
	resp, err := client.doHTTPRequest(
		"GET",
		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
		nil,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, target)
	}
	return fmt.Errorf("[ERROR] Statuspage could not find %s with ID: %s, http status %d", resourceType, ID, resp.StatusCode)
}

func updateResource(client *Client, pageID, resourceType, ID string, resource interface{}, result interface{}) error {
	resp, err := client.doHTTPRequest(
		"PATCH",
		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
		resource,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("Failed updating %s, request returned %d", resourceType, resp.StatusCode)
}

func deleteResource(client *Client, pageID, resourceType, ID string) error {
	resp, err := client.doHTTPRequest(
		"DELETE",
		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
		nil,
	)
	if err != nil {
		return err
	}

	// StatusGroup deletion returns StatusOK instead of StatsNoContent like other resources
	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("Failed deleting %s, request returned %d", resourceType, resp.StatusCode)
}
