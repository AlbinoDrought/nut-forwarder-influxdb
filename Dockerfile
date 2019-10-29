FROM golang:alpine as builder

RUN apk add -U --no-cache git ca-certificates

COPY . $GOPATH/src/github.com/AlbinoDrought/nut-forwarder-influxdb
WORKDIR $GOPATH/src/github.com/AlbinoDrought/nut-forwarder-influxdb

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go get -d -v && go build -a -installsuffix cgo -o /go/bin/nut-forwarder-influxdb

FROM scratch

WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/nut-forwarder-influxdb /go/bin/nut-forwarder-influxdb

CMD ["/go/bin/nut-forwarder-influxdb"]
