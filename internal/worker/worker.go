package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/NessibeliY/binance-ticker-parser/internal/dto"
)

const url = "https://api.binance.com/api/v3/ticker/price?symbol=%s"

type Worker struct {
	symbols      []string
	requestCount uint64
	stopChan     chan struct{}
}

func NewWorker(symbols []string) *Worker {
	return &Worker{
		symbols:  symbols,
		stopChan: make(chan struct{}),
	}
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	tickerPrices := make(map[string]string)

	for {
		select {
		case <-w.stopChan:
			fmt.Println("Worker stopping...")
			return
		default:
			for _, symbol := range w.symbols {
				resp, err := client.Get(fmt.Sprintf(url, symbol))
				atomic.AddUint64(&w.requestCount, 1)
				if err != nil {
					fmt.Println("Error fetching price:", err)
					continue
				}

				body, err := ioutil.ReadAll(resp.Body)
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
					fmt.Printf("%s price:%s changed\n", ticker.Symbol, ticker.Price)
				} else {
					fmt.Printf("%s price:%s\n", ticker.Symbol, ticker.Price)
				}
				tickerPrices[ticker.Symbol] = ticker.Price
			}
			// to make requests after 1 second each
			time.Sleep(1 * time.Second)
		}
	}
}

func (w *Worker) GetRequestsCount() int {
	return int(atomic.LoadUint64(&w.requestCount))
}

func (w *Worker) StopWorker(ctx context.Context) {
	close(w.stopChan)
}
