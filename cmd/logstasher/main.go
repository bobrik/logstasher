package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/bobrik/logstasher"
)

func main() {
	fields := flag.String("fields", "", "json-encoded map of fields")
	flag.Parse()

	extra, err := parseFields(*fields)
	if err != nil {
		log.Fatal(err)
	}

	w := logstasher.NewWriter(os.Stdout, extra)

	log.SetFlags(0)
	log.SetOutput(w)

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		log.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseFields(f string) (map[string]interface{}, error) {
	r := map[string]interface{}{}

	if f == "" {
		return r, nil
	}

	return r, json.Unmarshal([]byte(f), &r)
}
