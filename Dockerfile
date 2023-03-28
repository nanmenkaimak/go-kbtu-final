FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o final-go-kbtu ./cmd/web

CMD ["./final-go-kbtu"]
