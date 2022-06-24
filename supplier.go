package cyberrisk

import (
	"net/http"
)

type SupplierService service

func (srv SupplierService) List() ([]Supplier, error) {
	resp := []Supplier{}
	err := srv.client.Do(http.MethodGet, "/api/v1/suppliers", nil, &resp)
	return resp, err
}

func (srv SupplierService) Create(create []RequestSupplier) ([]Supplier, error) {
	resp := []Supplier{}
	err := srv.client.Do(http.MethodPost, "/api/v1/suppliers", create, &resp)
	return resp, err
}

func (srv SupplierService) Delete(id string) error { // supplierID or externalID
	err := srv.client.Do(http.MethodDelete, "/api/v1/suppliers/"+id, nil, nil)
	return err
}

func (srv SupplierService) GetRating(id string) (Rating, error) { // supplierID or externalID
	resp := Rating{}
	err := srv.client.Do(http.MethodGet, "/api/v1/suppliers/"+id+"/rating", nil, &resp)
	return resp, err
}
