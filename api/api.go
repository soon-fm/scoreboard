package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// A user from the API
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar_url"`
	FamilyName  string `json:"family_name"`
	GivenName   string `json:"given_name"`
}

// Response data decoder interface
type decoder interface {
	Decode(v interface{}) error
}

type Client struct {
	// Exported fields
	Config Configurer
	// Unexported Fields
	client  *http.Client // Optional custom http client
	decoder decoder      // Optional response decoder
}

// Builds the request url
func (c *Client) url(path string, values url.Values) string {
	u := url.URL{}
	u.Scheme = c.Config.Scheme()
	u.Host = c.Config.Host()
	u.Path = path
	u.RawQuery = values.Encode()
	return u.String()
}

// Makes a request to the API, returning the response status code,
// resonse data and any errors that occured
func (c *Client) request(method, path string, values url.Values, body io.Reader) (int, []byte, error) {
	client := c.client
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequest(method, c.url(path, values), body)
	if err != nil {
		return 0, nil, err
	}
	rsp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return rsp.StatusCode, nil, err
	}
	return rsp.StatusCode, data, nil
}

// Decodes data into a given interface
func (c *Client) decode(data []byte, v interface{}) error {
	decoder := c.decoder
	if decoder == nil {
		decoder = json.NewDecoder(bytes.NewReader(data))
	}
	return decoder.Decode(v)
}

// Get's a specific user from the API
func (c *Client) User(id string) (*User, error) {
	path := fmt.Sprintf("/users/%s", id)
	_, data, err := c.request(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}
	user := &User{}
	if err := c.decode(data, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Constructs a new API Client
func New(c Configurer) *Client {
	return &Client{Config: c}
}
