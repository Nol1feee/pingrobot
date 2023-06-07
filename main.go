package main

import (
	"pingRobot/pingrobot"
)

func main() {
	results := make(chan pingrobot.Result)
	pool := pingrobot.NewPool(pingrobot.Workers, results)

	go pingrobot.GenerateJobs(pool)
	pool.InitWorkers()
	go pingrobot.ShowInfo(results)

	pingrobot.GracefulShutdown()

	pool.Stop()
}