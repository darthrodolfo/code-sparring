# AeroStack — Roadmap Passo a Passo

## Diagnóstico: Onde estamos

O `Program.cs` atual tem CRUD completo do AircraftV2 com SQLite (Minimal API), incluindo tags, conflicts, specs, geolocation, metadata e manual archive. Tudo funcional mas concentrado num único arquivo, sem separação de responsabilidades, sem testes, sem autenticação, sem validação formal, sem observabilidade e sem contrato padronizado.

O plano abaixo organiza **tudo que precisa acontecer antes de AI** e depois **toda a progressão de AI Integrations**, em ordem de execução.

---

## PARTE 1 — FUNDAÇÕES PRÉ-AI

A ideia aqui é nivelar o backend a um padrão "senior interview-ready" antes de tocar em qualquer LLM. Cada etapa se aplica primeiro ao .NET, depois replica no NestJS, e eventualmente nos backends satélites (Go, Python, Dart).

---

### ETAPA 0 — Contrato JSON Padronizado (AeroStack Contract)

**Objetivo:** Todo microserviço/API do AeroStack recebe e responde JSON da mesma forma. Qualquer dev ou consumidor sabe exatamente o que esperar.

**O que fazer:**

1. Definir convenção global: `camelCase` para todas as propriedades JSON (request e response).

2. Criar um envelope de resposta padronizado para todos os endpoints:
   ```
   Sucesso → { data, meta }
   Erro    → { error: { code, message, details, traceId } }
   Lista   → { data: [], meta: { page, pageSize, totalCount, totalPages } }
   ```

3. Definir formato de datas: ISO 8601 (`yyyy-MM-ddTHH:mm:ssZ`), DateOnly como `yyyy-MM-dd`.

4. Definir formato de enums na serialização: string (não int), ex: `"status": "active"` e não `"status": 0`.

5. Definir tratamento de nulls: propriedades null podem ser omitidas ou explícitas — escolher um e ser consistente.

6. Configurar no .NET: `JsonSerializerOptions` com `PropertyNamingPolicy = CamelCase`, `JsonStringEnumConverter`, `DefaultIgnoreCondition`.

7. Configurar no NestJS: interceptor global que wrapa toda response no envelope, `class-transformer` com `@Expose()`.

**Endpoints afetados:** Todos. Isso é transversal.

**Entregável por linguagem:**
- .NET: middleware/filter + `JsonSerializerOptions` global
- NestJS: interceptor global + DTO base
- Python (FastAPI): modelo Pydantic base + middleware
- Go (Gin/Echo): struct de envelope + middleware
- Dart (shelf/dart_frog): modelo base + middleware

---

### ETAPA 1 — Validações Completas

**Objetivo:** Nenhum dado inválido entra no sistema. Toda validação é explícita, testável e retorna erro padronizado.

**O que fazer:**

1. Validar todos os campos do `CreateAircraftV2Request`:
   - `model`: required, min 2 chars, max 100
   - `manufacturer`: required, min 2 chars, max 100
   - `serialNumber`: opcional, formato alfanumérico, unique quando presente
   - `yearOfManufacture`: required, range 1900–2030
   - `priceMillions`: required, > 0, max 2 casas decimais
   - `emptyWeightKg`: required, > 0
   - `status`: required, valor válido do enum
   - `role`: required, valor válido do enum
   - `tags`: required, min 1, max 20, cada tag max 50 chars, sem duplicatas
   - `firstFlightDate`: required, não pode ser futura
   - `lastMaintenanceTime`: required, não pode ser futura
   - `baseLocation.latitude`: range -90 a 90
   - `baseLocation.longitude`: range -180 a 180
   - `specs.maxSpeedKmh`: > 0
   - `specs.wingspanMeters`: > 0
   - `specs.rangeKm`: > 0
   - `specs.maxAltitudeMeters`: opcional, > 0 quando presente
   - `specs.flightEndurance`: > 0
   - `conflicts`: cada um com name required, startYear <= endYear, anos razoáveis
   - `metadata`: max 20 keys, key max 50 chars, value max 500 chars
   - `photoUrl`: URL válida quando presente
   - `manualArchive`: tamanho máximo (ex: 10MB)

2. Retornar erros no formato do contrato:
   ```
   { error: { code: "VALIDATION_ERROR", message: "...", details: [ { field, message } ] } }
   ```

3. Validar também no PUT (mesmas regras + id no path deve existir).

**Entregável por linguagem:**
- .NET: `FluentValidation` ou `MinimalApis.Extensions` com endpoint filters
- NestJS: `class-validator` + `ValidationPipe` global
- Python: Pydantic validators
- Go: validação manual ou `go-playground/validator`
- Dart: validação manual no handler

---

