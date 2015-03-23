package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bobrik/logstasher"
)

type outWrapper struct {
	p string
	f *os.File
	m sync.Mutex
}

func newOutWrapper(path string, interval int) (*outWrapper, error) {
	w := &outWrapper{path, nil, sync.Mutex{}}
	err := w.reopen()
	if err != nil {
		return nil, err
	}

	go w.monitorFile(interval)

	return w, nil
}

func (w *outWrapper) reopen() error {
	w.m.Lock()
	defer w.m.Unlock()

	f, err := os.OpenFile(w.p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	if w.f != nil {
		w.f.Close()
	}

	w.f = f

	return nil
}

func (w *outWrapper) Write(p []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()

	return w.f.Write(p)
}

func (w *outWrapper) monitorFile(i int) {
	for {
		time.Sleep(time.Second * time.Duration(i))

		fs, err := os.Stat(w.p)
		if err != nil && os.IsNotExist(err) {
			if os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, "looks like output was deleted")
				w.reopen()
			} else {
				fmt.Fprintln(os.Stderr, "error doing stat on open file:", w.p)
			}

			continue
		}

		cs, err := w.f.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error doing stat on open file:", w.p)
			continue
		}

		if !os.SameFile(fs, cs) {
			fmt.Fprintln(os.Stderr, "looks like output was rotated")
			w.reopen()
		}
	}
}

func main() {
	fields := flag.String("fields", "", "json-encoded map of fields")
	out := flag.String("out", "", "output file to use instead of stdout")
	interval := flag.Int("i", 1, "interval to check for log rotation")
	flag.Parse()

	extra, err := parseFields(*fields)
	if err != nil {
		log.Fatal(err)
	}

	o, err := getOutput(*out, *interval)
	if err != nil {
		log.Fatal(err)
	}

	w := logstasher.NewWriter(o, extra)

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

func getOutput(p string, i int) (io.Writer, error) {
	if p == "" {
		return os.Stdout, nil
	}

	return newOutWrapper(p, i)
}

func parseFields(f string) (map[string]interface{}, error) {
	r := map[string]interface{}{}

	if f == "" {
		return r, nil
	}

	return r, json.Unmarshal([]byte(f), &r)
}
