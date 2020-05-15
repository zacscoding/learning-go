// Modification for custom reporter code example
// from https://github.com/openzipkin/zipkin-go/blob/master/reporter/http/http.go
package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// ===========================
// common for reporter impls
// ===========================

type Reporter interface {
	Close() error // Close the reporter
}

// ===========================
// http based reporter
// ===========================

// defaults
const (
	defaultTimeout       = 5 * time.Second // timeout for http request in seconds
	defaultBatchInterval = 1 * time.Second // BatchInterval in seconds
	defaultBatchSize     = 100
	defaultMaxBacklog    = 1000
)

// HTTPDoer will do a request to the Zipkin HTTP Collector
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// httpReporter will send spans to a Zipkin HTTP Collector using Zipkin V2 API.
type httpReporter struct {
	url           string
	client        HTTPDoer
	logger        *log.Logger
	batchInterval time.Duration
	batchSize     int
	maxBacklog    int
	batchMtx      *sync.Mutex
	batch         []*interface{}
	dataC         chan *interface{}
	sendC         chan struct{}
	quit          chan struct{}
	shutdown      chan error
	reqCallback   RequestCallbackFn
	reqTimeout    time.Duration
	serializer    reporter.SpanSerializer
}

// Send implements reporter
func (r *httpReporter) Send(s model.SpanModel) {
	r.spanC <- &s
}

// Close implements reporter
func (r *httpReporter) Close() error {
	// TODO : flush and write files if failed to report
	close(r.quit)
	return <-r.shutdown
}

func (r *httpReporter) loop() {
	var (
		nextSend   = time.Now().Add(r.batchInterval)
		ticker     = time.NewTicker(r.batchInterval / 10)
		tickerChan = ticker.C
	)
	defer ticker.Stop()

	for {
		select {
		case data := <-r.dataC:
			currentBatchSize := r.append(data)
			if currentBatchSize >= r.batchSize {
				nextSend = time.Now().Add(r.batchInterval)
				r.enqueueSend()
			}
		case <-tickerChan:
			if time.Now().After(nextSend) {
				nextSend = time.Now().Add(r.batchInterval)
				r.enqueueSend()
			}
		case <-r.quit:
			close(r.sendC)
			return
		}
	}
}

func (r *httpReporter) append(data *interface{}) (newBatchSize int) {
	r.batchMtx.Lock()

	r.batch = append(r.batch, data)
	if len(r.batch) > r.maxBacklog {
		dispose := len(r.batch) - r.maxBacklog
		r.logger.Printf("backlog too long, disposing %d spans", dispose)
		r.batch = r.batch[dispose:]
	}
	newBatchSize = len(r.batch)

	r.batchMtx.Unlock()
	return
}
