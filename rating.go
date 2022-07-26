package cyberrisk

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type RatingService service

func (srv RatingService) List(filter *RatingFilter) ([]Rating, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v1/ratings",
		RawQuery: v.Encode(),
	}

	resp := []Rating{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &resp)
	return resp, err
}

func (srv RatingService) Create(create []RequestRating) ([]Rating, error) {
	resp := []Rating{}
	err := srv.client.Do(http.MethodPost, "/api/v1/ratings", create, &resp)
	return resp, err
}
