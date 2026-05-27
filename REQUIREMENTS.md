# Project: 'lea'
Structural context operating system for AI-native software engineering.
lea is a terminal-first structural memory system that helps AI models and developers understand large codebases with minimal context and maximum precision.
Instead of treating code as flat text, lea models repositories as living structural graphs.
The goal is not to build another AI coding assistant. The goal is to build the missing context infrastructure layer between codebases and AI systems.

Vision
Modern AI coding systems suffer from several fundamental problems:
* Context window limitations
* Token inflation
* Context entropy
* Hallucinations
* Architectural drift
* Weak retrieval quality
* Poor understanding of large repositories
Most existing systems rely heavily on:
* embeddings
* semantic chunking
* vector databases
* giant prompts
* autonomous agent loops
These approaches often fail because software systems are not just text.
Software systems are:
* deterministic
* symbolic
* graph-oriented
* architecture-driven
lea focuses on:
Structural retrieval first.
Semantic retrieval second.

Core Philosophy
What lea IS
* Structural memory system
* Graph-based context engine
* AI context compiler
* Terminal-native developer infrastructure
* Persistent codebase cognition layer

What lea is NOT
* IDE replacement
* Chatbot wrapper
* Another autonomous AI agent
* Electron application
* Traditional RAG platform
* Generic vector database

Problem Statement
Large language models are extremely sensitive to context quality.
The problem is often not:
Model intelligence
The real problem is:
Wrong context
or:
Too much context
Traditional RAG systems chunk source code into disconnected text fragments. This destroys the structural relationships inside software systems.
For example:
func LoginHandler()
is nearly meaningless without understanding:
LoginHandler
  -> AuthService.Login
      -> UserRepository.FindByEmail
      -> TokenService.Generate
          -> RedisStore.SaveSession
lea exists to preserve and expose these relationships.

Key Design Principles
1. Structural Retrieval First
Codebases should be retrieved through:
* symbols
* dependencies
* call graphs
* interfaces
* flows
* architecture boundaries
before semantic search.

2. Deterministic Over Probabilistic
Prefer:
* AST analysis
* symbol graphs
* exact references
* compiler-grade parsing
before embeddings.

3. Local-First
The system should:
* work offline
* run locally
* support SSH workflows
* support terminal-first environments
* avoid cloud dependency

4. Incremental Context
Do not re-index entire repositories.
Only:
* detect changed files
* invalidate affected graph nodes
* rebuild local subgraphs

5. AI-Agnostic Infrastructure
lea should support:
* Claude
* GPT
* Gemini
* local models
* Pi Agent
* OpenCode
* Aider
* MCP ecosystems
without depending on a specific vendor.

High-Level Architecture
                ┌─────────────────┐
                │   Codebase      │
                └────────┬────────┘
                         │
                  Parse / Analyze
                         │
                ┌────────▼────────┐
                │ Structural Graph│
                └────────┬────────┘
                         │
                  Memory Layer
                         │
                ┌────────▼────────┐
                │ Retrieval Engine│
                └────────┬────────┘
                         │
         ┌───────────────┼────────────────┐
         │               │                │
      CLI/TUI         MCP API         AI Context

Core System Components
1. Parser Engine
Responsibility
Extract:
* functions
* methods
* structs
* interfaces
* imports
* references
* package dependencies
* call graphs

Go Native Parsing
Primary tooling:
go/parser
go/ast
go/types
go/packages
These provide compiler-grade structural analysis.

Multi-Language Support
Future language support via:
* Tree-sitter
* language-specific parsers
Target languages:
* Go
* Rust
* TypeScript
* Python
* Lua
* Zig

2. Structural Graph Engine
Core Concept
The repository becomes a graph.

Node Types
Function
Method
Struct
Interface
Package
Module
Flow

Edge Types
CALLS
IMPLEMENTS
USES
IMPORTS
BELONGS_TO
DEPENDS_ON
FLOWS_THROUGH

Example Graph
LoginHandler
  CALLS -> AuthService.Login

AuthService.Login
  USES -> UserRepository

UserRepository
  IMPLEMENTS -> UserRepo interface

AuthService
  BELONGS_TO -> auth module

