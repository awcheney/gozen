/*
  The zenclient package provides a Zendesk API client
*/
package zenclient

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
)

var defaultHeaders = map[string]string{
	"User-Agent":   "cheneyaw/gozen",
	"Content-Type": "application/json",
}

var subdomainRegexp = regexp.MustCompile("^[a-z][a-z0-9-]+[a-z0-9]$")

type ZenClient struct {
	RootURL    *url.URL
	UserAgent  string
	httpClient *http.client
}

func (c *ZenClient) ListUsers() ([]User, error) {
	// ListUsers : lists users
	rel := &url.URL{Path: "/users"}
	u := c.RootURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err
}