### ETAPA 2 — Idempotência

**Objetivo:** Toda operação de escrita é segura para retry. Nenhum POST duplica dado, nenhum PUT corrompe estado.

**O que fazer:**

1. POST `/aircraft-v2`: aceitar header `Idempotency-Key` (UUID). Se o mesmo key já foi processado, retornar o resultado original (201 com o mesmo body). Armazenar o resultado temporariamente (SQLite table ou Redis).

2. PUT `/aircraft-v2/{id}`: naturalmente idempotente (PUT replace completo). Garantir que se chamado N vezes com o mesmo body, o resultado é o mesmo.

3. DELETE `/aircraft-v2/{id}`: já idempotente (retorna 204 ou 404). Manter assim.

4. Criar tabela ou cache de idempotência:
   ```
   idempotency_keys
   ├── key (PK)
   ├── endpoint
   ├── response_status
   ├── response_body
   ├── created_at
   └── expires_at (TTL de 24h)
   ```

5. Middleware/filter que intercepta requests com `Idempotency-Key`, verifica cache, retorna se existe, processa se não.

**Entregável por linguagem:**
- .NET: endpoint filter ou middleware + SQLite/Redis
- NestJS: guard ou interceptor + Redis/SQLite
- Python/Go/Dart: middleware equivalente

---

### ETAPA 3 — Autenticação Mínima

**Objetivo:** Proteger os endpoints sem complexidade desnecessária para um lab. JWT simples, sem OAuth provider externo por enquanto.

**O que fazer:**

1. Criar endpoint `POST /auth/token` que recebe `{ apiKey }` (hardcoded no .env) e retorna um JWT com claims: `sub`, `role`, `iat`, `exp`.

2. Proteger todos os endpoints de escrita (POST, PUT, DELETE) com `[Authorize]` / guard.

3. GET pode ficar público ou com autenticação opcional (header presente = autenticado, ausente = anônimo).

4. Adicionar `userId` (do JWT sub) no envelope de response como parte do `meta` quando autenticado.

5. Rate limiting básico por API key: max N requests/minuto. Retornar `429 Too Many Requests` com `Retry-After`.

**Entregável por linguagem:**
- .NET: `AddAuthentication().AddJwtBearer()` + middleware de rate limit
- NestJS: `@nestjs/jwt` + guard + `@nestjs/throttler`
- Python: `python-jose` + middleware
- Go: middleware JWT manual ou `golang-jwt`
- Dart: middleware JWT manual

---

### ETAPA 4 — Observabilidade Mínima

**Objetivo:** Saber o que está acontecendo no sistema. Logs estruturados, health checks, métricas básicas, request tracing.

**O que fazer:**

1. **Logs estruturados (JSON):**
   - Todo request logado com: method, path, statusCode, durationMs, traceId, userId
   - Todo erro logado com: stack trace, request body (sanitizado), traceId
   - .NET: Serilog com sink Console (JSON)
   - NestJS: Pino com `pino-pretty` em dev, JSON em prod

2. **TraceId / CorrelationId:**
   - Gerar UUID por request (ou aceitar do header `X-Correlation-Id`)
   - Propagar em todos os logs e no response header
   - Incluir no envelope de erro

3. **Health checks:**
   ```
   GET /health          → { status: "healthy", uptime, version }
   GET /health/ready    → checa SQLite (e depois Redis, PostgreSQL quando entrar)
   GET /health/live     → sempre 200 (liveness probe pra K8s)
   ```

4. **Métricas básicas (salvar no SQLite por enquanto, migrar pra Redis depois):**
   ```
   request_log
   ├── id
   ├── method
   ├── path
   ├── status_code
   ├── duration_ms
   ├── user_id (nullable)
   ├── trace_id
   ├── created_at
   ```

5. **Endpoint de métricas:**
   ```
   GET /metrics → { totalRequests, avgDurationMs, errorRate, requestsByEndpoint, requestsByStatus }
   ```

6. **OpenAPI / Swagger:**
   - .NET: Swashbuckle ou Scalar
   - NestJS: `@nestjs/swagger`
   - Documentar todos os endpoints com schemas, exemplos, e códigos de erro

**Entregável por linguagem:**
- .NET: Serilog + health checks nativos + Swagger
- NestJS: Pino + Terminus health checks + Swagger
- Python: logging estruturado + health endpoint
- Go: middleware de logging + health endpoint
- Dart: middleware de logging + health endpoint

---

### ETAPA 5 — Redis

**Objetivo:** Subir Redis no Docker Compose e começar a usar para cache, rate limiting, e idempotência.

**O que fazer:**

1. Adicionar Redis ao `docker-compose.yml`.

2. Migrar o cache de idempotência (Etapa 2) do SQLite pro Redis com TTL nativo.

