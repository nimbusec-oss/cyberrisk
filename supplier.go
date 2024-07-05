package cyberrisk

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type SupplierService service

// func (srv SupplierService) List(filter *SupplierFilter) ([]Supplier, error) {
// 	v, err := query.Values(filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := url.URL{
// 		Path:     "/api/v2/suppliers",
// 		RawQuery: v.Encode(),
// 	}

// 	resp := []Supplier{}
// 	err = srv.client.Do(http.MethodGet, u.String(), nil, &resp)
// 	return resp, err
// }

func (srv SupplierService) Create(create []RequestSupplier) (PostSuppliersResponse, error) {
	resp := PostSuppliersResponse{}
	err := srv.client.Do(http.MethodPost, "/api/v2/suppliers", create, &resp)
	return resp, err
}

func (srv SupplierService) Unassign(unassign UnassignSupplier) error {
	v, err := query.Values(unassign)
	if err != nil {
		return err
	}
	u := url.URL{
		Path:     "/api/v2/suppliers/unassign",
		RawQuery: v.Encode(),
	}

	return srv.client.Do(http.MethodDelete, u.String(), nil, nil)
}
