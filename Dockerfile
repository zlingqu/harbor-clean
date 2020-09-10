FROM golang:1.15-alpine as builder
WORKDIR /data
ADD . .
RUN go build -o targets/harbor-clean

FROM alpine:3.12.0
WORKDIR /data
COPY --from=builder /data/targets/harbor-clean .
