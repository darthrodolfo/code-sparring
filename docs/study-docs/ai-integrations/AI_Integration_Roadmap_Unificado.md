# AI Integrations para Backend Engineers — Roadmap Unificado (.NET + Fastify)

## Sumário

1. [Objetivo do documento](#objetivo-do-documento)
2. [Visão estratégica](#visão-estratégica)
3. [Os 2 showcases](#os-2-showcases)
4. [Arquitetura-alvo por etapas](#arquitetura-alvo-por-etapas)
5. [Stack, plugs, bancos e componentes](#stack-plugs-bancos-e-componentes)
6. [Resumo executivo de endpoints e features](#resumo-executivo-de-endpoints-e-features)
7. [Roadmap prático: do hello world ao enterprise](#roadmap-prático-do-hello-world-ao-enterprise)
8. [AI integrations que precisam entrar no laboratório](#ai-integrations-que-precisam-entrar-no-laboratório)
9. [Boas práticas de arquitetura e operação](#boas-práticas-de-arquitetura-e-operação)
10. [C#/.NET: abordagem recomendada](#cnet-abordagem-recomendada)
11. [TypeScript/Fastify: abordagem recomendada](#typescriptfastify-abordagem-recomendada)
12. [Quando entrar com Kafka, mensageria e microserviços](#quando-entrar-com-kafka-mensageria-e-microserviços)
13. [Cloud: como treinar Azure, AWS e GCP sem se perder](#cloud-como-treinar-azure-aws-e-gcp-sem-se-perder)
14. [Timeline e fundamentos conceituais de AI Integration](#timeline-e-fundamentos-conceituais-de-ai-integration)
15. [Como explicar isso em entrevista](#como-explicar-isso-em-entrevista)
16. [Plano de execução sugerido](#plano-de-execução-sugerido)

---

## Objetivo do documento

Este documento unifica:

- o aulão teórico sobre AI Integrations;
- o roadmap prático para os clones em **.NET** e **Fastify (Node.js/TypeScript)**;
- as decisões que surgiram nas conversas sobre **monólito lean**, **versão enterprise**, **cloud**, **mensageria**, **RAG**, **agents**, **MCP**, **document pipelines** e **showcase de portfólio**;
- um resumo claro de **features**, **endpoints**, **bancos**, **plugs**, **integrações**, **boas práticas**, e a **ordem correta de estudo e implementação**.

A meta não é só “aprender AI”.
A meta é **completar o backend padrão ouro** e encaixar AI como capacidade real de produto e de arquitetura.

---

## Visão estratégica

A ideia central do laboratório é esta:

1. **Primeiro construir uma base premium de backend**, sem se preocupar com tela.
2. **Depois adicionar AI de forma incremental**, começando com o mínimo útil.
3. **Depois salvar uma versão monolítica muito boa como showcase lean**.
4. **Depois evoluir para uma versão mais enterprise**, extraindo serviços, adicionando mensageria, observabilidade distribuída, cloud resources e, quando fizer sentido, misturando stacks.

### Filosofia de aprendizado

A estratégia correta aqui é:

- **Passada 1 — breadth first**  
  Fazer o hello world / MVP mínimo de cada tópico importante de AI Integration.
- **Passada 2 — depth second**  
  Voltar tópico por tópico e expandir escopo, qualidade, resiliência e observabilidade.
- **Passada 3 — enterprise refinement**  
  Transformar a prova de conceito em arquitetura séria.

### O que isso evita

Você evita dois erros clássicos:

- aprofundar demais em um único tópico e travar o progresso;
- ver tudo superficialmente e nunca consolidar nada.

Aqui o fluxo é:

**construir o mapa → provar cada peça → depois sofisticar**.

---

## Os 2 showcases

### Showcase 1 — Lean / Monólito premium

Esse é o projeto que prova:

- modelagem de domínio;
- API bem desenhada;
- CRUD sólido;
- persistência relacional;
- upload e ingestão de documentos;
- busca semântica;
- RAG simples;
- chat com contexto;
- structured outputs;
- tool calling básico;
- resiliência;
- observabilidade;
- background jobs;
- Docker e ambiente local reproduzível.

Essa versão é ideal para mostrar:

- pragmatismo;
- coesão;
- capacidade de entrega;
- organização arquitetural;
- domínio de backend moderno sem overengineering precoce.

### Showcase 2 — Enterprise / Distribuído

Esse é o projeto que prova:

- extração gradual de responsabilidades;
- microserviços com propósito;
- mensageria;
- processamento assíncrono;
- pipelines de documentos;
- serviços poliglotas;
- cloud readiness;
- tracing distribuído;
- event-driven architecture;
- escalabilidade operacional;
- decisão arquitetural madura.

A narrativa correta não é “microserviço é melhor”.
A narrativa correta é:

> Primeiro estabilizamos domínio, contratos e fluxo no monólito.  
> Depois extraímos componentes que realmente ganharam com autonomia operacional, throughput próprio, isolamento de falhas e processamento assíncrono.

---

## Arquitetura-alvo por etapas

### Etapa A — Local, barata, simples e poderosa

Rodar tudo localmente com Docker:

- API .NET
- API Fastify
- PostgreSQL
- Redis
- armazenamento de arquivos local/MinIO
- Ollama
- Qdrant ou pgvector
- worker de background
- observabilidade local
- opcionalmente Redpanda/Kafka local na fase enterprise

Essa etapa serve para aprender sem custo e sem fricção de cloud.

### Etapa B — Monólito organizado

Mesmo antes de quebrar em serviços, organizar o código por camadas ou módulos:

- `Domain`
- `Application`
- `Infrastructure`
- `Api`
- `Integrations`

Ou por módulos de negócio:

- `Aircraft`
- `Documents`
- `Search`
- `Chat`
- `MissionBriefing`
- `Notifications`
- `Audit`

### Etapa C — Monólito com AI real

Adicionar:

- provider abstraction;
- embeddings;
- semantic search;
- RAG;
- upload de PDF/imagem;
- structured output;
- tool calling;
- streaming;
- logs e métricas;
- retry, timeout, circuit breaker;
- tarefas assíncronas.

### Etapa D — Extração gradual

Extrair apenas o que realmente merece vida própria, por exemplo:

- `document-ingestion-service`
- `embedding-service`
- `search-service`
- `chat-orchestrator-service`
- `notification-service`

### Etapa E — Enterprise distribuído

Adicionar:

- Kafka/Redpanda
- outbox/inbox
- retries
- DLQ
- idempotência
- tracing distribuído
- storage desacoplado
- orquestração de workflows
- cloud resources
- deploy sob demanda

---

## Stack, plugs, bancos e componentes

## Stack principal

### Backend 1 — .NET
- C#
- ASP.NET Core Web API / Minimal APIs
- Entity Framework Core
- Semantic Kernel opcional
- Microsoft.Extensions.AI opcional
- Polly para resiliência
- BackgroundService / Hangfire / Quartz dependendo do estágio

### Backend 2 — TypeScript
- Node.js
- TypeScript
- Fastify (pure — plugin-based architecture)
- Drizzle ou better-sqlite3
- LangChain.js opcional
- Opossum para resiliência (circuit breaker)
- BullMQ ou workers dedicados quando necessário
- JSON Schema + Ajv para validação nativa

### Backend 3 — opcional no estágio enterprise
- Go
- ideal para processamento concorrente de documentos, pipelines e workers

---

## Bancos e storages

### Banco relacional principal
**PostgreSQL**

Uso:
- CRUD de aircraft;
- status de documentos;
- metadata;
- histórico;
- auditoria;
- catálogos;
- perfis;
- jobs e estados do sistema.

### Vetorial
Você pode seguir dois caminhos:

#### Opção A — pgvector
Vantagens:
- tudo no PostgreSQL;
- menos moving parts;
- ótimo para showcase lean;
- excelente para aprender RAG sem inventar mais infraestrutura.

#### Opção B — Qdrant
Vantagens:
- experiência dedicada de vector DB;
- ótimo para comparação de arquitetura;
- interessante para mostrar ecossistema moderno.

### Cache
**Redis**

Uso:
- cache de consultas;
- cache de prompts/respostas;
- deduplicação;
- rate limiting;
- locks leves;
- status transitório;
- filas simples em alguns cenários.

### Object storage
- local: **MinIO**
- cloud: **S3 / Azure Blob / Google Cloud Storage**

Uso:
- PDFs;
- manuais;
- imagens;
- blueprints;
- anexos;
- arquivos de ingestão.

---

## Plugs e providers de AI

### Providers prioritários
- **OpenAI API**
- **Anthropic Claude API**
- **Ollama** para desenvolvimento local e testes baratos

### Providers/Plugs futuros
- Azure OpenAI
- AWS Bedrock
- Vertex AI
- MCP Servers
- web search / file search / tools internas
- weather mock
- OCR
- multimodal vision

### Regra arquitetural
Nunca acoplar o sistema diretamente a um provider só.

Criar uma abstração, por exemplo:

- `IAiProvider`
- `IEmbeddingProvider`
- `IChatProvider`
- `IToolExecutionService`
- `IDocumentExtractionService`

Assim você consegue:

- trocar OpenAI por Anthropic;
- comparar custo e qualidade;
- usar Ollama local;
- plugar Bedrock/Azure OpenAI depois;
- implementar fallback controlado.

---

## Componentes auxiliares

- Docker / Docker Compose
- Kubernetes local: `kind`, `k3d` ou `minikube`
- OpenTelemetry
- Prometheus/Grafana opcional
- Serilog / Pino
- Swagger / OpenAPI
- Postman / Bruno
- test containers
- CI/CD simples
- scripts de seed
- coleções de requests

---

## Resumo executivo de endpoints e features

Abaixo está o mapa de endpoints sugeridos.  
Nem todos precisam nascer no dia 1. A ideia é crescer por estágios.

### 1. CRUD base de aircraft

#### `POST /api/aircraft`
Cria um aircraft.

Pode incluir:
- nome;
- fabricante;
- país;
- tipo;
- ano;
- descrição;
- especificações;
- tags;
- links;
- metadados diversos.

Na fase com AI:
- gerar embedding da descrição;
- persistir no store vetorial;
- opcionalmente gerar tags automáticas.

#### `GET /api/aircraft/{id}`
Retorna dados completos do aircraft.

#### `GET /api/aircraft`
Lista/pagina aircrafts.

#### `PUT /api/aircraft/{id}`
Atualiza o aircraft.

#### `DELETE /api/aircraft/{id}`
Remove ou arquiva.

---

### 2. Busca semântica

#### `GET /api/aircraft/search?query={texto}`
Busca por significado, não só por termo exato.

Exemplos:
- “caça stealth de superioridade aérea”
- “avião de reconhecimento com grande alcance”
- “interceptadores invisíveis ao radar”

Fluxo:
1. gerar embedding da query;
2. buscar vetores mais próximos;
3. trazer IDs relevantes;
4. buscar dados completos no relacional;
5. devolver ranking.

---

### 3. RAG / perguntas sobre dados existentes

#### `POST /api/aircraft/ask`
Payload:
```json
{
  "question": "Comparando os dados do F-35 e do Su-57, qual tem maior alcance operacional?"
}
```

Fluxo:
1. recuperar registros e chunks relevantes;
2. montar prompt com contexto controlado;
3. pedir resposta baseada apenas no contexto;
4. retornar resposta com eventual lista de fontes/IDs.

---

### 4. Documentos e ingestão

#### `POST /api/documents/upload`
Faz upload de PDF, DOC, TXT, imagem ou blueprint.

Ação:
- salva no object storage;
- registra metadata no relacional;
- dispara pipeline de ingestão.

#### `POST /api/documents/{id}/index`
Indexa explicitamente um documento.

#### `GET /api/documents/{id}`
Retorna metadata e status.

#### `GET /api/documents/{id}/status`
Retorna status do pipeline:
- uploaded
- queued
- extracting
- chunking
- embedding
- indexed
- failed

#### `POST /api/documents/search`
Busca semântica em documentos.

Payload:
```json
{
  "query": "stealth coating maintenance procedure"
}
```

---

### 5. Chat geral / assistente do domínio

#### `POST /api/chat/ask`
Pergunta geral ao sistema.

Pode combinar:
- contexto do usuário;
- contexto de aircrafts;
- documentos;
- tools;
- histórico de conversa.

#### `POST /api/chat/stream`
Versão com streaming.

Ideal para:
- UX estilo chat;
- respostas longas;
- sensação de responsividade;
- treinamento de SSE.

---

### 6. Structured output / extração tipada

#### `POST /api/aircraft/extract-insights`
Recebe um texto bruto e devolve estrutura tipada.

Exemplo de saída:
```json
{
  "aircraftName": "F-22 Raptor",
  "roles": ["air superiority", "interceptor"],
  "keywords": ["stealth", "supercruise"],
  "confidence": 0.94
}
```

Serve para treinar:
- JSON Schema;
- DTOs;
- deserialização segura;
- validação de output.

---

### 7. Multimodalidade

#### `POST /api/aircraft/{id}/analyze-blueprint`
Upload de imagem/planta/foto.

O backend:
- lê o arquivo em memória;
- envia para modelo multimodal;
- pede insights;
- salva tags, notas ou observações.

---

### 8. Tool calling / function calling

#### `POST /api/mission-briefing`
Recebe uma tarefa complexa.

Exemplo:
> “Planeje uma missão de reconhecimento e verifique o clima na base.”

O modelo pode decidir chamar:
- busca de aircrafts;
- busca vetorial;
- status de documentos;
- clima mock;
- regras do domínio.

O backend executa as tools reais e devolve o resultado consolidado.

---

### 9. Agents

#### `POST /api/agents/summarize-manual`
Cria um fluxo multi-step para:
- ler documento;
- extrair partes relevantes;
- consultar dados relacionados;
- montar resumo final.

#### `POST /api/agents/recommend-aircraft`
Fluxo mais autônomo:
- entender objetivo;
- buscar aircrafts;
- comparar specs;
- consultar regras;
- devolver recomendação.

---

### 10. Observabilidade e operação

#### `GET /health`
Health check básico.

#### `GET /health/dependencies`
Valida dependências:
- banco;
- redis;
- vector store;
- object storage;
- provider de AI.

#### `GET /metrics`
Métricas para scraping ou exposição interna.

#### `GET /api/jobs/{id}`
Status de job assíncrono.

---

## Roadmap prático: do hello world ao enterprise

## Fase 0 — Fundações de backend

Antes de AI pesada, fechar:

- CRUD sólido;
- persistência;
- migrations;
- validações;
- logging;
- Docker;
- cache;
- background jobs;
- arquitetura interna limpa;
- testes básicos;
- health checks.

Objetivo:
ter um backend premium mesmo sem AI.

---

## Fase 1 — Embeddings e dual write

Primeiro contato real com AI Integration aplicada ao domínio.

### Hello world
Ao criar um aircraft:
- gerar embedding da descrição;
- salvar o registro no banco relacional;
- salvar vetor no vector store.

### O que aprender
- provider de embeddings;
- latência externa;
- dual write;
- idempotência;
- tratamento de falha;
- fallback inicial.

---

## Fase 2 — Semantic search

### Hello world
Buscar “caça stealth” e recuperar aircrafts semanticamente relevantes.

### O que aprender
- similarity search;
- cosine similarity;
- ranking;
- paginação;
- sincronização relacional + vetorial.

---

## Fase 3 — Resiliência e fallbacks

### Hello world
Parar o provider de embeddings/chat e garantir que o sistema não colapse.

### O que aprender
- timeout;
- retry;
- circuit breaker;
- fallback;
- graceful degradation.

Exemplos de fallback:
- busca textual tradicional;
- fila para reprocessar embedding depois;
- resposta parcial;
- status “temporariamente indisponível”.

---

## Fase 4 — RAG simples

### Hello world
Perguntar algo sobre aircrafts ou documentos e responder com contexto recuperado do próprio sistema.

### O que aprender
- chunking;
- retrieval;
- augmentation;
- grounding;
- limite de contexto;
- prompt assembly.

---

## Fase 5 — Structured outputs

### Hello world
Extrair dados em um DTO garantido por schema.

### O que aprender
- JSON Schema;
- tipagem segura;
- reduzir parsing frágil;
- output validado;
- diferenças entre texto livre e resposta estruturada.

---

## Fase 6 — Tool calling

### Hello world
O usuário faz uma pergunta natural e o modelo escolhe a função correta.

Exemplo:
- buscar aircraft;
- consultar clima mock;
- ler documento;
- consultar missão.

### O que aprender
- definição de tools;
- contratos;
- validação de argumentos;
- execução controlada;
- orchestration.

---

## Fase 7 — Streaming

### Hello world
Chat com SSE.

### O que aprender
- streaming token a token;
- cancelamento;
- timeouts;
- UX melhor sem frontend complexo.

---

## Fase 8 — Document pipeline completo

### Hello world
Upload de PDF:
- salvar;
- extrair;
- chunkar;
- embedar;
- indexar;
- consultar.

### O que aprender
- object storage;
- workers;
- pipelines;
- status tracking;
- jobs assíncronos;
- reprocessamento.

---

## Fase 9 — Multimodal

### Hello world
Receber imagem de blueprint e pedir análise simples.

### O que aprender
- upload;
- buffer/stream;
- modelos multimodais;
- persistência de insights.

---

## Fase 10 — Agentic workflows

### Hello world
Fluxo em loop com ferramentas múltiplas e limite de iterações.

### O que aprender
- planner/executor;
- iteração máxima;
- token budget;
- guardrails;
- confirmação humana para ações críticas.

---

## Fase 11 — MCP

### Hello world
Criar um MCP Server expondo algumas tools do domínio.

Exemplos:
- `search_aircraft`
- `get_aircraft_details`
- `search_documents`
- `summarize_manual`

### O que aprender
- tools;
- resources;
- prompts;
- conectividade padronizada com clientes compatíveis.

---

## AI integrations que precisam entrar no laboratório

Abaixo está o checklist do que deve existir pelo menos em forma de MVP.

## 1. Chat completions / responses
- system prompt;
- user prompt;
- histórico;
- controle de temperatura;
- limites de tokens;
- mensagens bem definidas.

## 2. Embeddings
- geração de vetores;
- armazenamento;
- busca semântica;
- comparação entre providers.

## 3. RAG
- recuperação;
- chunking;
- ranking;
- prompt com contexto;
- grounded answer.

## 4. Structured outputs
- schema tipado;
- parsing confiável;
- DTOs fortes.

## 5. Tool calling
- tools internas do domínio;
- tool de clima mock;
- tool de consulta documental;
- tool de busca vetorial.

## 6. Agents
- loop multi-step;
- decisão autônoma;
- orçamento de tokens;
- limite de iterações.

## 7. Streaming
- SSE;
- resposta progressiva;
- cancelamento.

## 8. Document ingestion
- upload;
- extração;
- chunking;
- indexing;
- reprocessamento.

## 9. Multimodal
- análise de imagem/planta;
- enriquecimento com tags.

## 10. MCP
- server do domínio;
- tools;
- resources;
- prompts.

## 11. Evaluation mínima
Mesmo que simples, você deve medir:
- qualidade da resposta;
- grounding;
- latência;
- custo;
- taxa de falha.

---

## Boas práticas de arquitetura e operação

## 1. Provider abstraction obrigatória

Nunca deixar OpenAI/Anthropic/Ollama espalhados pelo código.

Criar interfaces.

Exemplo conceitual:
- `IAiProvider`
- `IEmbeddingService`
- `IChatCompletionService`
- `IToolRouter`
- `IDocumentRagService`

---

## 2. Prompts versionados

Não esconder prompt importante dentro de controller.

Ter:
- arquivos;
- templates;
- classes;
- ou banco/versionamento simples.

Guardar:
- propósito;
- versão;
- parâmetros;
- contexto esperado;
- schema de saída.

---

## 3. Resiliência desde cedo

Aplicar:
- timeout;
- retry com backoff;
- circuit breaker;
- fallback;
- cancellation token;
- limites claros.

---

## 4. Custo e latência como cidadania de primeira classe

Registrar por request:
- provider;
- modelo;
- input tokens;
- output tokens;
- custo estimado;
- duração;
- cache hit/miss;
- rota;
- tenant ou usuário.

---

## 5. Segurança

Nunca:
- colocar secret no prompt;
- mandar PII desnecessária;
- concatenar SQL gerado pelo modelo;
- deixar tool executar ação sem validação.

Sempre:
- validar entrada;
- validar argumentos das tools;
- validar output estruturado;
- usar segredo em vault/env seguro;
- tratar prompt injection como risco real.

---

## 6. Idempotência e pipelines

Para indexação e documentos:
- usar correlation id;
- usar job id;
- permitir reprocessamento;
- impedir duplicação indevida;
- guardar checkpoints de status.

---

## 7. Observabilidade

Você quer enxergar:
- qual endpoint chamou AI;
- qual provider respondeu;
- quanto custou;
- qual tool foi invocada;
- quantas iterações o agent fez;
- quais chunks foram usados no RAG;
- onde falhou.

---

## 8. Testes

Ter pelo menos:
- unit tests de orquestração;
- integration tests com banco;
- testes de fallback;
- testes de schema output;
- smoke tests de pipeline;
- mocks/fakes de provider.

---

## 9. Não exagerar cedo

Nem toda feature precisa de:
- Kafka;
- agent;
- vector DB dedicado;
- microserviço;
- K8s gerenciado.

A ordem madura é:
1. fazer funcionar;
2. organizar;
3. medir dor;
4. extrair;
5. sofisticar.

---

## C#/.NET: abordagem recomendada

## Estrutura sugerida

- `Api`
- `Application`
- `Domain`
- `Infrastructure`
- `Ai`
- `Documents`
- `Search`
- `Workers`

## Bibliotecas e peças úteis

- ASP.NET Core
- EF Core
- Npgsql
- `pgvector-dotnet` se usar pgvector
- Qdrant client se usar Qdrant
- Polly
- OpenTelemetry
- Serilog
- FluentValidation
- MediatR opcional
- Semantic Kernel opcional
- Microsoft.Extensions.AI opcional

## Padrões que fazem sentido em .NET

- `HttpClientFactory`
- `Options pattern`
- `IHostedService` / `BackgroundService`
- Resilience pipelines
- endpoints finos + services/coordinators
- DTOs bem definidos
- health checks por dependência

## Coisas boas para treinar em .NET

- Minimal APIs e Controllers
- SSE com `IAsyncEnumerable`
- providers tipados
- serialização segura
- workers
- DI limpa
- testes com WebApplicationFactory
- pipeline assíncrono de documentos

---

## TypeScript/Fastify: abordagem recomendada

## Estrutura sugerida

- `plugins/` (database, auth, redis, swagger, ai-provider)
- `routes/aircraft/`
- `routes/documents/`
- `routes/search/`
- `routes/chat/`
- `routes/agents/`
- `routes/health/`
- `hooks/` (correlationId, requestLogging, idempotency)
- `services/` (ai, storage, embedding)

## Bibliotecas e peças úteis

- Fastify (core + AutoLoad)
- Drizzle ou better-sqlite3
- `pg` / PostgreSQL
- LangChain.js opcional
- JSON Schema + Ajv (validação nativa do Fastify)
- Pino (built-in no Fastify)
- Opossum
- BullMQ
- KafkaJS quando entrar em eventos
- OpenTelemetry

## Coisas boas para explorar no Fastify

- plugin system com encapsulação (`fp()` vs scoped)
- hooks (onRequest, preHandler, onSend, onError)
- JSON Schema validation + serialization
- `fastify.decorate()` para DI via instância
- SSE/streaming via `reply.raw`
- workers
- filas
- adapters Fastify
- DI e organização modular

---

## Quando entrar com Kafka, mensageria e microserviços

Kafka não entra porque “é bonito”.
Kafka entra quando o pipeline realmente ganha com eventos, retenção e consumo por múltiplos serviços.

## Onde Kafka faz muito sentido

### Pipeline de documentos
Exemplo de fluxo:

1. `document.uploaded`
2. `document.extracted`
3. `document.chunked`
4. `document.embedded`
5. `document.indexed`
6. `document.failed`

A partir disso:
- um worker extrai texto;
- outro chunka;
- outro gera embeddings;
- outro indexa;
- outro notifica.

### Integração entre stacks
- .NET publica evento
- Go consome para processamento pesado
- Fastify atualiza status
- outro serviço envia notificação

### Reprocessamento
Kafka também é ótimo para:
- replay;
- reconstrução de projeções;
- reindexação;
- troubleshooting.

## Onde Kafka é exagero cedo
- CRUD simples;
- autenticação básica;
- operações síncronas pequenas;
- estágio inicial do laboratório.

## Regra madura
- showcase lean: fila simples ou background jobs;
- showcase enterprise: Kafka/Redpanda nos fluxos que realmente merecem.

---

## Cloud: como treinar Azure, AWS e GCP sem se perder

A estratégia correta é:

### Modo 1 — local padrão
Você desenvolve e valida tudo localmente.

### Modo 2 — demo cloud
Você sobe quando quiser demonstrar e desliga depois.

Isso é ótimo para custo baixo.

## Mapa de conceitos portáveis

### Compute / APIs
- Azure: Container Apps / Functions
- AWS: ECS/Fargate / Lambda
- GCP: Cloud Run / Cloud Run functions

### Mensageria
- Azure: Service Bus
- AWS: SQS, SNS, EventBridge
- GCP: Pub/Sub

### Object storage
- Azure: Blob Storage
- AWS: S3
- GCP: Cloud Storage

### Containers / Orquestração
- Azure: Container Apps / AKS
- AWS: ECS / EKS
- GCP: Cloud Run / GKE

### Workflow / orquestração
- Azure: Functions + Durable patterns / Dapr
- AWS: Step Functions
- GCP: Workflows

## O que treinar primeiro
Primeiro dominar o conceito.
Depois aprender o nome da cloud.

Exemplo:
- fila → Service Bus / SQS / Pub/Sub
- storage → Blob / S3 / Cloud Storage
- serverless → Functions / Lambda / Cloud Run functions
- container service → Container Apps / ECS/Fargate / Cloud Run

---

## Timeline e fundamentos conceituais de AI Integration

## 2020–2021 — GPT-3 API
POST com texto, recebe texto.
Integração mais crua, output menos estruturado.

## 2022 — ChatGPT
Explosão de awareness e adoção pública.

## 2023 — Chat Completions
Mensagens com roles:
- system
- user
- assistant

Essa mudança é a base de tudo.

## 2023 — Function Calling / Tool Use
O modelo deixa de ser só uma caixa de texto e passa a atuar como decisor de chamada de ferramentas.

Ponto central:
**o modelo não executa nada sozinho**.  
Seu backend define as tools, valida os argumentos e executa a ação real.

## 2023 — RAG
Padrão arquitetural para responder usando contexto recuperado do seu próprio sistema.

Pipeline clássico:
1. chunking
2. embeddings
3. vector storage
4. retrieval
5. augmented prompt
6. grounded response

## 2024 — Embeddings mais baratos e melhores
Embeddings viram commodity operacional.
O custo maior geralmente fica no modelo gerador, não no embedding.

## 2024 — Structured outputs
Respostas aderindo a schema.
Isso aproxima AI de um mundo muito mais tipado e confiável para backend.

## 2024 — Agents
Tool calling em loop com autonomia controlada.

## 2024 — MCP
Protocolo aberto para conectar aplicações AI a ferramentas, recursos e prompts de forma padronizada.

## 2025–2026 — Consolidação
O stack padrão de AI Integration fica cada vez mais claro:

- chat/responses
- tools
- RAG
- structured output
- agents
- MCP
- multimodalidade
- observabilidade e custo

---

## Conceitos que você precisa dominar

### Tokens
Impactam:
- custo;
- latência;
- limite de contexto.

### Temperature
Use baixo para cenários enterprise e previsíveis.

### Streaming
Melhora UX e te obriga a entender resposta progressiva.

### Provider abstraction
Essencial para liberdade arquitetural.

### RAG
Base de document search útil.

### Tool calling
Base de orquestração com dados reais.

### Agents
Layer acima de tools para tarefas multi-step.

### MCP
Padrão moderno para conectar AI a capabilities externas.

---

## Como explicar isso em entrevista

Estrutura boa de resposta:

1. Clarificar a feature de AI necessária.
2. Identificar se o fluxo é síncrono ou assíncrono.
3. Definir abstração de provider.
4. Escolher o padrão adequado:
   - chat;
   - structured output;
   - RAG;
   - tool calling;
   - agent;
   - MCP.
5. Definir tratamento de falhas.
6. Definir custo e modelo.
7. Definir observabilidade.
8. Definir rollout mínimo e evolução.

Exemplo de fala madura:

> Eu começaria clarificando a capacidade de AI específica que o serviço precisa.  
> Se for chat simples, uso um provider abstraction com prompts controlados e streaming.  
> Se precisar responder com base em dados do domínio, avalio tool calling ou RAG dependendo da origem da verdade.  
> Para documentos, monto pipeline de ingestão, chunking, embeddings e retrieval.  
> Para tarefas multi-step, introduzo um loop agentic com limites de iteração e budget.  
> Tudo isso com timeout, retries, circuit breaker, observabilidade de custo/tokens e fallback para manter a plataforma funcional mesmo quando a AI falha.

---

## Plano de execução sugerido

## Sprint 1 — fundação premium
- CRUD
- banco
- Docker
- Redis
- object storage local
- health checks
- logs
- background jobs

## Sprint 2 — primeiros MVPs de AI
- embeddings
- semantic search
- RAG simples
- structured output
- streaming

## Sprint 3 — capabilities mais modernas
- tool calling
- document pipeline completo
- multimodal simples
- avaliação mínima

## Sprint 4 — showcase lean fechado
- README forte
- diagramas
- exemplos de requests
- métricas/custo
- demonstração end-to-end

## Sprint 5 — enterprise evolution
- extração de serviços
- mensageria
- Kafka/Redpanda
- tracing distribuído
- cloud demo mode
- eventual serviço em Go para pipeline pesado

---

## Conclusão

A grande sacada deste laboratório é que ele finalmente faz todas as peças conversarem entre si:

- backend clássico;
- APIs;
- persistência;
- filas;
- cache;
- cloud;
- observabilidade;
- microserviços;
- AI integrations.

