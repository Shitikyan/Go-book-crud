FROM golang:1.16-alpine AS builder

WORKDIR $GOPATH/src/go-book-crud/

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-book-crud

FROM scratch

WORKDIR $GOPATH/src/go-book-crud/

COPY --from=builder /go/bin/go-book-crud /go/bin/go-book-crud

ENTRYPOINT ["/go/bin/go-book-crud"]