package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nevill/jiangjing/api"
)

type Synonyms struct {
	List   SynonymsList
	Get    SynonymsGet
	Create SynonymsCreate
	Update SynonymsUpdate
	Delete SynonymsDelete
}

// SynonymsList retrieves all available synonym sets for the engine.
// see https://www.elastic.co/guide/en/app-search/current/synonyms.html#synonyms-get for details.
type SynonymsList func(name string, o ...func(*SynonymsListRequest)) (*api.Response, error)

func (h SynonymsList) WithContext(ctx context.Context) func(*SynonymsListRequest) {
	return func(r *SynonymsListRequest) {
		r.Context = ctx
	}
}

func (h SynonymsList) WithBody(body io.Reader) func(*SynonymsListRequest) {
	return func(r *SynonymsListRequest) {
		r.Body = body
	}
}

func newSynonymsListFunc(tp api.Transport) SynonymsList {
	return func(name string, o ...func(*SynonymsListRequest)) (*api.Response, error) {
		r := SynonymsListRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: name,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SynonymsListRequest struct {
	api.Request
	Engine string
	Body   io.Reader
}

func (r SynonymsListRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/synonyms", r.Engine)

	req, err := api.NewRequest(http.MethodGet, path, r.Body)
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

// SynonymsGet Retrieves a synonym set by ID.
// see https://www.elastic.co/guide/en/app-search/current/synonyms.html#synonyms-list-one for details.
type SynonymsGet func(name, id string, o ...func(*SynonymsGetRequest)) (*api.Response, error)

func (h SynonymsGet) WithContext(ctx context.Context) func(*SynonymsGetRequest) {
	return func(r *SynonymsGetRequest) {
		r.Context = ctx
	}
}

func newSynonymsGetFunc(tp api.Transport) SynonymsGet {
	return func(name, id string, o ...func(*SynonymsGetRequest)) (*api.Response, error) {
		r := SynonymsGetRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: name,
			Id:     id,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SynonymsGetRequest struct {
	api.Request
	Engine string
	Id     string
}

func (r SynonymsGetRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", r.Engine, r.Id)
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

// SynonymsCreate creates a new synonym set for the engine.
// see https://www.elastic.co/guide/en/app-search/current/synonyms.html#synonyms-create for details.
type SynonymsCreate func(name string, synonyms []string, o ...func(*SynonymsCreateRequest)) (*api.Response, error)

func (h SynonymsCreate) WithContext(ctx context.Context) func(*SynonymsCreateRequest) {
	return func(r *SynonymsCreateRequest) {
		r.Context = ctx
	}
}

func newSynonymsCreateFunc(tp api.Transport) SynonymsCreate {
	return func(name string, synonyms []string, o ...func(*SynonymsCreateRequest)) (*api.Response, error) {
		r := SynonymsCreateRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine:   name,
			Synonyms: synonyms,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SynonymsCreateRequest struct {
	api.Request
	Engine   string
	Synonyms []string
}

func (r SynonymsCreateRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/synonyms", r.Engine)

	body, err := json.Marshal(map[string]interface{}{
		"synonyms": r.Synonyms,
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

// SynonymsUpdate updates a synonym set by ID.
// see https://www.elastic.co/guide/en/app-search/current/synonyms.html#synonyms-update for details.
type SynonymsUpdate func(name, id string, synonyms []string, o ...func(*SynonymsUpdateRequest)) (*api.Response, error)

func (h SynonymsUpdate) WithContext(ctx context.Context) func(*SynonymsUpdateRequest) {
	return func(r *SynonymsUpdateRequest) {
		r.Context = ctx
	}
}

func newSynonymsUpdateFunc(tp api.Transport) SynonymsUpdate {
	return func(name, id string, synonyms []string, o ...func(*SynonymsUpdateRequest)) (*api.Response, error) {
		r := SynonymsUpdateRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine:   name,
			Id:       id,
			Synonyms: synonyms,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SynonymsUpdateRequest struct {
	api.Request
	Engine   string
	Id       string
	Synonyms []string
}

func (r SynonymsUpdateRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", r.Engine, r.Id)

	body, err := json.Marshal(map[string]interface{}{
		"id":       r.Id,
		"synonyms": r.Synonyms,
	})
	if err != nil {
		return nil, err
	}

	req, err := api.NewRequest(http.MethodPut, path, bytes.NewReader(body))
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

// SynonymsUpdate deletes a synonym set by ID.
// see https://www.elastic.co/guide/en/app-search/current/synonyms.html#synonyms-delete for details.
type SynonymsDelete func(name, id string, o ...func(*SynonymsDeleteRequest)) (*api.Response, error)

func (h SynonymsDelete) WithContext(ctx context.Context) func(*SynonymsDeleteRequest) {
	return func(r *SynonymsDeleteRequest) {
		r.Context = ctx
	}
}

func newSynonymsDeleteFunc(tp api.Transport) SynonymsDelete {
	return func(name, id string, o ...func(*SynonymsDeleteRequest)) (*api.Response, error) {
		r := SynonymsDeleteRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: name,
			Id:     id,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SynonymsDeleteRequest struct {
	api.Request
	Engine string
	Id     string
}

func (r SynonymsDeleteRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", r.Engine, r.Id)
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
