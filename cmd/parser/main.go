package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/NessibeliY/binance-ticker-parser/internal/config"
	"github.com/NessibeliY/binance-ticker-parser/internal/crobjobs"
	"github.com/NessibeliY/binance-ticker-parser/internal/worker"
	"github.com/NessibeliY/binance-ticker-parser/pkg"
	"github.com/go-co-op/gocron"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	numCPU := runtime.NumCPU()
	if cfg.MaxWorkers > numCPU {
		cfg.MaxWorkers = numCPU
	}

	symbolGroups := pkg.DivideSlice(cfg.Symbols, cfg.MaxWorkers)
	workers := make([]*worker.Worker, cfg.MaxWorkers)
	var wg sync.WaitGroup

	for i, symbols := range symbolGroups {
		workers[i] = worker.NewWorker(symbols)
		wg.Add(1)
		go workers[i].Run(&wg)
	}

	go func() {
		scheduler := gocron.NewScheduler(time.UTC)
		crobjobs.CountPrintJob(scheduler, workers)
		scheduler.StartBlocking()
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(input) == "STOP" {
			break
		}
	}

	fmt.Println("Stopping workers...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// channel to signal when all workers are done
	done := make(chan struct{})

	go func() {
		for _, worker := range workers {
			go worker.StopWorker(ctx)
		}
		wg.Wait()
		close(done)
	}()

	<-done
	fmt.Println("All workers stopped.")
}
