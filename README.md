# ProjetoFinalNeoCamp 
## Como rodar
```bash
cp .env.example .env
docker compose up --build -d
# ver logs
docker compose logs -f db
docker compose logs -f api
```

API básica:
- `GET /healthz` (verifica conexão com o MySQL)
- `GET /ping` → `{"pong": true}`

## Próximos passos
- Implementar handlers e repositórios de `ingredient`, `dish`, `menu`.
- Montar rotas no `internal/http/router.go`.
- Criar testes com `go test ./...`.