3. Migrar rate limiting (Etapa 3) pro Redis (counter com TTL).

4. Cache de leitura: cachear `GET /aircraft-v2` e `GET /aircraft-v2/{id}` no Redis com TTL de 5 min. Invalidar no POST/PUT/DELETE.

5. Health check do Redis no `/health/ready`.

6. Abstrair o acesso ao Redis atrás de interface (`ICacheService`) pra trocar implementação depois.

**Entregável:**
- .NET: `StackExchange.Redis` + `IDistributedCache`
- NestJS: `@nestjs/cache-manager` + `cache-manager-ioredis`

---

### ETAPA 6 — Testes Unitários e de Integração

**Objetivo:** Cobertura mínima que prove que cada peça funciona. Foco em validações, contrato, idempotência, e CRUD.

**O que fazer:**

1. **Testes unitários — validações:**
   - Testar cada regra de validação isolada
   - Testar combinações de campos inválidos
   - Testar que erro retornado segue o contrato

2. **Testes unitários — serialização:**
   - Testar que o JSON de response está em camelCase
   - Testar que enums serializam como string
   - Testar que o envelope de erro está correto

3. **Testes de integração — CRUD completo:**
   - POST cria e retorna 201 com body correto
   - GET retorna o que foi criado
   - PUT atualiza e retorna 200
   - DELETE retorna 204
   - GET após DELETE retorna 404
   - POST com body inválido retorna 400 com detalhes

4. **Testes de integração — idempotência:**
   - POST com mesmo `Idempotency-Key` retorna mesmo resultado
   - POST com key diferente cria novo registro

5. **Testes de integração — autenticação:**
   - Request sem token retorna 401
   - Request com token inválido retorna 401
   - Request com token válido passa

6. **Testes de integração — rate limiting:**
   - Burst de requests excedendo limite retorna 429

**Entregável por linguagem:**
- .NET: xUnit + `WebApplicationFactory` + FluentAssertions
- NestJS: Jest + Supertest + test module
- Python: pytest + httpx/TestClient
- Go: `testing` + `httptest`
- Dart: `test` + `shelf_test_handler` ou equivalente

---

### ETAPA 7 — Refatoração Estrutural (somente .NET e NestJS)

**Objetivo:** Sair do arquivo único e organizar pra escalar. Não é DDD completo, é organização pragmática.

**O que fazer no .NET:**

```
AeroStack.Api/
├── Program.cs (bootstrap limpo)
├── Endpoints/
│   ├── AircraftEndpoints.cs
│   ├── AuthEndpoints.cs
│   ├── HealthEndpoints.cs
│   └── MetricsEndpoints.cs
├── Models/
│   ├── Aircraft.cs
│   ├── Requests/
│   └── Responses/
├── Validators/
│   └── CreateAircraftValidator.cs
├── Services/
│   ├── IAircraftService.cs
│   ├── AircraftService.cs
│   ├── ICacheService.cs
│   └── RedisCacheService.cs
├── Data/
│   ├── IAircraftRepository.cs
│   └── SqliteAircraftRepository.cs
├── Middleware/
│   ├── CorrelationIdMiddleware.cs
│   ├── IdempotencyFilter.cs
│   └── RequestLoggingMiddleware.cs
└── Configuration/
```

**O que fazer no NestJS:**

```
aerostack-api/
├── src/
│   ├── main.ts
│   ├── app.module.ts
│   ├── aircraft/
│   │   ├── aircraft.module.ts
│   │   ├── aircraft.controller.ts
│   │   ├── aircraft.service.ts
│   │   ├── aircraft.repository.ts
│   │   ├── dto/
│   │   └── validators/
│   ├── auth/
│   ├── health/
│   ├── common/
│   │   ├── interceptors/
│   │   ├── guards/
│   │   ├── filters/
│   │   └── dto/
│   └── config/
```

---

### CHECKPOINT — "Backend Premium sem AI"

Neste ponto o AeroStack tem, em pelo menos .NET e NestJS:

- CRUD completo com modelo rico (18+ campos, nested objects, collections)
- Contrato JSON padronizado (camelCase, envelope, erros estruturados)
- Validações completas em todos os campos
- Idempotência em operações de escrita
- Autenticação JWT mínima + rate limiting
- Observabilidade: logs estruturados, traceId, health checks, métricas, OpenAPI
- Redis para cache, rate limit e idempotência
- Testes unitários e de integração
- Código organizado em camadas/módulos

Isso é exatamente o que o roadmap unificado chama de "Fase 0 — Fundações de backend".

**Prioridade por linguagem:**
1. .NET (principal) — tudo acima completo
2. NestJS+Fastify (clone) — tudo acima completo
3. Go, Python, Dart — CRUD + contrato + validações + testes básicos (nivelar "nível médio")

