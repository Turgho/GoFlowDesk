FROM golang:1.25.7-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build the Go application with CGO disabled for a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o goflowdesk ./cmd/api

FROM alpine:3.23.3
WORKDIR /app
COPY --from=build /app/goflowdesk .
COPY .env .

RUN apk add --no-cache curl

EXPOSE 8080
CMD ["./goflowdesk"]