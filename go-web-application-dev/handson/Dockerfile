FROM golang:1.19-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

FROM debian:bullseye-slim AS deploy

RUN apt update

COPY --from=builder /app/app .

CMD ["./app"]

# For hot reload
FROM golang:1.19-bullseye AS development

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1

CMD ["air"]
EXPOSE 80
