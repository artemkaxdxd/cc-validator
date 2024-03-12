FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

RUN go build -ldflags="-w -s" -o /go/bin/result

FROM alpine

COPY --from=builder /go/bin/result /go/bin/result
ENV SERVER_PORT=8080
EXPOSE 8080

CMD ["/go/bin/result"]
