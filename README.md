# Monitoring package

Monitoring and tracing package for go services.

This package implements a server for monitoring and tracing

Launch a server like this:

```go
ep, err := NewServer(":0")
if err != nil {
    t.Fatalf("Got error creating endpoint: %v", err)
}

if err := ep.Start(); err != nil {
    t.Fatalf("Got error starting endpoint: %v", err)
}
defer ep.Shutdown()
```

`:0` will listen on an auto-assigned port on all interfaces. Note that
`localhost:0` might be `127.0.0.1` or `::1/128`.

Use the `ServerURL()` method to retrieve the server address.
