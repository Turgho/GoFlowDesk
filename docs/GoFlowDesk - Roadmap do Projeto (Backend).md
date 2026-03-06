# Roadmap Backend — GoFlowDesk

Este documento define um roteiro **passo a passo** para construir a API
GoFlowDesk com foco em **Clean Architecture** e **Domain‑Driven Design (DDD)**.
A ideia é manter uma base sólida na fase inicial e, a partir dela, evoluir de
forma incremental para mensageria, workers, observabilidade e implantação.

> 🧱 A estrutura de pastas sugerida segue as camadas internas de DDD e pode
>     evoluir conforme novas responsabilidades surgem (cache, mensageria,
>     workers, observabilidade etc.). Um esqueleto inicial mínimo fica assim
>     mas deve permitir crescer de forma ordenada:
>
> ```text
> .
> ├── cmd/
> │   ├── api/
> │   │   └── main.go
> │   ├── worker-sla/
> │   │   └── main.go
> │   ├── worker-notifications/
> │   │   └── main.go
> │   └── ...
> │
> ├── internal/
> │   ├── bootstrap/
> │   │   ├── app.go
> │   │   └── container.go
> │   │
> │   ├── domain/
> │   │   ├── user.go
> │   │   ├── ticket.go
> │   │   ├── ticket_status.go
> │   │   ├── sla_rule.go
> │   │   ├── log_entry.go
> │   │   ├── errors.go
> │   │   └── event/
> │   │       ├── ticket_created.go
> │   │       └── sla_breached.go
> │   │
> │   ├── application/
> │   │   ├── ports/
> │   │   │   ├── user_repository.go
> │   │   │   └── ticket_repository.go
> │   │   ├── security/
> │   │   │   └── password_hasher.go
> │   │   ├── user/
> │   │   │   ├── create.go
> │   │   │   ├── get.go
> │   │   │   ├── list.go
> │   │   │   ├── authenticate.go
> │   │   │   ├── update.go
> │   │   │   └── delete.go
> │   │   └── ticket/
> │   │       ├── create.go
> │   │       ├── update_status.go
> │   │       └── list.go
> │   │
> │   ├── adapters/
> │   │   ├── http/
> │   │   │   ├── handler/
> │   │   │   │   ├── user_handler.go
> │   │   │   │   └── ticket_handler.go
> │   │   │   ├── middleware/
> │   │   │   │   ├── auth.go
> │   │   │   │   └── ratelimit.go
> │   │   │   ├── dto/
> │   │   │   │   └── user_dto.go
> │   │   │   ├── routes.go
> │   │   │   └── json.go
> │   │   ├── postgres/
> │   │   │   ├── user_repository.go
> │   │   │   ├── ticket_repository.go
> │   │   │   └── queries.sql
> │   │   ├── redis/
> │   │   │   └── cache.go
> │   │   └── rabbitmq/
> │   │       └── producer.go
> │   │
> │   └── infrastructure/
> │       ├── database/
> │       │   └── postgres.go
> │       ├── messaging/
> │       │   └── rabbitmq.go
> │       ├── logging/
> │       │   └── logger.go
> │       ├── metrics/
> │       │   └── prometheus.go
> │       └── server.go
> │
> ├── pkg/
> │   └── config/
> │       └── config.go
> ├── migrations/
> ├── scripts/
> ├── docs/
> ├── docker-compose.yml
> ├── Dockerfile
> ├── go.mod
> └── README.md
> ```

---

## 🔵 FASE 1 — Fundação (Estrutura, DB e CRUD)

Objetivo: ter a API funcionando, organizada em camadas e com CRUD básico de
usuários e tickets.

### 1. Estruturação e Clean Architecture

1. Criar a árvore de pastas acima e colocar responsabilidades em cada camada.
2. Definir contratos através de interfaces (_ports_) para permitir inversão de
dependências.
3. Evitar que camadas internas-importem externas; apenas _application_ usa
   _domain_ e _interfaces_ depende de _application_.

### 2. Banco de Dados (Postgres)

