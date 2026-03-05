# Roadmap: Do Zero ao Avançado em Qualquer Linguagem

> Estrutura genérica e progressiva para aprender qualquer linguagem de programação com profundidade real.  
> Adaptável para: Go, Python, TypeScript, Rust, Kotlin, Java, ou qualquer outra.

---

## Estrutura de Repositórios

```
/linguagem
  /01-fundamentos
  /02-linguagem-core
  /03-conectando-mundo
  /04-distribuido
  /05-agentes
```

Cada pasta com um `README.md` curto explicando:
- O que foi estudado
- Por que foi estudado
- O que aprendeu
- Pacotes e ferramentas utilizados

---

## Fase 1 — Fundamentos da Linguagem

> O que sustenta tudo. Não é básico — é a base que não pode ter furo.

- Hello World
- Tipos primitivos e compostos
- Variáveis, constantes e escopo
- Condicionais e loops
- Funções
- Tratamento de erro *(como ESSA linguagem trata)*
- Coleções — array, slice, map, lista
- Structs / classes / records
- Interfaces / traits / protocolos
- Generics

---

## Fase 2 — A Linguagem Sendo Ela Mesma

> O que essa linguagem faz diferente de todas as outras. Aqui está a identidade dela.

- O diferencial da linguagem
  - **Rust** → ownership, borrow checker, lifetimes
  - **Go** → goroutines, channels, simplicidade forçada
  - **Python** → duck typing, dynamic, ecosystem
  - **TypeScript** → tipagem estática sobre JS, structurals
- Concorrência e paralelismo
- Gerenciamento de memória
- Módulos e organização de projeto
- Testes unitários e de integração
- Manipulação de string, arquivo e JSON

---

## Fase 3 — Conectando com o Mundo

> A linguagem saindo do isolamento e interagindo com sistemas reais.

- HTTP Client — consumir API externa
- HTTP Server — expor API própria (minimal API)
- Banco de dados — leitura e escrita (SQL e/ou NoSQL)
- Variáveis de ambiente e configuração
- Logging e observabilidade básica
- Docker — containerizar a aplicação
- Deploy simples — subir em algum lugar real

---

## Fase 4 — Distribuído e Profissional

> Como a linguagem se comporta em sistemas reais de produção.

- Autenticação e segurança
- Mensageria — RabbitMQ, Kafka ou Redis
- gRPC
- Cache
- Testes de carga e performance
- CI/CD básico

---

## Fase 5 — AI e Agentes

> O diferencial do momento. A linguagem interagindo com inteligência artificial.

### 5.1 — Consumindo LLM
- Integrar API de LLM (OpenAI, Anthropic, Gemini)
- Enviar prompt, receber resposta
- Gerenciar contexto e histórico de conversa

### 5.2 — LLM Local
- Rodar LLM local via Ollama
- Integrar com modelo offline
- Comparar performance local vs cloud

### 5.3 — RAG (Retrieval Augmented Generation)
- Conceito de Embedding — transformar texto em vetor
- Vector Database — ChromaDB, Pinecone, pgvector
- Similarity Search — busca por significado, não por palavra
- Chunking — dividir documentos grandes
- Pipeline completo: documento → embedding → busca → resposta

### 5.4 — Agente Simples
- Tool calling — plugar funções no LLM
- Loop de decisão — o LLM decide qual tool chamar
- Critério de parada — quando o agente conclui
- Memória de contexto entre iterações

### 5.5 — Multi-Agente
- Orquestrador sem LLM — fluxo fixo entre agentes
- Orquestrador com LLM — fluxo dinâmico
- Handoff entre agentes
- Agentes especializados com responsabilidade única

### 5.6 — MCP (Model Context Protocol)
- Criar servidor MCP na linguagem
- Expor tools via protocolo padrão
- Conectar com Claude, Cursor, ou qualquer cliente MCP

---

## Projeto Âncora por Fase

> Um projeto simples e consistente que cresce junto com o aprendizado.  
> Sugestão: **CRUD de uma entidade rica** (ex: aeronaves militares, veículos, produtos).

| Fase | O projeto faz |
|------|--------------|
| 1 e 2 | CRUD local, sem banco, só memória |
| 3 | CRUD com banco real + API exposta |
| 4 | CRUD distribuído com cache e mensageria |
| 5 | CRUD com agente que interpreta linguagem natural |

Exemplo final da Fase 5:
```
Usuário: "me mostra os caças russos com mais de 2 motores"
    ↓
Agente interpreta a intenção
    ↓
Monta a query dinamicamente
    ↓
Retorna em linguagem natural
```

---

## Dois Repositórios Paralelos

```
Repo 1 — Fundamentos (Fases 1 a 3)
"A academia — aprender a linguagem"
Foco em linguagem pura, tipos, validações,
testes, organização mínima de projeto.
Sem pressão de arquitetura complexa.

Repo 2 — Labs (Fases 4 e 5)
"O sparring — integrar o mundo real"
APIs reais, LLMs, agentes, mensageria,
deploy, observabilidade, MCP.
Simula cenários de produção.
```

---

## Prompt Sugerido para Outro Chat

```
Estou aprendendo [LINGUAGEM]. 
Tenho 15 anos de experiência em .NET/C#.

Quero seguir este roadmap:
- Fase 1: Fundamentos
- Fase 2: Identidade da linguagem
- Fase 3: Conectar com o mundo
- Fase 4: Distribuído e profissional
- Fase 5: AI e agentes

Meu projeto âncora é um CRUD de [ENTIDADE].
Pode me guiar pela Fase [X], partindo do que já tenho em [LINK/CÓDIGO]?
Quero entender o porquê de cada decisão, não só o como.
Prefiro digitar o código — não gere tudo automaticamente.
```

---

> **Princípio central:** Você programa as capacidades. O LLM navega entre elas.  
> O mesmo vale pro seu aprendizado — você monta a estrutura, a curiosidade navega.
