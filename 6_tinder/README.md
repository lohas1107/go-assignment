# Tinder Matching

## Project Layout
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

## API List
> Host: [http://localhost:8080](http://localhost:8080)

### Query Possible Singles
Request
```shell
GET /v1/singles?most_possible={count}
```
Response
```shell
[{"gender":"GIRL","height":165,"wantedDates":1},{"gender":"GIRL","height":165,"wantedDates":1}]
```

### Add Single and Match
Request
```shell
POST /v1/singles 
``` 
Response
```shell
[{"gender":"GIRL","height":165,"wantedDates":1},{"gender":"GIRL","height":165,"wantedDates":1}]
```

### Remove Single
Request
```shell
DELETE /v1/singles/{id}
``` 
Response
```shell
null
```