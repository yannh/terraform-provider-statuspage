package statuspage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiRoot = "https://api.statuspage.io/v1"

type Client struct {
	token string
}

type Component struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	GroupID            string `json:"group_id,omitempty"`
	Showcase           bool   `json:"showcase,omitempty"`
	Status             string `json:"status,omitempty"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded,omitempty"`
}

type componentFull struct {
	Component
	ID        string `json:"id"`
	Position  int32  `json:"position"`
	CreatedAt string `json:"created_at"`
}

func (client *Client) doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error) {
	httpClient := &http.Client{}
	componentURL := apiRoot + endpoint

	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", componentURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "OAuth "+client.token)

	log.Printf("Request %+v\n", req)
	log.Printf("Data %s\n", data)
	resp, err = httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (client *Client) CreateComponent(pageID string, component *Component) (id string, err error) {
	resp, err := client.doHTTPRequest(
		"POST",
		"/pages/"+pageID+"/components",
		struct{ Component *Component }{component},
	)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusCreated {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		c := componentFull{}
		err = json.Unmarshal(bodyBytes, &c)
		if err != nil {
			return "", err
		}
		return c.ID, nil
	}

	return "", fmt.Errorf("Failed creating component, request returned %d", resp.StatusCode)
}

func (client *Client) GetComponent(pageID string, componentID string) (*Component, error) {
	resp, err := client.doHTTPRequest(
		"GET",
		"/pages/"+pageID+"/component/"+componentID,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		c := &Component{}
		err = json.Unmarshal(bodyBytes, c)
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	return nil, nil
}

func (client *Client) DeleteComponent(pageID, componentID string) (err error) {
	resp, err := client.doHTTPRequest(
		"DELETE",
		"/pages/"+pageID+"/component/"+componentID,
		nil,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return fmt.Errorf("Failed deleting component, request returned %d", resp.StatusCode)
}

func (client *Client) UpdateComponent(pageID, componentID string, component *Component) (err error) {
	resp, err := client.doHTTPRequest(
		"POST",
		"/pages/"+pageID+"/components/"+componentID,
		struct{ Component *Component }{component},
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
		c := componentFull{}
		err = json.Unmarshal(bodyBytes, &c)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("Failed creating component, request returned %d", resp.StatusCode)
}

func NewClient(token string) *Client {
	return &Client{token}
}
