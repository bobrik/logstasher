package logstasher

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

// Writer wraps another writer and writes json events to it
type Writer struct {
	e      *json.Encoder
	t      string
	tags   []string
	fields map[string]interface{}
}

// NewWriter creates new wrapper around existing writer
func NewWriter(w io.Writer, t string, tags []string, fields map[string]interface{}) Writer {
	return Writer{
		e:      json.NewEncoder(w),
		t:      t,
		tags:   tags,
		fields: fields,
	}
}

func (w Writer) Write(p []byte) (n int, err error) {
	e := Event{
		Version:   1,
		Type:      w.t,
		Timestamp: time.Now(),
		Message:   strings.TrimSpace(string(p)),
		Tags:      w.tags,
		Fields:    w.fields,
	}

	return len(p), w.e.Encode(e)
}
