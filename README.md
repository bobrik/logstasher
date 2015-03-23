# Logstasher

Logstasher provides logstash-friendly logging interface for Go.
It wraps any writer and writes logstash events to it.

This is how output lines look like:

```json
{"@timestamp":"2015-03-22T13:34:43.973174775Z","@version":1,"message":"hello","shards":[1,2,3],"tags":["wow","such"]}
```

## Usage

```go
f := map[string]interface{}{
    "app":    "myapp",
    "things": []string{"something", "another"},
}

w := logstasher.NewWriter(os.Stdout, f)

log.SetFlags(0)
log.SetOutput(w)

log.Println("hey, logstash!")
```

## CLI

Logstasher provides command `logstasher` that can be used in unix pipes.
It reads lines from stdin, turns them into json events and writes to stdout:

```
λ echo reporting from $(hostname -s) | logstasher -fields '{"app":"myapp","pi":3.14}'
{"@timestamp":"2015-03-22T13:33:42.344282854Z","@version":1,"app":"myapp","message":"reporting from hyperion","pi":3.14}
```

You can also pass `-out` option with filename to write and `logstasher` would
happily write to that file, detecting log rotation and removal.

## License

MIT
