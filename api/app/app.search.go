package app

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/nevill/jiangjing/api"
)

type Search func(o ...func(*SearchRequest)) (*api.Response, error)

func (Search) WithContext(ctx context.Context) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Context = ctx
	}
}

func (Search) WithEngine(engine string) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Engine = engine
	}
}

func (Search) WithBody(body io.Reader) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Body = body
	}
}

func newSearchFunc(tp api.Transport) Search {
	return func(o ...func(*SearchRequest)) (*api.Response, error) {
		r := SearchRequest{
			Request: api.Request{
				Transport: tp,
			},
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
	Body   io.Reader
}

func (r SearchRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/search", r.Engine)
	req, err := api.NewRequest(http.MethodPost, path, r.Body)

	if err != nil {
		return nil, err
	}

	if r.Context != nil {
		req = req.WithContext(r.Context)
	}

	if r.Body != nil {
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