---

## PARTE 2 — AI INTEGRATIONS

A partir daqui, .NET e NestJS caminham juntos. Cada feature é implementada nos dois e deve produzir o mesmo resultado (ou o mais próximo possível).

---

### FASE 1 — Provider Abstraction + Chat Completions

**Objetivo:** Primeira integração real com LLM. Abstrair provider desde o dia 1.

**O que fazer:**

1. Criar interfaces de abstração:
   ```
   IAiChatProvider
   ├── SendMessageAsync(messages, options) → ChatResponse
   ├── SendMessageStreamAsync(messages, options) → IAsyncEnumerable<string>
   ```

2. Implementar para Ollama (local, gratuito) como primeiro provider.

3. Criar endpoint:
   ```
   POST /api/chat/completions
   Body: { messages: [{ role, content }], temperature?, maxTokens? }
   Response: { data: { content, model, usage: { inputTokens, outputTokens } }, meta: { durationMs, provider } }
   ```

4. System prompt fixo do domínio AeroStack: "You are an expert military aviation assistant..."

5. Logging de cada chamada: provider, modelo, tokens in/out, duração, custo estimado.

**Entregável:** Endpoint funcional que conversa via Ollama nos dois backends.

---

### FASE 2 — Embeddings + Dual Write

**Objetivo:** Ao criar um aircraft, gerar embedding da descrição e salvar no vector store.

**O que fazer:**

1. Criar interface `IEmbeddingProvider` com método `GenerateEmbeddingAsync(text) → float[]`.

2. Implementar com Ollama (`nomic-embed-text`).

3. Subir pgvector no Docker Compose (extensão do PostgreSQL) — ou usar SQLite com extensão vetorial para dev.

4. No `POST /api/aircraft-v2`: após salvar no relacional, gerar embedding da concatenação `model + manufacturer + role + description + tags` e salvar no vector store.

5. Tratar falha do embedding como não-bloqueante: salvar no relacional mesmo se embedding falhar, marcar status "embedding_pending" pra reprocessar depois.

**Entregável:** Aircraft criado com embedding salvo. Fallback funcional quando Ollama cai.

---

### FASE 3 — Semantic Search

**Objetivo:** Buscar aircraft por significado, não por texto exato.

**Endpoint:**
```
GET /api/aircraft/search?query={texto}&limit=10
```

**Fluxo:**
1. Gerar embedding da query
2. Buscar N vetores mais próximos (cosine similarity)
3. Recuperar dados completos do relacional pelos IDs
4. Retornar com score de similaridade

**Entregável:** Busca "caça stealth de superioridade aérea" retorna F-22, F-35, Su-57.

---

### FASE 4 — Resiliência nas chamadas AI

**Objetivo:** O sistema não morre quando o LLM está fora.

**O que fazer:**

1. Timeout: máximo 30s para chat, 10s para embeddings.
2. Retry: exponential backoff, max 3 tentativas.
3. Circuit breaker: abre após 3 falhas em 30s, half-open após 60s.
4. Fallback: busca semântica degrada para busca textual (`LIKE`), chat retorna "AI temporarily unavailable".
5. Teste: parar Ollama e validar que tudo degrada gracefully.

**Entregável:**
- .NET: Polly resilience pipeline
- NestJS: Opossum circuit breaker

---

### FASE 5 — RAG Simples

**Objetivo:** Perguntar algo sobre os aircraft e receber resposta grounded nos dados do sistema.

**Endpoint:**
```
POST /api/aircraft/ask
Body: { question: "Qual caça tem maior alcance operacional?" }
Response: { data: { answer, sources: [{ aircraftId, model, score }] }, meta: { ... } }
```

**Fluxo:**
1. Embedding da pergunta
2. Busca vetorial → top 5 chunks/aircraft relevantes
3. Montar prompt: "Based ONLY on the following data: [contexto]. Answer: [pergunta]"
4. Enviar pro LLM
5. Retornar resposta com lista de fontes usadas

**Entregável:** RAG funcional nos dois backends com mesma qualidade de resposta.

---

### FASE 6 — Structured Outputs

**Objetivo:** LLM retorna JSON tipado, não texto livre.

**Endpoint:**
```
POST /api/aircraft/extract-insights
Body: { text: "The F-22 Raptor is a stealth air superiority fighter..." }
Response: { data: { aircraftName, roles: [], keywords: [], confidence } }
```

**O que fazer:**
1. Definir JSON Schema do output esperado.
2. Enviar schema no request ao LLM (OpenAI `response_format` ou prompt engineering para Ollama).
3. Validar o output recebido contra o schema.
4. Fallback: se parsing falhar, retry com prompt mais restritivo.

