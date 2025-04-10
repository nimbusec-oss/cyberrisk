package cyberrisk

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type TimelineService service

func (srv TimelineService) List(filter *TimelineFilter) ([]RatingTimeline, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v2/ratingtimelines",
		RawQuery: v.Encode(),
	}

	resp := []RatingTimeline{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &resp)
	return resp, err
}
