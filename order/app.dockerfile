FROM golang:1.12.7-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/kaansari/spidey/order
COPY vendor ../vendor
COPY account ../account
COPY catalog ../catalog
COPY order ./
RUN go build -o /go/bin/app ./cmd/order/main.go

FROM alpine:3.10
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
