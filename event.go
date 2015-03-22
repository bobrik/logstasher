package logstasher

import "time"

// Event represents logstash event
type Event struct {
	Version   int                    `json:"@version"`
	Type      string                 `json:"@type"`
	Timestamp time.Time              `json:"@timestamp"`
	Message   string                 `json:"@message"`
	Tags      []string               `json:"@tags,omitempty"`
	Fields    map[string]interface{} `json:"@fields,omitempty"`
}
