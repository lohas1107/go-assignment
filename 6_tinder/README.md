# Tinder Matching

### Project Layout
```shell
.
|-- /api                        # API documentation
|-- /build                      # Docker image build
|-- /cmd
|   `-- /matching
|       |-- /router             # Endpoint handlers
|       `-- main.go             # Application entry point
|-- /internal
|   |-- /matching               # Business logic
|   `-- /pkg
|-- /test                       # Integration tests
|-- Makefile
|-- README.md
`-- go.mod
```