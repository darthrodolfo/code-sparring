# Roadmap: Do Zero ao Avancado em Qualquer Linguagem

> Estrutura generica e progressiva para aprender qualquer linguagem de programacao com profundidade real.
> Adaptavel para: Go, Python, TypeScript, Rust, Kotlin, Java, ou qualquer outra.

---

## Estrutura de Repositorios

```
/linguagem
  /01-fundamentos
  /02-linguagem-core
  /03-conectando-mundo
  /04-distribuido
  /05-agentes
```

Cada pasta com um README curto explicando: o que foi estudado, por que, o que aprendeu, pacotes e ferramentas.

---

## Fase 1 — Fundamentos da Linguagem

> O que sustenta tudo. Nao e basico — e a base que nao pode ter furo.

- Hello World
- Tipos primitivos e compostos
- Variaveis, constantes e escopo
- Condicionais e loops
- Funcoes
- Tratamento de erro (como ESSA linguagem trata)
- Colecoes — array, slice, map, lista
- Structs / classes / records
- Interfaces / traits / protocolos
- Generics

### Regra de mentoria para Fase 1

- Sempre que surgir sintaxe nova, explicar no ato em 2-5 linhas com exemplo minimo.
- Nao esperar o aluno travar para explicar semantica basica de loops, atribuicao, ponteiros/referencias, escopo e tratamento de erro.
- Em qualquer transicao de stack, explicitar diferencas de modelo mental entre a linguagem anterior e a atual.

---

## Fase 2 — A Linguagem Sendo Ela Mesma

> O que essa linguagem faz diferente de todas as outras. Aqui esta a identidade dela.

- O diferencial da linguagem:
  - **Rust** — ownership, borrow checker, lifetimes
  - **Go** — goroutines, channels, simplicidade forcada
  - **Python** — duck typing, dynamic, ecosystem
  - **TypeScript** — tipagem estatica sobre JS, structural types
- Concorrencia e paralelismo
- Gerenciamento de memoria
- Modulos e organizacao de projeto
- Testes unitarios e de integracao
- Manipulacao de string, arquivo e JSON

---

## Fase 3 — Conectando com o Mundo

> A linguagem saindo do isolamento e interagindo com sistemas reais.

- HTTP Client — consumir API externa
- HTTP Server — expor API propria (minimal API)
- Banco de dados — leitura e escrita (SQL e/ou NoSQL)
- Variaveis de ambiente e configuracao
- Logging e observabilidade basica
- Docker — containerizar a aplicacao
- Deploy simples — subir em algum lugar real

---

## Fase 4 — Distribuido e Profissional

> Como a linguagem se comporta em sistemas reais de producao.

- Autenticacao e seguranca
- Mensageria — RabbitMQ, Kafka ou Redis
- gRPC
- Cache
- Testes de carga e performance
- CI/CD basico

---

## Fase 5 — AI e Agentes

> O diferencial do momento. A linguagem interagindo com inteligencia artificial.

### 5.1 — Consumindo LLM

- Integrar API de LLM (OpenAI, Anthropic, Gemini)
- Enviar prompt, receber resposta
- Gerenciar contexto e historico de conversa

### 5.2 — LLM Local

- Rodar LLM local via Ollama
- Integrar com modelo offline
- Comparar performance local vs cloud

### 5.3 — RAG (Retrieval Augmented Generation)

- Embedding — transformar texto em vetor
- Vector Database — ChromaDB, Pinecone, pgvector
- Similarity Search — busca por significado, nao por palavra
- Chunking — dividir documentos grandes
- Pipeline completo: documento -> embedding -> busca -> resposta

### 5.4 — Agente Simples

- Tool calling — plugar funcoes no LLM
- Loop de decisao — o LLM decide qual tool chamar
- Criterio de parada — quando o agente conclui
- Memoria de contexto entre iteracoes

### 5.5 — Multi-Agente

- Orquestrador sem LLM — fluxo fixo entre agentes
- Orquestrador com LLM — fluxo dinamico
- Handoff entre agentes
- Agentes especializados com responsabilidade unica

### 5.6 — MCP (Model Context Protocol)

- Criar servidor MCP na linguagem
- Expor tools via protocolo padrao
- Conectar com Claude, Cursor, ou qualquer cliente MCP

---

## Projeto Ancora por Fase

> Um projeto simples e consistente que cresce junto com o aprendizado.
> Sugestao: CRUD de uma entidade rica (ex: aeronaves militares, veiculos, produtos).

| Fase | O projeto faz |
|------|--------------|
| 1 e 2 | CRUD local, sem banco, so memoria |
| 3 | CRUD com banco real + API exposta |
| 4 | CRUD distribuido com cache e mensageria |
| 5 | CRUD com agente que interpreta linguagem natural |

Exemplo final da Fase 5:

```
Usuario: "me mostra os cacas russos com mais de 2 motores"
  -> Agente interpreta a intencao
  -> Monta a query dinamicamente
  -> Retorna em linguagem natural
```

---

## Dois Repositorios Paralelos

```
Repo 1 — Fundamentos (Fases 1 a 3)
  "A academia — aprender a linguagem"
  Foco em linguagem pura, tipos, validacoes,
  testes, organizacao minima de projeto.
  Sem pressao de arquitetura complexa.

Repo 2 — Labs (Fases 4 e 5)
  "O sparring — integrar o mundo real"
  APIs reais, LLMs, agentes, mensageria,
  deploy, observabilidade, MCP.
  Simula cenarios de producao.
```

---

## Prompt Sugerido para Novo Chat

```
Estou aprendendo [LINGUAGEM].
Tenho 15 anos de experiencia em .NET/C#.

Quero seguir este roadmap:
- Fase 1: Fundamentos
- Fase 2: Identidade da linguagem
- Fase 3: Conectar com o mundo
- Fase 4: Distribuido e profissional
- Fase 5: AI e agentes

Meu projeto ancora e um CRUD de [ENTIDADE].
Pode me guiar pela Fase [X], partindo do que ja tenho em [LINK/CODIGO]?
Quero entender o porque de cada decisao, nao so o como.
Prefiro digitar o codigo — nao gere tudo automaticamente.
```

---

> Principio central: Voce programa as capacidades. O LLM navega entre elas.
> O mesmo vale pro seu aprendizado — voce monta a estrutura, a curiosidade navega.
