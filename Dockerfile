# syntax=docker/dockerfile:1

FROM golang:1.22.2-bookworm

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /pipeline
CMD ["/pipeline"]