package pingrobot

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	interval = time.Second * 3
	timeout = time.Second
	Workers = 7
)

var urls = []string{
	"https://apple.com",
	"https://ya.ru",
	"https://golang.org",
	"https://bitrix24.ru",
	"https://mail.ru",
	"https://yelp.com",
	"https://yahoo.ru",
}

func GenerateJobs(pool *Pool) {
	for {
		for _, job := range urls {
			pool.Push(Job{URL:job})
		}
		time.Sleep(interval)
	}
}

func ShowInfo(results <-chan Result) {
	func() {
		for res := range results {
			fmt.Println(res.Info())
		}
	}()
}

func GracefulShutdown() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit
}