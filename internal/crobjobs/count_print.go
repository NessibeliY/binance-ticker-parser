package crobjobs

import (
	"fmt"

	"github.com/NessibeliY/binance-ticker-parser/internal/worker"
	"github.com/go-co-op/gocron"
)

func CountPrintJob(scheduler *gocron.Scheduler, workers []*worker.Worker) {
	//nolint:errcheck
	scheduler.Every(5).Seconds().Do(func() {
		totalRequests := 0
		for _, worker := range workers {
			totalRequests += worker.GetRequestsCount()
		}
		fmt.Printf("workers requests total: %d\n", totalRequests)
	})
}
