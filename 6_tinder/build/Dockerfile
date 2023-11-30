FROM golang:1.21.4-alpine AS build
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o ./build/bin/matching ./cmd/matching

FROM alpine
WORKDIR /app
COPY --from=build /src/build/bin/matching /app/
CMD ./matching