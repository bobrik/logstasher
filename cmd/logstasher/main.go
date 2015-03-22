package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/bobrik/logstasher"
)

func main() {
	t := flag.String("type", "", "type for logstash (app name)")
	tags := flag.String("tags", "", "comma-separated list of tags")
	fields := flag.String("fields", "", "json-encoded map of fields")

	flag.Parse()

	if *t == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	extra, err := parseFields(*fields)
	if err != nil {
		log.Fatal(err)
	}

	w := logstasher.NewWriter(os.Stdout, *t, parseTags(*tags), extra)

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

func parseTags(t string) []string {
	if t == "" {
		return []string{}
	}

	return strings.Split(t, ",")
}

func parseFields(f string) (map[string]interface{}, error) {
	r := map[string]interface{}{}

	if f == "" {
		return r, nil
	}

	return r, json.Unmarshal([]byte(f), &r)
}
