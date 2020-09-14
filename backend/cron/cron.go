package cron

import (
    "sync"
    "time"
    "github.com/yellia1989/tex-go/tools/log"
)

var wg sync.WaitGroup
var stop chan bool

func init() {
    stop = make(chan bool)
}

type croner interface {
    run(now time.Time)
}

func startCron(name string, f croner, d time.Duration) {
    wg.Add(1)
    go func(stop <- chan bool, name string) {
        defer wg.Done()

        ticker := time.NewTicker(d)
        defer ticker.Stop()
        for {
            select {
            case <- stop:
                log.Debug("%s cron stop", name)
                return
            case t := <- ticker.C: 
                f.run(t)
            }
        }
    }(stop, name)
}

func Start() {

    // 免费福利
    startCron("welfare", newWelfare(), time.Second * 5)

    log.Debug("all cron start")
}

func Stop() {
    close(stop)
    wg.Wait()

    log.Debug("all cron stop")
}
