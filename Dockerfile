FROM golang:alpine AS build
WORKDIR /go/src/github.com/juliandev/bookstore_items-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/github.com/juliandev/bookstore_items-api src/main.go

FROM scratch
COPY --from=build /go/bin/github.com/juliandev/bookstore_items-api /go/bin/github.com/juliandev/bookstore_items-api
CMD ["/go/bin/github.com/juliandev/bookstore_items-api"]
