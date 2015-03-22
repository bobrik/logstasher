# Logstasher

Logstasher provides logstash-friendly logging interface for Go.
It wraps any writer and writes logstash events to it.

This is how output lines look like:

```json
{"@version":1,"@type":"something","@timestamp":"2015-03-22T12:46:23.632679998Z","@message":"haha","@tags":["one","two"],"@fields":{"pi":3.14}}
```

## Usage

```go
w := logstasher.NewWriter(os.Stdout, "myapp", []string{"nice", "tags"}, nil)

log.SetFlags(0)
log.SetOutput(w)

log.Println("hey, logstash!")
```

## CLI

Logstasher provides command `logstasher` that can be used in unix pipes as filter.
It reads lines from stdin, turns them into logstash events and writes to stdout:

```
λ echo reporting from $(hostname -s) | logstasher -type something
{"@version":1,"@type":"something","@timestamp":"2015-03-22T12:45:23.991968768Z","@message":"reporting from hyperion"}
```

## License

MIT