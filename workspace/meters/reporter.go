// Modification for custom reporter code example
// from https://github.com/openzipkin/zipkin-go/blob/master/reporter/http/http.go
package main

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// ===========================
// common for reporter interface
// ===========================

var (
	ErrParse = errors.New("can't marshalling spans batch")
	ErrCanRetryable = errors.New("")
)

type Reporter interface {
	Send(span ApiCallSpan)
	Close() error // Close the reporter
}

// ===========================
// http based reporter
// ===========================

// defaults
const (
	defaultTimeout       = 5 * time.Second // timeout for http request in seconds
	defaultBatchInterval = 5 * time.Second // BatchInterval in seconds
	defaultBatchSize     = 10
	defaultMaxBacklog    = 20
)

type ApiCallSpan struct {
	Url      string
	FullPath string
	Headers  map[string][]string
}

// HTTPDoer will do a request to the Zipkin HTTP Collector
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Send implements reporter
func (r *httpReporter) Send(s ApiCallSpan) {
	r.spanC <- &s
}

// Close implements reporter
func (r *httpReporter) Close() error {
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
		case data := <-r.spanC:
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

func (r *httpReporter) append(span *ApiCallSpan) (newBatchSize int) {
	r.batchMtx.Lock()

	r.batch = append(r.batch, span)
	if r.maxBacklog > 0 && len(r.batch) > r.maxBacklog {
		dispose := len(r.batch) - r.maxBacklog
		r.logger.Printf("backlog too long, disposing %d spans", dispose)
		r.batch = r.batch[dispose:]
	}
	newBatchSize = len(r.batch)

	r.batchMtx.Unlock()
	return
}

func (r *httpReporter) sendLoop() {
	for range r.sendC {
		r.logger.Println("sendLoop()")
		_ = r.sendBatch()
	}
	r.shutdown <- r.sendBatch()
}

func (r *httpReporter) enqueueSend() {
	select {
	case r.sendC <- struct{}{}:
	default:
		// Do nothing if there's a pending send request already
	}
}

// TODO : add handle failures
func (r *httpReporter) sendBatch() error {
	// Select all current spans in the batch to be sent
	r.batchMtx.Lock()
	sendBatch := r.batch[:]
	r.batchMtx.Unlock()

	if len(sendBatch) == 0 {
		return nil
	}

	body, err := r.serializer.Serialize(sendBatch)
	if err != nil {
		r.logger.Printf("failed when marshalling the spans batch: %s\n", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", r.url, bytes.NewReader(body))
	if err != nil {
		r.logger.Printf("failed when creating the request: %s\n", err.Error())
		return err
	}
	req.Header.Set("Content-Type", r.serializer.ContentType())
	if r.reqCallback != nil {
		r.reqCallback(req)
	}

	ctx, cancel := context.WithTimeout(req.Context(), r.reqTimeout)
	defer cancel()

	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		r.logger.Printf("failed to send the request: %s\n", err.Error())
		return err
	}
	_ = resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		r.logger.Printf("failed the request with status code %d\n", resp.StatusCode)
	}

	// Remove sent spans from the batch even if they were not saved
	r.batchMtx.Lock()
	r.batch = r.batch[len(sendBatch):]
	r.batchMtx.Unlock()

	return nil
}

// RequestCallbackFn receives the initialized request from the Collector before
// sending it over the wire. This allows one to plug in additional headers or
// do other customization.
type RequestCallbackFn func(*http.Request)

// ReporterOption sets a parameter for the HTTP Reporter
type ReporterOption func(r *httpReporter)

// httpReporter will send spans to a Zipkin HTTP Collector using Zipkin V2 API.
type httpReporter struct {
	url           string
	client        HTTPDoer
	logger        *log.Logger
	batchInterval time.Duration
	batchSize     int
	maxBacklog    int
	batchMtx      *sync.Mutex
	batch         []*ApiCallSpan
	spanC         chan *ApiCallSpan
	sendC         chan struct{}
	quit          chan struct{}
	shutdown      chan error
	reqCallback   RequestCallbackFn
	reqTimeout    time.Duration
	serializer    SpanSerializer
}

// MaxBacklog sets the maximum backlog size. When batch size reaches this
// threshold, spans from the beginning of the batch will be disposed.
func MaxBacklog(n int) ReporterOption {
	return func(r *httpReporter) { r.maxBacklog = n }
}

// NewReporter returns a new HTTP Reporter.
// url should be the endpoint to send the spans to, e.g.
// http://localhost:9411/api/v2/spans
func NewReporter(url string, opts ...ReporterOption) Reporter {
	r := httpReporter{
		url:           url,
		logger:        log.New(os.Stderr, "", log.LstdFlags),
		client:        &http.Client{},
		batchInterval: defaultBatchInterval,
		batchSize:     defaultBatchSize,
		maxBacklog:    defaultMaxBacklog,
		batch:         []*ApiCallSpan{},
		spanC:         make(chan *ApiCallSpan),
		sendC:         make(chan struct{}, 1),
		quit:          make(chan struct{}, 1),
		shutdown:      make(chan error, 1),
		batchMtx:      &sync.Mutex{},
		serializer:    JSONSerializer{},
		reqTimeout:    defaultTimeout,
	}

	for _, opt := range opts {
		opt(&r)
	}

	go r.loop()
	go r.sendLoop()

	return &r
}
