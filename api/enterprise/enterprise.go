package enterprise

import (
	"context"
	"net/http"

	"github.com/nevill/jiangjiang/api"
)

func newHealthFunc(tp api.Transport) Health {
	return func(o ...func(*HealthRequest)) (*api.Response, error) {
		r := HealthRequest{
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

type Health func(o ...func(*HealthRequest)) (*api.Response, error)

func (h Health) WithContext(ctx context.Context) func(*HealthRequest) {
	return func(r *HealthRequest) {
		r.Context = ctx
	}
}

type HealthRequest struct {
	api.Request
}

func (r HealthRequest) Do() (*api.Response, error) {
	path := "/api/ent/v1/internal/health"

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

type API struct {
	Health Health
}

func New(t api.Transport) *API {
	return &API{
		Health: newHealthFunc(t),
	}
}
