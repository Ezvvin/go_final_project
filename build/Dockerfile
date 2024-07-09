FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=${TODO_GOOS} GOARCH=${TODO_GOARCH} go build -o /main ./cmd

CMD ["/main"]