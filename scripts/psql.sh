#!/usr/bin/env bash
set -euo pipefail

# Nome do container
CONTAINER_NAME="goflowdesk-postgres-1"

# Carrega .env da raiz do projeto
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ENV_FILE="$SCRIPT_DIR/../.env"

if [ -f "$ENV_FILE" ]; then
  set -a
  source "$ENV_FILE"
  set +a
else
  echo ".env não encontrado em $ENV_FILE"
  exit 1
fi

: "${DB_USER:?DB_USER não definido}"
: "${DB_PASSWORD:?DB_PASSWORD não definido}"
: "${DB_NAME:?DB_NAME não definido}"

# Verifica se o container está rodando
if ! docker ps --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
  echo "❌ Container ${CONTAINER_NAME} não está rodando"
  exit 1
fi

echo "🔌 Conectando ao banco dentro do container..."

docker exec -it "$CONTAINER_NAME" \
  env PGPASSWORD="$DB_PASSWORD" \
  psql -U "$DB_USER" -d "$DB_NAME"