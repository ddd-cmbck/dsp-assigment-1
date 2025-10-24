# Spelling Bee — Distributed gRPC Word Game

### Developed in Go using gRPC, Protocol Buffers, and Design Patterns

---

## Overview

This project implements a **distributed spelling bee game** built with **Go** and **gRPC**.
It consists of three independent services communicating over RPC:

1. **Core Service** – Generates random letters, calculates user scores, and communicates with the Dictionary Service.
2. **Dictionary Service** – Validates if a submitted word exists in the English dictionary.
3. **Client Application** – Console-based interface for the player to play the spelling bee interactively.

The system demonstrates the use of **design patterns**, **scalable architecture**, and **modular design principles** with **separation of concerns** between services.

---

## System Architecture

```
┌────────────────────┐      gRPC      ┌─────────────────────┐
│    Client (UI)     │ <------------> │    Core Service     │
│ (Interactive Game) │                │ (Logic + Score Eval)│
└────────────────────┘                └────────────┬────────┘
                                                   │
                                                   │ gRPC
                                                   ▼
                                      ┌─────────────────────────┐
                                      │   Dictionary Service    │
                                      │ (Word Validation + DB)  │
                                      └─────────────────────────┘
```

Each service runs independently, exposing a gRPC API that enables inter-service communication with serialized protobuf messages.

---

## Functionality and Architecture 

### Core Service

* Generates **random sets of letters** (7 letters per game).
* Selects a **center letter** using a random utility function.
* Evaluates the **score of a word** based on:

  * Word length
  * Full pangram bonus (all 7 letters used)
  * Minimum length validation

### Dictionary Service

* Loads a **JSON-based dictionary** (`words_dictionary.json`) into memory using the **Singleton pattern** for efficiency.
* Provides the gRPC method `CheckWord` to validate if a word exists.

### Client Service

* Requests letters and submits words through the gRPC Core Service.
* Displays results, score updates, and handles user interaction in a console UI.

The system is **modular**, **scalable**, and **separated by responsibility**:

* `core/` → Logic and gRPC server for game mechanics
* `database/` → Dictionary microservice
* `client/` → User interface and network communication layer

---

## Design Patterns 

### 1. **Singleton Pattern** – in `word_checker.go`

* Ensures the dictionary file (`words_dictionary.json`) is loaded **only once** into memory.
* Implemented with `sync.Once` and `GetInstance()`.
* **Rationale:** Reduces memory overhead and guarantees thread-safe access to the dictionary data across gRPC calls.

```go
var once sync.Once
var instance *WordChecker
```

### 2. **Strategy Pattern** – in `grpc_server.go`

* The `CoreServer` injects functions for generating letters and selecting a center letter:

  ```go
  GenerateLetters func() []string
  PickCenter func([]string) string
  ```
* **Rationale:** Enables flexible substitution of word-generation strategies without changing server logic (e.g., for difficulty levels or different letter sets).

### 3. **Facade Pattern** – in `dictionary_client.go`

* The `DictionaryClient` encapsulates complex gRPC connection logic:

  ```go
  func (dc *DictionaryClient) WordExists(ctx context.Context, word string) (bool, error)
  ```
* **Rationale:** Simplifies interaction between Core Service and Dictionary Service by providing a clean interface for a single RPC call.

---

## Use of gRPC 

### Protobuf Definitions

#### **letters.proto**

Defines the RPC methods for the Core Service:

```proto
service Core {
  rpc GetLetters(LettersRequest) returns (LettersResponse);
  rpc GetScore(UserWord) returns (Score);
}
```

#### **dictionary.proto**

Defines the RPC methods for the Dictionary Service:

```proto
service Dictionary {
  rpc CheckWord(WordRequest) returns (WordResponse);
}
```

### gRPC Features Implemented

* Both proto files were compiled using `protoc-gen-go` and `protoc-gen-go-grpc`.
* Generated Go code is **imported and used** across client and server modules.
* Communication between services uses **typed protobuf messages**, ensuring type safety and performance.
* Core ↔ Dictionary and Client ↔ Core communication handled entirely via gRPC.

---

## Organization & Coding Standards 

* **Modular project layout:** `/client`, `/core`, `/database`, `/proto`
* **Encapsulation:** Clear separation of gRPC logic, utilities, and business rules.
* **Consistent naming:** Functions and variables follow Go conventions.
* **Error handling:** All RPC calls include contextual error logging.
* **Flags for configuration:** (`--port`, `--dictionary_addr`, `--core_addr`) for flexibility.

Example:

```bash
go run core/main.go --port=4000 --dictionary_addr=localhost:4050
go run database/main.go --port=4050
go run client/server.go --core_addr=localhost:4000
```

---

## Scoring System Example

| Word    | Letters Used | Pangram | Score            |
| ------- | ------------ | ------- | ---------------- |
| bee     | 3            | ❌       | 0 (too short)    |
| read    | 4            | ❌       | 1                |
| reading | 7            | ✅       | 14 (7 + bonus 7) |

---

## Technologies Used

* **Go 1.22+**
* **gRPC / Protobuf**
* **sync.Once (thread-safe singleton)**
* **flag (command-line configuration)**
* **JSON-based dictionary storage**

---

## How to Run

```bash
# 1. Run Dictionary Service
go run database/main.go --port=4050

# 2. Run Core Service
go run core/main.go --port=4000 --dictionary_addr=localhost:4050

# 3. Run Client
go run client/server.go --core_addr=localhost:4000
```

You’ll see:

```
Spelling Bee!
Enter \qt or press CTRL + C if you want to close the game
```