**Entregável:** Extração tipada funcionando com validação de output.

---

### FASE 7 — Streaming (SSE)

**Objetivo:** Chat com resposta progressiva token a token.

**Endpoint:**
```
POST /api/chat/stream
Headers: Accept: text/event-stream
Body: { messages: [...] }
```

**O que fazer:**
1. .NET: `IAsyncEnumerable<string>` no endpoint, `Content-Type: text/event-stream`
2. NestJS: `@Sse()` decorator ou response manual com Fastify
3. Suportar cancelamento (CancellationToken / abort signal)
4. Timeout do stream total (ex: 120s)

**Entregável:** Chat com streaming funcional, testável via curl/Postman.

---

### FASE 8 — Tool Calling / Function Calling

**Objetivo:** O LLM decide quais funções do backend chamar baseado na pergunta do usuário.

**Endpoint:**
```
POST /api/mission-briefing
Body: { prompt: "Find stealth fighters and check weather at Edwards AFB" }
```

**Tools disponíveis:**
```
search_aircraft(query, role?, status?) → lista de aircraft
get_aircraft_details(id) → aircraft completo
get_weather(location) → mock de clima
search_documents(query) → busca em documentos (quando existir)
```

**Fluxo:**
1. Enviar prompt + definição de tools pro LLM
2. LLM retorna tool_calls com nome e argumentos
3. Backend executa cada tool real
4. Retorna resultados pro LLM
5. LLM formula resposta final consolidada

**Entregável:** Mission briefing funcional com pelo menos 2 tools nos dois backends.

---

### FASE 9 — Document Pipeline (Upload + Ingestão)

**Objetivo:** Upload de PDF/TXT, extração de texto, chunking, embedding, indexação.

**Endpoints:**
```
POST /api/documents/upload         → recebe arquivo, salva, dispara pipeline
GET  /api/documents/{id}/status    → status: uploaded/extracting/chunking/embedding/indexed/failed
POST /api/documents/search         → busca semântica nos documentos
POST /api/documents/reindex        → reprocessa todos
```

**Pipeline:**
```
upload → save to storage → extract text → split chunks (500-800 tokens) → generate embeddings → store in vector DB
```

**O que fazer:**
1. Object storage: pasta local em dev, MinIO no Docker pra simular S3
2. Background processing: `BackgroundService` no .NET, BullMQ no NestJS
3. Status tracking: tabela `document_jobs` com status por etapa
4. Usar manuais de aviões como documentos de teste (3 manuais, ~100 páginas, ~300 chunks)

**Entregável:** Pipeline completo de ingestão com status tracking.

---

### FASE 10 — RAG com Documentos

**Objetivo:** Combinar dados de aircraft + documentos no RAG.

**Atualizar:**
```
POST /api/aircraft/ask  → agora busca em aircraft E documentos
POST /api/chat/ask      → endpoint geral que combina todas as fontes
```

**Entregável:** Perguntas como "What is the maintenance procedure for stealth coating?" respondidas com base nos documentos ingeridos.

---

### FASE 11 — Multimodal (Análise de Imagem)

**Endpoint:**
```
POST /api/aircraft/{id}/analyze-blueprint
Content-Type: multipart/form-data
Body: file (imagem)
```

**O que fazer:**
1. Receber imagem via multipart
2. Enviar para modelo multimodal (LLaVA no Ollama)
3. Prompt: "Identify weapons, engines, and notable features in this aircraft blueprint"
4. Salvar insights como tags automáticas no aircraft

**Entregável:** Upload de imagem retorna análise textual + tags geradas.

---

### FASE 12 — Agentic Workflows

**Endpoints:**
```
POST /api/agents/summarize-manual    → resume manual multi-step
POST /api/agents/recommend-aircraft  → recomenda baseado em objetivo
```

**O que fazer:**
1. Loop de execução: LLM planeja → executa tool → avalia → repete ou finaliza
2. Guardrails: max 5 iterações, token budget por request
3. Logging de cada iteração do agente
4. Confirmação humana para ações destrutivas (opcional)

**Entregável:** Agente funcional com loop controlado e observabilidade por iteração.

---

### FASE 13 — MCP Server

**Objetivo:** Expor as capabilities do AeroStack como MCP Server.

**Tools expostas:**
```
search_aircraft
get_aircraft_details
search_documents
summarize_manual
get_weather (mock)
```

**O que fazer:**
1. Implementar MCP Server em .NET e NestJS
2. Testar com Claude Desktop como client
3. Documentar tools, resources e prompts disponíveis

**Entregável:** MCP Server conectável a qualquer client compatível.

---

### FASE 14 — Evaluation Mínima

