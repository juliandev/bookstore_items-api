FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/myapp src/main.go

FROM scratch
COPY --from=build /go/bin/myapp /go/bin/myapp
CMD ["/go/bin/myapp"]
