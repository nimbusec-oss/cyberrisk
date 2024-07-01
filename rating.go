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
		Path:     "/api/v2/ratings",
		RawQuery: v.Encode(),
	}

	resp := []Rating{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &resp)
	return resp, err
}

func (srv RatingService) CreateCRR(create []RequestRatingCRR) (ContingentUsage, error) {
	resp := ContingentUsage{}
	err := srv.client.Do(http.MethodPost, "/api/v2/ratings/crr", create, &resp)
	return resp, err
}

func (srv RatingService) CreateDPR(create []RequestRatingDPR) (ContingentUsage, error) {
	resp := ContingentUsage{}
	err := srv.client.Do(http.MethodPost, "/api/v2/ratings/dpr", create, &resp)
	return resp, err
}