**Objetivo:** Medir qualidade antes de considerar "pronto".

**O que medir:**
- Qualidade da resposta RAG (relevância dos chunks recuperados)
- Grounding (resposta baseada no contexto vs. alucinação)
- Latência por endpoint de AI
- Custo estimado por request (tokens × preço do modelo)
- Taxa de falha por provider
- Precisão do tool calling (chamou a tool certa?)

**Entregável:** Dashboard simples ou tabela no README com métricas coletadas.

---

## RESUMO DA ORDEM DE EXECUÇÃO

```
PARTE 1 — Fundações (nivelar backend)
├── Etapa 0: Contrato JSON padronizado
├── Etapa 1: Validações completas
├── Etapa 2: Idempotência
├── Etapa 3: Autenticação mínima + rate limiting
├── Etapa 4: Observabilidade (logs, tracing, health, métricas, OpenAPI)
├── Etapa 5: Redis (cache, rate limit, idempotência)
├── Etapa 6: Testes unitários e de integração
└── Etapa 7: Refatoração estrutural (.NET e NestJS)

PARTE 2 — AI Integrations
├── Fase 1:  Provider abstraction + Chat Completions
├── Fase 2:  Embeddings + Dual Write
├── Fase 3:  Semantic Search
├── Fase 4:  Resiliência nas chamadas AI
├── Fase 5:  RAG simples (aircraft)
├── Fase 6:  Structured Outputs
├── Fase 7:  Streaming (SSE)
├── Fase 8:  Tool Calling / Function Calling
├── Fase 9:  Document Pipeline (upload + ingestão)
├── Fase 10: RAG com documentos
├── Fase 11: Multimodal (análise de imagem)
├── Fase 12: Agentic Workflows
├── Fase 13: MCP Server
└── Fase 14: Evaluation mínima
```

## PRIORIDADE POR LINGUAGEM

```
Tier 1 (feature-complete): .NET + NestJS/Fastify
  → Todas as etapas e fases, resultado idêntico

Tier 2 (nível médio): Go, Python, Dart
  → Parte 1 completa (Etapas 0-6)
  → Parte 2 até Fase 5 (RAG simples)
  → Sem necessidade de paridade total com Tier 1

Tier 3 (quando fizer sentido): Go como worker
  → Document pipeline (Fase 9) em Go pra mostrar concorrência
  → Embedding worker em Go consumindo fila
```

## DIFERENÇAS PLANEJADAS ENTRE .NET E NESTJS

Ambos entregam o mesmo resultado, mas com stacks diferentes:

```
Feature              .NET                          NestJS
─────────────────────────────────────────────────────────────
ORM/DB               EF Core / Dapper + SQLite     Prisma / TypeORM + SQLite
Validação            FluentValidation              class-validator
Resiliência          Polly                         Opossum
Logging              Serilog                       Pino
Background jobs      BackgroundService             BullMQ
Cache                StackExchange.Redis           cache-manager-ioredis
AI SDK               Semantic Kernel (opcional)    LangChain.js (opcional)
Testes               xUnit + WebAppFactory         Jest + Supertest
OpenAPI              Swagger/Scalar                @nestjs/swagger
```

## PARTE 3 — MULTI-AGENT EVOLUTION

A ideia desta parte é fechar o treinamento premium de AI Integration mostrando que o AeroStack consegue evoluir de:

1. um agente com tools
2. para múltiplos agentes dentro da mesma API
3. para orquestração distribuída entre serviços
4. para exposição padronizada via MCP
5. com guardrails, observabilidade e avaliação mínima

A ordem recomendada é sempre:
**in-process first → distributed later**.

---

### FASE 15 — Multi-Agent In-Process (Dentro da Mesma API)

**Objetivo:** Implementar múltiplos agentes especializados rodando dentro do mesmo backend, compartilhando a mesma base, as mesmas tools e a mesma observabilidade.

**Quando usar:**
- Primeiro passo de multi-agent
- Menor complexidade operacional
- Melhor para debug, testes e evolução do lab
- Ideal para .NET e NestJS manterem paridade funcional

**Agentes sugeridos:**
- `RouterAgent` → decide quem deve agir
- `AircraftResearchAgent` → consulta aircrafts e specs
- `DocumentResearchAgent` → consulta chunks e manuais
- `ComparisonAgent` → compara duas ou mais aeronaves
- `QAGuardAgent` → revisa grounding, consistência e formato final

**Endpoint principal:**

```http
POST /api/agents/hello-world
```

**Body:**

```json
{
  "objective": "Summarize the F-22 and list its most relevant stealth-related characteristics."
}
```

