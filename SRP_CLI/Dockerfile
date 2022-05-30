FROM golang:latest

WORKDIR /SRP_server
RUN bash -c 'mkdir -p /SRP_server/{cmd/server,internal,pkg}'

COPY cmd/server /SRP_server/cmd/server
COPY internal /SRP_server/internal
COPY pkg /SRP_server/pkg
COPY go.mod go.sum config_server.yml /SRP_server/

RUN go mod tidy -go 1.16
RUN go mod tidy -go 1.17
RUN go build ./cmd/server/.