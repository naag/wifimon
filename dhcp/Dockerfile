FROM golang:1.16-alpine AS base
WORKDIR /go/src/github.com/naag/wifimon/dhcp
ADD . .
RUN go install ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/bin/dhcp .
ENTRYPOINT /app/dhcp
