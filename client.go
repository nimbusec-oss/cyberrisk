package cyberrisk

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type service struct {
	client *Client
}

type Client struct {
	http *http.Client
	base string
	dry  bool

	common    service // Reuse a single struct instead of allocating one for each service on the heap.
	Suppliers *SupplierService
	Ratings   *RatingService
	Account   *AccountService
}

type Error struct {
	Message    string            `json:"message"`
	Validation map[string]string `json:"validation,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	Endpoint     Endpoint
	Dry          bool
}

type Endpoint struct {
	AuthURL   string
	ServerURL string
}

var DefaultEndpoint = Endpoint{
	AuthURL:   "https://auth.nimbusec.com/oauth2/token",
	ServerURL: "https://cr.nimbusec.com",
}

func NewClient(config ClientConfig) *Client {
	oauth := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.AuthURL,
	}

	c := &Client{http: oauth.Client(context.Background()), base: config.Endpoint.ServerURL, dry: config.Dry}
	c.common = service{client: c}
	c.Suppliers = (*SupplierService)(&c.common)
	c.Ratings = (*RatingService)(&c.common)
	c.Account = (*AccountService)(&c.common)

	return c
}

func (client *Client) Ping() error {
	err := client.Do(http.MethodHead, "/health", nil, nil)
	return err
}

func (client *Client) Do(method, path string, in, out interface{}) error {
	var req *http.Request
	var err error

	// encode `in` object as json body if provided
	if in != nil {
		body := &bytes.Buffer{}
		err = json.NewEncoder(body).Encode(in)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, client.base+path, body)
	} else {
		req, err = http.NewRequest(method, client.base+path, nil)
	}

	if err != nil {
		return err
	}

	// dry mode active, do not make an action except for read-only GET requests
	if client.dry && method != http.MethodGet {
		return nil
	}

	// perform request
	req = req.WithContext(context.Background())
	resp, err := client.http.Do(req)
	if err != nil {
		return err
	}

	// decode response body into `out` object if provided
	defer resp.Body.Close()

	// check for api errors, try to decode in error object
	if resp.StatusCode >= 300 {
		msg := Error{}
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			return err
		}

		return msg
	}

	if out != nil {
		err = json.NewDecoder(resp.Body).Decode(&out)
		if err != nil {
			return err
		}
	}

	return nil
}
