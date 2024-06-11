package test

import (
	"sync/atomic"
	"testing"

	"github.com/NessibeliY/binance-ticker-parser/internal/worker"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

type MockHttpClient struct {
	mock.Mock
}

func TestWorker_GetRequestsCount(t *testing.T) {
	w := worker.NewWorker([]string{"BTCUSDT"}, nil)
	atomic.AddUint64(&w.RequestCount, 5)
	assert.Equal(t, 5, w.GetRequestsCount())
}
