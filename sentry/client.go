package sentry

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	CreateRelease(*Release) error
	CreateDeploy(*Deploy) error
}

type client struct {
	org     string
	project string
	api_key string
}

func NewClient(api_key string, org string, project string) Client {
	return &client{org, project, api_key}
}

func (c *client) CreateRelease(msg *Release) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	resp, err := http.Post(
		"https://sentry.io/api/0/projects/" + c.org + "/" + c.project + "/releases",
		"application/json",
		buf)
	if err != nil {
		return err
	}

	// 201 = Created, 208 = Already Exists
	if resp.StatusCode != 201 && resp.StatusCode != 208 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &Error{resp.StatusCode, string(t)}
	}

	return nil
}


func (c *client) CreateDeploy(msg *Deploy) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	resp, err := http.Post(
		"https://sentry.io/api/0/organizations/" + c.org +
			"/releases/" + msg.Name + "/deploys",
		"application/json",
		buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &Error{resp.StatusCode, string(t)}
	}

	return nil
}
