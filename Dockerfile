# syntax=docker/dockerfile:1

# ------------------------
# Estágio de build (golang)
# ------------------------
FROM golang:1.25.7-alpine AS build-stage
WORKDIR /app

# Instala ferramentas básicas para build
RUN apk add --no-cache git

# Dependências
COPY go.mod go.sum ./
RUN go mod download

# Código fonte
COPY . .

# Compila a API e o seed
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/goflowdesk-api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/goflowdesk-seed ./cmd/seed

# ------------------------
# Estágio de testes
# ------------------------
FROM build-stage AS run-test-stage
WORKDIR /app
# Ao iniciar, executa todos os testes
ENTRYPOINT ["go", "test", "./...", "-v"]

# ------------------------
# Estágio final (produção) — Alpine leve
# ------------------------
FROM alpine:3.23 AS build-release-stage
WORKDIR /app

# Dependências básicas para rodar binário Go
RUN apk add --no-cache ca-certificates

# Copia os binários já compilados
COPY --from=build-stage /app/goflowdesk-api /app/goflowdesk-api
COPY --from=build-stage /app/goflowdesk-seed /app/goflowdesk-seed

# Permissões
RUN chmod +x /app/goflowdesk-api /app/goflowdesk-seed

# Ponto de entrada padrão: API
ENTRYPOINT ["/app/goflowdesk-api"]