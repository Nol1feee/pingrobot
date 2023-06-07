package pingrobot

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func (w worker) process(j Job) Result {
	t := time.Now()

	result := Result{Url: j}
	resp, err := w.client.Get(j.URL)

	if err != nil {
		result.error = err
		return result
	}

	result.statusCode = resp.StatusCode
	result.responseTime = time.Since(t)

	return result
}

func newWorker(timeout time.Duration) *worker {
	return &worker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}