**Fluxo mínimo (hello world):**
1. `RouterAgent` recebe o objetivo
2. chama `search_aircraft("F-22")`
3. chama `get_aircraft_details(id)`
4. opcionalmente chama `search_documents("F-22 stealth")`
5. `QAGuardAgent` valida se a resposta final está grounded
6. retorna resposta consolidada

**Resposta esperada:**

```json
{
  "data": {
    "finalAnswer": "The F-22 is a stealth air superiority fighter...",
    "agentsUsed": [
      "RouterAgent",
      "AircraftResearchAgent",
      "QAGuardAgent"
    ],
    "sources": [
      { "type": "aircraft", "id": "..." },
      { "type": "document", "id": "..." }
    ]
  },
  "meta": {
    "iterations": 3,
    "durationMs": 842
  }
}
```

**Endpoints adicionais sugeridos:**

```http
POST /api/agents/research-aircraft
POST /api/agents/compare-aircraft
POST /api/agents/summarize-document
POST /api/agents/qa-check
```

**Entregável:**
- Multi-agent funcional dentro de uma única API
- Logs por iteração
- Lista de tools chamadas
- Resposta final com fontes e agentes utilizados

---

### FASE 16 — Multi-Agent por Especialização de Fluxo

**Objetivo:** Separar o comportamento dos agentes por tipo de trabalho, mas ainda no mesmo deploy.

**Casos de uso ideais:**
- extração de dados de manual
- comparação entre aeronaves
- recomendação orientada a missão
- validação de inconsistências entre cadastro e documento

#### Fluxo A — Catalogação assistida

```http
POST /api/agents/catalog-aircraft-from-document
```

**Body:**

```json
{
  "documentId": "doc_123"
}
```

**Passos:**
1. `DocumentResearchAgent` lê o documento
2. `ExtractionAgent` extrai campos estruturados
3. `CatalogAgent` monta um draft compatível com o DTO
4. `QAGuardAgent` marca campos incertos
5. retorna draft para revisão humana

**Entregável hello world:**
Subir um PDF simples e receber um draft parcial de aircraft.

#### Fluxo B — Comparação orientada a missão

```http
POST /api/agents/recommend-aircraft
```

**Body:**

```json
{
  "mission": "Recommend the best aircraft for long-range stealth interception."
}
```

**Passos:**
1. `RouterAgent` identifica intenção
2. `AircraftResearchAgent` busca candidatos
3. `ComparisonAgent` compara range, speed, stealth e ceiling
4. `QAGuardAgent` remove claims não suportadas
5. retorna ranking final

**Entregável hello world:**
Uma resposta com top 3 aeronaves, justificativa e fontes usadas.

---

### FASE 17 — Multi-Agent Distribuído (Orchestrator + Serviços)

**Objetivo:** Evoluir do multi-agent in-process para uma arquitetura em que existe um **orquestrador** chamando serviços internos separados.

**Importante:**
Aqui, **multi-agent** e **microservices** continuam sendo coisas diferentes:

- **multi-agent** = coordenação cognitiva
- **microservices** = separação arquitetural

Você pode ter:
- um orquestrador agentic chamando serviços determinísticos
- ou vários serviços, cada um com seu próprio agente local

**Arquitetura mínima sugerida:**
- `orchestrator-service`
- `aircraft-service`
- `document-service`
- `evaluation-service`

**Fluxo mínimo (hello world distribuído):**

```http
POST /api/orchestrator/hello-world
```

**Body:**

```json
{
  "objective": "Compare the F-22 and the Su-57 using registered aircraft data and imported manuals."
}
```

**Passos:**
1. Orquestrador recebe o objetivo
2. chama `aircraft-service`
3. chama `document-service`
4. consolida os resultados
5. chama `evaluation-service` para grounding básico
6. retorna resposta final

**Endpoints internos sugeridos:**

```http
GET  /internal/aircraft/search
GET  /internal/aircraft/{id}
POST /internal/documents/ask
POST /internal/documents/search
POST /internal/evals/grounding-check
```

**Entregável:**
- Orquestrador separado
- Pelo menos 2 serviços internos reais
- Contrato JSON idêntico entre serviços
- `traceId` propagado entre chamadas

---

### FASE 18 — Serviço com Agente Próprio

**Objetivo:** Mostrar a versão mais madura, em que cada serviço pode ter um agente local especializado.

**Exemplo de distribuição:**
- `aircraft-service` → `AircraftResearchAgent`
- `document-service` → `DocumentResearchAgent`
- `catalog-service` → `CatalogAgent`
- `evaluation-service` → `QAGuardAgent`

**Vantagem:**
Cada serviço conhece profundamente seu próprio domínio e expõe capacidades mais ricas.

**Cuidado:**
Só fazer isso depois que a versão com orquestrador + serviços determinísticos já estiver estável.

**Hello world mínimo:**

