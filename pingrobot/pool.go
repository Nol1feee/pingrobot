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

type pool struct {
	worker       *worker
	Jobs         chan Job
	workersCount int
	Res          chan Result
	Wg           *sync.WaitGroup
}

func NewPool(workersCount int, result chan Result) *pool {
	return &pool{
		worker:       newWorker(timeout),
		Jobs:         make(chan Job),
		workersCount: workersCount,
		Res:          result,
		Wg:           new(sync.WaitGroup),
	}
}


func (p *pool) doHtpp() {
	for {
		p.Res <- p.worker.process(<-p.Jobs)
		p.Wg.Done()
	}
}

func (p *pool) InitWorkers() {
	for i := 0; i < p.workersCount; i++ {
		go p.doHtpp()
	}
}

func (p *pool) push(j Job) {
	p.Jobs <- j
	p.Wg.Add(1)
}

func (p *pool) Stop() {
	p.Wg.Wait()
	close(p.Jobs)
}

func (r Result) info() string {
	if r.error != nil {
		//можно чекать на 5xx ответы + отправлять письмо на почту
		return fmt.Sprintf("\n[ERROR]\nURL - %s, %s", r.Url, r.error)
	}
	return fmt.Sprintf("[SUCCESS]; URL - %s, statusCode - %d, responseTime - %s", r.Url, r.statusCode, r.responseTime)
}