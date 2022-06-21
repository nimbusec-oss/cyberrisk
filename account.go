package cyberrisk

import (
	"net/http"
)

type AccountService service

func (srv AccountService) Get() (Account, error) {
	resp := Account{}
	err := srv.client.Do(http.MethodGet, "/api/v1/account", nil, &resp)
	return resp, err
}
