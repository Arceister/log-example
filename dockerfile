FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go -o /log-example

EXPOSE 8080

CMD ["/log-example"]