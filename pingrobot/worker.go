package pingrobot

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func (w worker) process(j Job) Result {
	result := Result{Url: j}
	
	t := time.Now()
	resp, err := w.client.Get(j.URL)

	if err != nil {
		result.error = err
		return result
	}

	result.statusCode = resp.StatusCode
	//если количество потоков >= количество urls, то responseTime некорректен
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