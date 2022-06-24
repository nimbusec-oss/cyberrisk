package cyberrisk

import "net/http"

type RatingService service

func (srv RatingService) List() ([]Rating, error) {
	resp := []Rating{}
	err := srv.client.Do(http.MethodGet, "/api/v1/ratings", nil, &resp)
	return resp, err
}

func (srv RatingService) Create(create []RequestRating) ([]Rating, error) {
	resp := []Rating{}
	err := srv.client.Do(http.MethodPost, "/api/v1/ratings", create, &resp)
	return resp, err
}