3. Storage Layer
Recommended Database
SQLite
Why:
* lightweight
* embedded
* fast
* portable
* supports recursive CTEs
* ideal for local-first tooling

Suggested Schema
nodes(
  id,
  type,
  name,
  file,
  metadata
)

edges(
  from_id,
  to_id,
  edge_type,
  metadata
)

Critical Indexes
CREATE INDEX idx_edges_from
ON edges(from_id);

CREATE INDEX idx_edges_to
ON edges(to_id);

CREATE INDEX idx_edges_type
ON edges(edge_type);

4. Incremental Indexing Engine
Goal
Avoid full repository re-indexing.

Strategy
When a file changes:
1. Detect affected symbols
2. Remove stale outbound edges
3. Remove stale inbound edges
4. Re-parse file
5. Reconnect local graph

Important Principle
Prefer:
invalidate and rebuild local subgraphs
instead of overly complex mutation systems.

File Watching
Recommended package:
fsnotify

5. Retrieval Engine
This is the heart of the system.
The goal is not visualization. The goal is fast structural understanding.

Core Retrieval Types
Neighbor Query
lea neighbors AuthService

Trace Query
lea trace LoginHandler
Output:
HTTP Request
  -> LoginHandler
      -> AuthService.Login
          -> UserRepository.Find
          -> TokenService.Generate

Impact Query
lea impact TokenService

Architecture Violations
lea violations

AI Context Export
lea context AuthService

6. AI Context Compiler
Purpose
Generate:
small high-signal context
for AI systems.

Example Output
## AuthService

Layer: service

Uses:
- UserRepository
- RedisCache

Called by:
- LoginHandler

Constraints:
- cannot access HTTP layer

Key Goal
Minimize:
* noisy context
* token waste
* hallucinations
* architectural confusion

7. MCP Integration
Future versions should expose:
* symbol graph queries
* flow tracing
* dependency retrieval
* architecture constraints
* context generation
through MCP-compatible APIs.

Example MCP Tools
get_symbol_neighbors
trace_execution_path
find_architecture_violations
get_related_context

CLI Philosophy
The CLI should expose:
graph intelligence
not giant visual graphs.

Example Commands
Search
lea find AuthService

Explain
lea explain AuthService

Trace
lea trace LoginFlow

AI Context
lea context auth

TUI Philosophy
Visualization is secondary.
The real value is:
structural retrieval
The TUI should focus on:
* fuzzy symbol navigation
* expandable dependency trees
* flow tracing
* architecture exploration

Recommended Technology Stack
Layer	Technology
Language	Go
CLI	Cobra
TUI	Bubble Tea
Parsing	go/ast + Tree-sitter
Storage	SQLite
File Watching	fsnotify
Search	ripgrep integration
Serialization	JSON
Config	YAML / TOML
AI Integration	MCP
Caching	SQLite / Badger
Development Roadmap
Phase 1 — MVP (Go Only)
Goal
Build a stable structural graph engine.

Features
* parse Go repositories
* extract symbols
* build dependency graph
* build basic call graph
* SQLite storage
* CLI retrieval

Commands
lea index
lea neighbors
lea trace
lea impact

Phase 1.5 — Stable Incremental Updates
Add
* file watching
* graph invalidation
* partial re-indexing

Phase 2 — AI Context Layer
Add
* context compiler
* architecture constraints
* flow summaries
* AI-optimized markdown export

Phase 3 — MCP Integration
Add
* MCP server
* AI tooling APIs
* external agent integration

Phase 4 — Interactive TUI
Add
* graph exploration
* fuzzy symbol navigation
* live tracing
* interactive architecture browsing

Phase 5 — Multi-Language Support
Add
* Tree-sitter integration
* Rust support
* TypeScript support
* Python support
* Lua support

Long-Term Direction
lea is not trying to become:
* a giant autonomous AI agent
* an IDE replacement
* a full software architect
Instead, it aims to become:
persistent structural cognition infrastructure
for:
* developers
* AI systems
* terminal-native workflows
* large repositories
* local models
* AI-native engineering

Final Thesis
The future of AI coding is likely not:
bigger prompts
+ bigger agents
+ more embeddings
The future is more likely:
smaller context
+ stronger retrieval
+ better structure
+ deterministic tooling
lea exists to help make that future practical.
