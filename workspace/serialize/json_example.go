package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title    string    `json:"title"`
	Internal string    `json:"-"`
	Empty    string    `json:"empty,omitempty"`
	Status   status    `json:"status"`
	ID       int64     `json:"id,string"`
	Deadline *Deadline `json:"deadline"`
}

func main() {
	// marshal
	t := Task{
		"Laundry",
		"Internal value",
		"",
		DONE,
		1,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output
	// {"title":"Laundry","Status":"DONE","id":"1","Deadline":"2015-08-16T15:43:00Z"}
}

// MarshalJSON implements the json.Marshaler interface.
func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}
