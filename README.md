# GoFlowDesk

Projeto inicial para API seguindo Clean Architecture e DDD.

A estrutura mínima é:

```
cmd/api         # ponto de entrada
internal/
  domain/        # entidades e regras de domínio
  application/   # casos de uso e serviços
  interfaces/    # adaptadores (HTTP, CLI, etc.)
  infrastructure/# implementações (DB, cache, mensageria...)
pkg/            # bibliotecas auxiliares
```

Edite os pacotes internos e implemente seus casos de uso. Este repositório
serve como ponto de partida para o desenvolvimento.
