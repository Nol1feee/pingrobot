This is an extremely minimalist analog of pingrobot.io or UptimeRobot.com, which is used for real-time server monitoring.

An example of usage can be found below, where you simply need to specify your websites in the common.go folder and the frequency (interval) at which you want to ping the specified websites.

>##### example
```go
package main

import (
    "pingRobot/pingrobot"
)
c main() {

    results := make(chan pingrobot.Result)
    pool := pingrobot.NewPool(pingrobot.Workers, results)

 	go pingrobot.GenerateJobs(pool)
 	pool.InitWorkers()
 	go pingrobot.ShowInfo(results)

 	pingrobot.GracefulShutdown()

 	pool.Stop()
 }
 ```