package jiangjing

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/nevill/jiangjiang/api/enterprise"
)

type Client struct {
	EnterpriseSearch
	AppSearch AppSearch
}

type Config struct {
	Address  string
	Username string
	Password string
	Token    string
}

type EnterpriseSearch struct {
	*enterprise.API
}

type AppSearch struct {
}

func addrToUrls(address string) ([]*url.URL, error) {
	u, err := url.Parse(strings.TrimRight(address, "/"))
	if err != nil {
		return nil, fmt.Errorf("cannot parse url: %v", err)
	}

	return []*url.URL{u}, nil
}

func NewClient(cfg Config) (*Client, error) {
	urls, err := addrToUrls(cfg.Address)

	if err != nil {
		return nil, err
	}

	if len(cfg.Token) > 0 && len(cfg.Username) > 0 {
		return nil, errors.New("cannot create client: both APIKey and Username are set")
	}

	tp, err := elastictransport.New(elastictransport.Config{
		URLs:         urls,
		Username:     cfg.Username,
		Password:     cfg.Password,
		ServiceToken: cfg.Token,
	})

	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}

	c := &Client{
		EnterpriseSearch: EnterpriseSearch{
			enterprise.New(tp),
		},
	}

	return c, nil
}
