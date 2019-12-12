FROM golang:1.12.7-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/kaansari/spidey/catalog
COPY vendor ../vendor
COPY catalog ./
RUN go build -o /go/bin/app ./cmd/catalog/main.go

FROM alpine:3.10
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
