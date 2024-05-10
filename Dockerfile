# syntax=docker/dockerfile:1

FROM golang:1.22.2
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-test
CMD [ "/docker-go-test" ]
