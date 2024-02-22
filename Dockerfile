FROM golang:1.21 as builder

COPY . /queue-reader
WORKDIR /queue-reader
RUN go build .

FROM ubuntu
COPY --from=builder /queue-reader/queue-reader /queue-reader
ENTRYPOINT ["/queue-reader"]
