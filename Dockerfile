FROM golang:1.19

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o final-go-kbtu ./cmd/web

CMD ["./final-go-kbtu"]

# docker-compose up final-go-kbtu
