FROM golang:1.25.7-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goflowdesk-api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goflowdesk-cli ./cmd/cli


FROM alpine:3.23.3

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=build /app/goflowdesk-api /app/goflowdesk-api
COPY --from=build /app/goflowdesk-cli /app/goflowdesk-cli

RUN chmod +x /app/goflowdesk-api /app/goflowdesk-cli

CMD ["/app/goflowdesk-api"]