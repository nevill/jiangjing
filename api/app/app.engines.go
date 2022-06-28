package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nevill/jiangjing/api"
)

type Engines struct {
	List   EnginesList
	Get    EnginesGet
	Create EnginesCreate
	Delete EnginesDelete
}

// EnginesList Retrieve all engines.
// see https://www.elastic.co/guide/en/app-search/current/engines.html#engines-list for details.
type EnginesList func(o ...func(*EnginesListRequest)) (*api.Response, error)

func (h EnginesList) WithContext(ctx context.Context) func(*EnginesListRequest) {
	return func(r *EnginesListRequest) {
		r.Context = ctx
	}
}

func newEnginesListFunc(tp api.Transport) EnginesList {
	return func(o ...func(*EnginesListRequest)) (*api.Response, error) {
		r := EnginesListRequest{
			api.Request{
				Transport: tp,
			},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EnginesListRequest struct {
	api.Request
}

func (r EnginesListRequest) Do() (*api.Response, error) {
	path := "/api/as/v1/engines"

	req, err := api.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	if r.Context != nil {
		req = req.WithContext(r.Context)
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

// EnginesGet Retrieves details of a given engine by its name.
// see https://www.elastic.co/guide/en/app-search/current/engines.html#engines-get for details.
type EnginesGet func(name string, o ...func(*EnginesGetRequest)) (*api.Response, error)

func (h EnginesGet) WithContext(ctx context.Context) func(*EnginesGetRequest) {
	return func(r *EnginesGetRequest) {
		r.Context = ctx
	}
}

func newEnginesGetFunc(tp api.Transport) EnginesGet {
	return func(name string, o ...func(*EnginesGetRequest)) (*api.Response, error) {
		r := EnginesGetRequest{
			Request: api.Request{
				Transport: tp,
			},
			Name: name,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EnginesGetRequest struct {
	api.Request
	Name string
}

func (r EnginesGetRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s", r.Name)

	req, err := api.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	if r.Context != nil {
		req = req.WithContext(r.Context)
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

// EnginesCreate Creates an App Search Engine.
// see https://www.elastic.co/guide/en/app-search/current/engines.html#engines-create for details.
type EnginesCreate func(name string, o ...func(*EnginesCreateRequest)) (*api.Response, error)

func (h EnginesCreate) WithContext(ctx context.Context) func(*EnginesCreateRequest) {
	return func(r *EnginesCreateRequest) {
		r.Context = ctx
	}
}

func newEngineCreateFunc(tp api.Transport) EnginesCreate {
	return func(name string, o ...func(*EnginesCreateRequest)) (*api.Response, error) {
		r := EnginesCreateRequest{
			Request: api.Request{
				Transport: tp,
			},
			Name: name,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EnginesCreateRequest struct {
	api.Request
	Name string
}

func (r EnginesCreateRequest) Do() (*api.Response, error) {
	path := "/api/as/v1/engines"

	body, err := json.Marshal(map[string]string{
		"name": r.Name,
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

// EnginesCreate Deletes an engine by name.
// see https://www.elastic.co/guide/en/app-search/current/engines.html#engines-delete for details.
type EnginesDelete func(name string, o ...func(*EnginesDeleteRequest)) (*api.Response, error)

func (h EnginesDelete) WithContext(ctx context.Context) func(*EnginesDeleteRequest) {
	return func(r *EnginesDeleteRequest) {
		r.Context = ctx
	}
}

func newEngineDeleteFunc(tp api.Transport) EnginesDelete {
	return func(name string, o ...func(*EnginesDeleteRequest)) (*api.Response, error) {
		r := EnginesDeleteRequest{
			Request: api.Request{
				Transport: tp,
			},
			Name: name,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EnginesDeleteRequest struct {
	api.Request
	Name string
}

func (r EnginesDeleteRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s", r.Name)

	req, err := api.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	if r.Context != nil {
		req = req.WithContext(r.Context)
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