- Modelar entidades DDD: `User`, `Ticket`, `TicketStatus`, `SlaRule`, `LogEntry`.
- Criar migrations com `golang-migrate` ou equivalente.
- Implementar *repositories* na camada de infraestrutura.
- Escrever testes para conexões e operações básicas.

### 3. CRUD de Tickets (e Users básicos)

- Expor endpoints HTTP em `interfaces/http`:
  - `POST /tickets` → criar
  - `PATCH /tickets/:id/status` → alterar status
  - `GET /tickets` → listar com filtros/paginação
  - `DELETE /tickets/:id` → soft delete
- Implementar casos de uso em `application/ticket`.
- Cobrir com testes unitários e integração mínima.

### 4. Redis (Cache)

- Implementar padrão *Cache‑Aside*:
  - Verificar cache antes de ir ao banco.
  - Atualizar cache em criações/atualizações.
  - Definir TTL apropriado.

---

## 🟡 FASE 2 — Mensageria e Eventos Assíncronos

Permitir comunicação desacoplada entre a API e os processos de background.

### 5. Escolher Broker

RabbitMQ é recomendado para filas de tarefas. Kafka pode ser usado para
streaming de eventos, mas não é necessário inicialmente.

### 6. Eventos de Domínio

- Definir eventos como `TicketCreated`, `TicketUpdated`, `SlaBreached`.
- Serializar e publicar em `infrastructure/messaging`.
- Produzir eventos nos serviços de criação/alteração de tickets.

### 7. Workers Independentes

- **worker‑sla** – aplica regras de SLA e publica `SlaBreached`.
- **worker‑notifications** – envia notificações/alertas.
- **worker‑reports** – gera métricas periódicas.

Cada worker:
1. Conecta-se ao broker.
2. Consome fila específica.
3. Aplica lógica de aplicação/importa use cases.

Containerizar cada worker de maneira isolada.

---

## 🟠 FASE 3 — Motor de SLA

- Regras configuráveis por prioridade, cliente ou tipo de ticket.
- Worker SLA verifica tickets abertos e atualiza status quando uma violação
  ocorre.
- Publicar evento `SlaBreached` para que outros consumidores reajam.

---

## 🔴 FASE 4 — Observabilidade

- Utilizar logs estruturados (`slog`, `zap`).
- Middleware para `request-id`/correlation-id.
- Expor métricas Prometheus em `/metrics`.
- Health endpoints: `/health/liveness` e `/health/readiness`.
- (Opcional) Dashboards Grafana.

---

## 🟣 FASE 5 — Segurança

- Autenticação com JWT.
- Middleware de autorização baseado em roles.
- Limitação de taxa (rate‑limit) usando Redis.
- Configurar CORS adequadamente.
- Validações de entrada na borda.

---

## 🟤 FASE 6 — Dockerização Profissional

- Multi‑stage build para API e cada worker.
- `docker-compose.yml` com:
  - `api`, `postgres`, `redis`, `rabbitmq`, e os workers.
- Prover healthchecks e rede isolada.

---

## ⚫ FASE 7 — Deploy em Kubernetes (Opcional)

- Criar Deployments e Services (ClusterIP/LoadBalancer).
- Usar ConfigMaps e Secrets para configuração.
- Configurar probes de liveness/readiness.
- HPA para escalabilidade horizontal.

---

## 📊 Ordem Recomendada de Execução

1. Estrutura inicial e usuários CRUD.
2. Tickets CRUD + migrations.
3. Cache Redis.
4. Refatoração para extração de interfaces e desacoplamento.
5. Mensageria (RabbitMQ) e eventos.
6. Workers de background.
7. SLA engine.
8. Observabilidade.
9. Segurança.
10. Testes automatizados e CI.
11. Docker completo.
12. Kubernetes ou similar.

---

## 🎯 Arquitetura Conceitual

```text
Cliente → API Go
              ↓
         PostgreSQL    Redis ←─┐
              ↓           ↑   │
        RabbitMQ         └─ workers (cache/invalidate)
           ↓    ↓    ↓
        SLA  Notif  Reports
```

Seguindo essas fases com atenção à separação de capas e ao domínio, o
roadmap serve tanto como guia de implementação quanto como documentação para
novos desenvolvedores.
