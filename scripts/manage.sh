#!/usr/bin/env bash
set -euo pipefail

# ---------------------------------------------
# Carrega variáveis do .env
# ---------------------------------------------
load_env() {
  if [ -f .env ]; then
    set -a
    source .env
    set +a
  fi
}

load_env

: "${DB_USER:?DB_USER não definido}"
: "${DB_PASSWORD:?DB_PASSWORD não definido}"
: "${DB_PORT:?DB_PORT não definido}"
: "${DB_NAME:?DB_NAME não definido}"

# Host temporário usando IP de outro computador na rede local
DB_HOST="192.168.1.9"

DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
MIGRATIONS_PATH="./migrations"

# ---------------------------------------------
# Funções de help
# ---------------------------------------------
show_help() {
  cat <<EOF
Uso: $0 <comando> [subcomando]

Comandos:
  migrate             | Gerencia migrations
  ├── up              | Aplica todas as migrations pendentes
  ├── down            | Reverte a última migration
  ├── force <version> | Força versão
  ├── version         | Mostra versão atual
  └── create <name>   | Cria nova migration

  seed                | Executa o seed para popular o banco de dados

  test                | Executa testes unitários
EOF
}

# ---------------------------------------------
# Funções de migration
# ---------------------------------------------
migrate_up()      { migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" up; }
migrate_down()    { migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" down 1; }
migrate_force()   { migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" force "${1:?Você precisa informar a versão}"; }
migrate_version() { migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" version; }
migrate_create()  { migrate create -ext sql -dir "$MIGRATIONS_PATH" -seq "${1:?Você precisa informar o nome da migration}"; }

# ---------------------------------------------
# Função de seed
# ---------------------------------------------
run_seed() {
  echo "Executando seed..."
  
  # Se estiver rodando via docker-compose
  docker compose run --rm seed
}

# ---------------------------------------------
# Função de testes unitários
# ---------------------------------------------
run_tests() {
  echo "Executando testes unitários..."

  # Docker compose com volume apontando para raiz do projeto, para enxergar go.mod
  docker compose run --rm test-run
}

# ---------------------------------------------
# Dispatcher de comandos
# ---------------------------------------------
command="${1:-}"
subcommand="${2:-}"

case "$command" in
  migrate)
    case "$subcommand" in
      up|down|version) "migrate_${subcommand}" ;;
      force) migrate_force "${3:-}" ;;
      create) migrate_create "${3:-}" ;;
      *) show_help; exit 1 ;;
    esac
    ;;
  seed) run_seed ;;
  test) run_tests ;;
  *) show_help; exit 1 ;;
esac