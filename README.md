# GoFlowDesk

API simples de Help Desk escrita em Go. Este repositório foi refatorado
como um **MVP leve** para uso em portfólio: a arquitetura é mínima, fácil
de entender e executar.

## Estrutura atual

```
cmd/                  # executáveis (api + seed)
internal/
  app/                # bootstrapping da aplicação
  domain/             # entidades de domínio
  repository/         # acesso a dados (Postgres)
  service/            # lógica de negócio
  handler/            # HTTP handlers + render helpers
  router/             # chi router e rotas
  infrastructure/     # server, db, logging, etc.
```

Esses pacotes representam a camada HTTP ➜ serviço ➜ repositório. Não há
interfaces excessivas nem contêiner de injeção.

## Executando localmente

1. inicie o banco PostgreSQL com `docker-compose up -d` (contém somente o
   serviço de banco).
2. configure as variáveis `DATABASE_URL` / `DATABASE_URL_TEST` conforme
   necessário.
3. rode a API: `go run ./cmd/api`.
4. para popular dados use `go run ./cmd/seed`.
5. testes: `DATABASE_URL_TEST=... go test ./...`.


## Docker

A imagem construída pelo `Dockerfile` não é usada em `docker-compose`; ela é
apenas um placeholder. O compose traz apenas o container de Postgres, já que
você chama a aplicação diretamente na máquina local.

```sh
# start DB only
docker-compose up -d
# ...then run API locally as acima
```


---

Sinta‑se livre para estender o projeto com tickets, autenticação, etc.
