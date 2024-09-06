package product

import "net/http"

type ProductClient struct {
	hostName string
	token    string
	client   http.Client
}

func NewProductClient(c http.Client, host string, token string) *ProductClient {
	return &ProductClient{client: c, hostName: host, token: token}
}
