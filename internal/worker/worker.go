package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/NessibeliY/binance-ticker-parser/internal/dto"
	"github.com/NessibeliY/binance-ticker-parser/internal/values"
)

type TickerUpdate struct {
	Symbol string
	Price  string
}

type Worker struct {
	symbols      []string
	RequestCount uint64
	StopChan     chan struct{}
	UpdateChan   chan TickerUpdate
}

func NewWorker(symbols []string, updateChan chan TickerUpdate) *Worker {
	return &Worker{
		symbols:    symbols,
		StopChan:   make(chan struct{}),
		UpdateChan: updateChan,
	}
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	tickerPrices := make(map[string]string)

	for {
		select {
		case <-w.StopChan:
			fmt.Println("Worker stopping...")
			return
		default:
			for _, symbol := range w.symbols {
				resp, err := client.Get(fmt.Sprintf(values.Url, symbol))
				atomic.AddUint64(&w.RequestCount, 1)
				if err != nil {
					fmt.Println("Error fetching price:", err)
					continue
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					resp.Body.Close()
					continue
				}
				resp.Body.Close()

				var ticker dto.TickerResponse
				if err := json.Unmarshal(body, &ticker); err != nil {
					fmt.Println("Error unmarshalling response:", err)
					continue
				}

				if oldPrice, ok := tickerPrices[ticker.Symbol]; ok && oldPrice != ticker.Price {
					w.UpdateChan <- TickerUpdate{Symbol: ticker.Symbol, Price: ticker.Price}
				}
				tickerPrices[ticker.Symbol] = ticker.Price
			}
			// to make requests after 1 second each
			time.Sleep(1 * time.Second)
		}
	}
}

func (w *Worker) GetRequestsCount() int {
	return int(atomic.LoadUint64(&w.RequestCount))
}

func (w *Worker) StopWorker(ctx context.Context) {
	close(w.StopChan)
}