```http
POST /api/orchestrator/research-mission
```

**Fluxo:**
1. Orquestrador chama `document-service`
2. `document-service` usa seu agente local para resumir trechos
3. Orquestrador chama `aircraft-service`
4. `aircraft-service` usa seu agente local para montar comparações
5. Orquestrador consolida a resposta final

**Entregável:**
- Pelo menos 1 serviço com agente local
- Orquestrador central ainda controlando o fluxo
- Logs separados por serviço e por agente

---

### FASE 19 — MCP Mesh (Tools e Resources Externos)

**Objetivo:** Expor as capabilities do AeroStack como uma malha de tools/resources reutilizáveis por qualquer client compatível com MCP.

**Estratégias possíveis:**
1. **1 MCP Server único** expondo tudo
2. **1 MCP Server por domínio**
   - aircraft MCP server
   - documents MCP server
   - evaluations MCP server

**Hello world mínimo:**
Expor somente estas tools:

```text
search_aircraft
get_aircraft_details
search_documents
compare_aircraft
```

**E estes resources:**

```text
aircraft://{id}
document://{id}
manual-chunk://{id}
```

**Prompt reutilizável sugerido:**

```text
compare-aircraft-for-mission
```

**Smoke test mínimo:**
- conectar um client MCP
- listar tools
- chamar `search_aircraft`
- chamar `get_aircraft_details`
- validar retorno estruturado

**Entregável:**
- MCP Server funcional
- 2 a 4 tools úteis
- resources documentados
- prompt reutilizável publicado

---

### FASE 20 — Human-in-the-Loop

**Objetivo:** Treinar o ponto em que AI deixa de ser só automação e passa a ser sistema confiável com revisão humana.

**Casos em que revisão humana é obrigatória:**
- criação automática de aircraft
- alteração de cadastro existente
- classificação de conflito histórico
- resposta com baixa confiança
- respostas sem grounding suficiente

**Endpoints sugeridos:**

```http
POST /api/review/approve-draft
POST /api/review/reject-draft
GET  /api/review/pending
```

**Hello world mínimo:**
1. documento é processado
2. agente gera draft
3. draft vai para fila de revisão
4. humano aprova
5. sistema cria ou atualiza aircraft

**Entregável:**
- workflow com aprovação humana
- fila de drafts pendentes
- decisão auditável

---

### FASE 21 — Evaluation de Multi-Agent

**Objetivo:** Medir se o sistema multi-agent realmente ficou melhor, e não só mais complexo.

**O que medir:**
- taxa de sucesso por workflow
- número médio de iterações por request
- latência total por fluxo
- quais tools são chamadas com mais frequência
- quantas vezes o `QAGuardAgent` corrige a resposta
- grounding score
- custo por execução
- taxa de fallback para fluxo determinístico

**Endpoints sugeridos:**

```http
POST /api/evals/run-multi-agent
GET  /api/evals/multi-agent-report
```

**Dataset mínimo para hello world:**
- 10 perguntas sobre aircraft
- 10 perguntas sobre documentos
- 5 comparações
- 5 tarefas de catalogação
- 5 casos ambíguos para forçar roteamento entre agentes

**Entregável:**
- relatório mínimo com métricas por fluxo
- comparação entre:
  - single-agent
  - multi-agent in-process
  - multi-agent distribuído

---

## RESUMO DA ESCADA DE EVOLUÇÃO

```text
Nível 1 → Single-agent com tools
Nível 2 → Multi-agent dentro da mesma API
Nível 3 → Multi-agent por fluxos especializados
Nível 4 → Orquestrador chamando microserviços
Nível 5 → Serviços com agentes próprios
Nível 6 → MCP mesh
Nível 7 → Human-in-the-loop + evaluation comparativa
```

---

## ORDEM RECOMENDADA DE IMPLEMENTAÇÃO

```text
1. Fase 15 — Multi-Agent In-Process
2. Fase 16 — Multi-Agent por Especialização de Fluxo
3. Fase 20 — Human-in-the-Loop
4. Fase 21 — Evaluation de Multi-Agent
5. Fase 17 — Multi-Agent Distribuído
6. Fase 18 — Serviço com Agente Próprio
7. Fase 19 — MCP Mesh
```

---

## CHECKPOINT — "AI Integration Premium com Multi-Agent"

Neste ponto o AeroStack deve ser capaz de:

- executar múltiplos agentes dentro da mesma API
- rotear tarefas para agentes especializados
- expor tools e resources via MCP
- evoluir de monólito agentic para orquestração distribuída
- usar revisão humana nos fluxos críticos
- medir qualidade, grounding, custo, latência e taxa de sucesso
- manter paridade funcional entre **.NET** e **NestJS/Fastify**

