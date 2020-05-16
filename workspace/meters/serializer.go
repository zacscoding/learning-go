package main

import (
	"encoding/json"
)

// SpanSerializer describes the methods needed for allowing to set Span encoding
// type for the various Zipkin transports.
type SpanSerializer interface {
	Serialize([]*ApiCallSpan) ([]byte, error)
	ContentType() string
}

// JSONSerializer implements the default JSON encoding SpanSerializer.
type JSONSerializer struct{}

// Serialize takes an array of Zipkin SpanModel objects and returns a JSON
// encoding of it.
func (JSONSerializer) Serialize(spans []*ApiCallSpan) ([]byte, error) {
	return json.Marshal(spans)
}

// ContentType returns the ContentType needed for this encoding.
func (JSONSerializer) ContentType() string {
	return "application/json"
}
