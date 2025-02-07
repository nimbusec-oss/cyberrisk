package cyberrisk

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type CertificationService service

func (srv CertificationService) List(filter *CertificationFilter) ([]SupplierCertifications, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path: "/api/v2/certifications",
		RawQuery: v.Encode(),
	}
	
	resp := []SupplierCertifications{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &resp)
	return resp, err
}