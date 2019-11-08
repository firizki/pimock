FROM golang:1.13.4-alpine3.10 AS builder

RUN apk update && apk add --no-cache git
RUN go get github.com/firizki/ailea

WORKDIR $GOPATH/src/github.com/firizki/ailea/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/ailea

FROM alpine:3.10

COPY --from=builder /go/bin/ailea /usr/bin/ailea
ADD responses/ responses/

EXPOSE 8080

ENTRYPOINT ["ailea"]
