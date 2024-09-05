package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type RetryRT struct {
	R http.RoundTripper
}

func NewRetryRT(r http.RoundTripper) *RetryRT {
	return &RetryRT{R: r}
}

func (r *RetryRT) RoundTrip(req *http.Request) (resp *http.Response, err error) {

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := req.Body.Close(); err != nil {
		return nil, err
	}

	// 1 запрос и 3 ретрая
	for i := 0; i < 4; i++ {
		req.Body = io.NopCloser(bytes.NewReader(data))
		resp, err = r.R.RoundTrip(req)
		if err != nil || resp.StatusCode != 420 && resp.StatusCode != http.StatusTooManyRequests {
			break
		}
		log.Printf("%v got status %v, Retrying", req.URL, resp.Status)
	}

	return resp, err
}
