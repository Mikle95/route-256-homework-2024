package middleware

import (
	"log"
	"net/http"
)

type RetryRT struct {
	R http.RoundTripper
}

func NewRetryRT(r http.RoundTripper) *RetryRT {
	return &RetryRT{R: r}
}

// TODO: Протестировать
func (r *RetryRT) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	for i := 0; i < 3; i++ {
		resp, err = r.R.RoundTrip(req)
		if err != nil || resp.StatusCode != 404 && resp.StatusCode != http.StatusTooManyRequests {
			break
		}
		log.Printf("%v got status %v, Retrying", req.URL, resp.Status)
	}

	return resp, err
}

// func resetBody(request *http.Request, originalBody []byte) {
// 	request.Body = io.NopCloser(bytes.NewBuffer(originalBody))
// 	request.GetBody = func() (io.ReadCloser, error) {
// 		return io.NopCloser(bytes.NewBuffer(originalBody)), nil
// 	}
// }
