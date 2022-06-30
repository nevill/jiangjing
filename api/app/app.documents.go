package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nevill/jiangjing/api"
)

type Documents struct {
	Create DocumentsCreate
	Delete DocumentsDelete
	List   DocumentsList
}

type DocumentsCreate func(string, ...func(*DocumentsCreateRequest)) (*api.Response, error)

func (DocumentsCreate) WithContext(ctx context.Context) func(*DocumentsCreateRequest) {
	return func(r *DocumentsCreateRequest) {
		r.Context = ctx
	}
}

func (DocumentsCreate) WithDocuments(docs ...map[string]interface{}) func(*DocumentsCreateRequest) {
	return func(r *DocumentsCreateRequest) {
		r.Documents = append(r.Documents, docs...)
	}
}

func newDocumentsCreateFunc(tp api.Transport) DocumentsCreate {
	return func(engine string, o ...func(*DocumentsCreateRequest)) (*api.Response, error) {
		r := DocumentsCreateRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: engine,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type DocumentsCreateRequest struct {
	api.Request
	Engine    string
	Documents []map[string]interface{}
}

func (r DocumentsCreateRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/documents", r.Engine)

	body, err := json.Marshal(r.Documents)
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

type DocumentsDelete func(string, ...func(*DocumentsDeleteRequest)) (*api.Response, error)

func (DocumentsDelete) WithContext(ctx context.Context) func(*DocumentsDeleteRequest) {
	return func(r *DocumentsDeleteRequest) {
		r.Context = ctx
	}
}

func (DocumentsDelete) WithIds(ids ...string) func(*DocumentsDeleteRequest) {
	return func(r *DocumentsDeleteRequest) {
		r.Ids = append(r.Ids, ids...)
	}
}

func newDocumentsDeleteFunc(tp api.Transport) DocumentsDelete {
	return func(engine string, o ...func(*DocumentsDeleteRequest)) (*api.Response, error) {
		r := DocumentsDeleteRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: engine,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type DocumentsDeleteRequest struct {
	api.Request
	Engine string
	Ids    []string
}

func (r DocumentsDeleteRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/documents", r.Engine)

	body, err := json.Marshal(r.Ids)
	if err != nil {
		return nil, err
	}

	req, err := api.NewRequest(http.MethodDelete, path, bytes.NewReader(body))
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

type DocumentsList func(string, ...func(*DocumentsListRequest)) (*api.Response, error)

func (DocumentsList) WithContext(ctx context.Context) func(*DocumentsListRequest) {
	return func(r *DocumentsListRequest) {
		r.Context = ctx
	}
}

func newDocumentsListFunc(tp api.Transport) DocumentsList {
	return func(engine string, o ...func(*DocumentsListRequest)) (*api.Response, error) {
		r := DocumentsListRequest{
			Request: api.Request{
				Transport: tp,
			},
			Engine: engine,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type DocumentsListRequest struct {
	api.Request
	Engine string
}

func (r DocumentsListRequest) Do() (*api.Response, error) {
	path := fmt.Sprintf("/api/as/v1/engines/%s/documents/list", r.Engine)
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
