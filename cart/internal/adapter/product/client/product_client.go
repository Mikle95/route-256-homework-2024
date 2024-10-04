package product

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type ProductClient struct {
	hostName string
	token    string
	client   http.Client
	limiter  *rate.Limiter
}

func NewProductClient(c http.Client, host string, token string, rps time.Duration) *ProductClient {
	rt := rate.Every(time.Second / rps)
	lim := rate.NewLimiter(rt, 1)

	return &ProductClient{client: c, hostName: host, token: token, limiter: lim}
}
