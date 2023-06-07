package pingrobot

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	URL string
}

type Result struct {
	statusCode   int
	Url          Job
	responseTime time.Duration
	error        error
}

type Pool struct {
	worker       *worker
	Jobs         chan Job
	workersCount int
	Res          chan Result
	Wg           *sync.WaitGroup
}

func NewPool(workersCount int, result chan Result) *Pool {
	return &Pool{
		worker:       newWorker(timeout),
		Jobs:         make(chan Job),
		workersCount: workersCount,
		Res:          result,
		Wg:           new(sync.WaitGroup),
	}
}


func (p *Pool) doHtpp() {
	for {
		p.Res <- p.worker.process(<-p.Jobs)
		p.Wg.Done()
	}
}

func (p *Pool) InitWorkers() {
	for i := 0; i < p.workersCount; i++ {
		go p.doHtpp()
	}
}

func (p *Pool) Push(j Job) {
	p.Jobs <- j
	p.Wg.Add(1)
}

func (p *Pool) Stop() {
	p.Wg.Wait()
	close(p.Jobs)
}

func (r Result) Info() string {
	if r.error != nil {
		return fmt.Sprintf("\n[ERROR]\nURL - %s, %s", r.Url, r.error)
	}
	return fmt.Sprintf("[SUCCESS]; URL - %s, statusCode - %d, responseTime - %s", r.Url, r.statusCode, r.responseTime)
}