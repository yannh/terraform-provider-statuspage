package main

import (
	"log"
	"os"

	gostatuspage "github.com/yannh/terraform-provider-statuspage/go-statuspage"
)

type Config struct {
	AuthToken string
}

// Client returns a new client for accessing Packet's API.
func (c *Config) Client() *gostatuspage.Client {
	token := ""

	if v := os.Getenv("STATUSPAGE_TOKEN"); v != "" {
		token = v
	}

	log.Printf("[INFO] Statuspage Client configured!")
	return gostatuspage.NewClient(token)
}
