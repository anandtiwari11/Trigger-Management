package initializers

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-co-op/gocron/v2"
)

var Scheduler gocron.Scheduler
var once sync.Once

func InitScheduler() {
	once.Do(func() {
		Scheduler, _ = gocron.NewScheduler()
		Scheduler.Start()
		go func() {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan
			log.Println("\nInterrupt signal received. Exiting...")
			_ = Scheduler.Shutdown()
			os.Exit(0)
		}()
	})
}