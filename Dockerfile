FROM golang:1.22.1 as builder
LABEL authors="Anurag"

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3.14 as production

RUN apk add --no-cache ca-certificates

COPY --from=builder app .

EXPOSE 9090:9090

CMD ./app

