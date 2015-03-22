package logstasher_test

import (
	"log"
	"os"

	"github.com/bobrik/logstasher"
)

func ExampleWriter() {
	f := map[string]interface{}{
		"app":    "myapp",
		"things": []string{"something", "another"},
	}

	w := logstasher.NewWriter(os.Stdout, f)

	log.SetFlags(0)
	log.SetOutput(w)

	log.Println("hey, logstash!")
}
