package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nevill/jiangjing/api"
)

type Search func(engine, query string, o ...func(*SearchRequest)) (*api.Response, error)

func (Search) WithContext(ctx context.Context) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Context = ctx
	}
}

func newSearchFunc(tp api.Transport) Search {
	return func(engine, query string, o ...func(*SearchRequest)) (*api.Response, error) {
		r := SearchRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: engine,
			Query:  query,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SearchRequest struct {
	api.Request
	Engine string
	Query  string
}

func (r SearchRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/search", r.Engine)

	body, err := json.Marshal(map[string]interface{}{
		"query": r.Query,
	})
	if err != nil {
		return nil, err
	}

	req, err := api.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if r.Context != nil {
		req = req.WithContext(r.Context)
	}

	if len(body) > 0 {
		req.Header[api.HeaderContentType] = api.HeaderContentTypeJSON
	}

	res, err := r.Transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := api.Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}
