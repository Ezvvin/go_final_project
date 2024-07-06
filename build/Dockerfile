FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 7540

RUN GOOS=${TODO_GOOS} GOARCH=${TODO_GOARCH} go build -o /my_app ./cmd

CMD ["/my_app"]