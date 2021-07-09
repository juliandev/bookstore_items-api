FROM golang:alpine AS build

WORKDIR /go/src/github.com/juliandev/bookstore_items-api

COPY . .

RUN go build -o /go/bin/github.com/juliandev/bookstore_items-api src/main.go

FROM scratch

COPY --from=build /go/bin/github.com/juliandev/bookstore_items-api /go/bin/github.com/juliandev/bookstore_items-api

EXPOSE 9000

CMD ["./bookstore_items-api"]
