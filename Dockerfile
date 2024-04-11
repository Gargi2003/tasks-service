FROM golang:1.20

WORKDIR /tasks-service

COPY . /tasks-service

EXPOSE 8081

RUN go mod download

RUN go build -o tasks-service .

CMD ["./tasks-service"]
