package pingrobot

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	//Every "interval" the service will send GET requests	
	interval = time.Second * 10
	//timeout for hhtp.Client
	timeout = time.Second * 3
	//goroutines
	Workers = 10
)

//websites that you wanna check
var urls = []string{
	"https://apple.com",
	"https://ya.ru",
	"https://golang.org",
	"https://bitrix24.ru",
	"https://mail.ru",
	"https://yelp.com",
	"https://yahoo.ru",
}

func GenerateJobs(pool *pool) {
	for {
		for _, job := range urls {
			pool.push(Job{URL:job})
		}
		time.Sleep(interval)
	}
}

func ShowInfo(results <-chan Result) {
	func() {
		for res := range results {
			fmt.Println(res.info())
		}
	}()
}

func GracefulShutdown() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit
}