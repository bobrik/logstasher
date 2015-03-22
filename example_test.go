package logstasher_test

import (
	"log"
	"os"

	"github.com/bobrik/logstasher"
)

func ExampleLogWriter() {
	w := logstasher.NewWriter(os.Stdout, "myapp", []string{"nice", "tags"}, nil)

	log.SetFlags(0)
	log.SetOutput(w)

	log.Println("hey, logstash!")
}
