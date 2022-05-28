FROM golang:latest

WORKDIR /SRP_web
RUN bash -c 'mkdir /SRP_web/{cert,cmd,internal,migrations,pkg,web}'

COPY cert /SRP_web/cert
COPY cmd /SRP_web/cmd
COPY internal /SRP_web/internal
COPY migrations /SRP_web/migrations
COPY pkg /SRP_web/pkg
COPY web /SRP_web/web
COPY go.mod go.sum /SRP_web/

RUN go mod tidy
RUN go build ./cmd/SRP_web/.