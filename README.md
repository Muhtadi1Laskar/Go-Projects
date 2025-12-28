# 🧠 Go Projects

A single repository containing **multiple Go projects** focused on **computer science fundamentals**, **systems programming**, **cryptography**, **data structures**, and **practical tooling** — all implemented from scratch wherever possible.

This repo serves as:
- A **learning lab** for low-level and core CS concepts  
- A **portfolio** showcasing real engineering depth in Go  
- A **reference** for clean, minimal, dependency-light implementations  

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Paradigms:** Procedural, Data-Structure–Driven, Systems-Oriented
- **Focus Areas:**  
  - Algorithms & Data Structures  
  - Cryptography  
  - File Systems  
  - Networking  
  - Text Processing  
  - Performance & Correctness  

---

## 📁 Project Index

Each project is self-contained inside this monorepo.

---

### 1. 🎵 Audio Cutter
Cut and trim audio files programmatically.

**Concepts:**
- File I/O
- Binary data handling
- Audio stream manipulation

---

### 2. 🔍 Auto-Complete (BST)
A simple auto-complete engine implemented using a **Binary Search Tree**.

**Concepts:**
- Trees
- Recursive traversal
- Prefix matching

---

### 3. 🔐 BIP-39 (From Scratch)
Full implementation of **BIP-39 mnemonic generation** without external libraries.

**Concepts:**
- Cryptographic hashes
- Entropy → mnemonic conversion
- Checksum validation

---

### 4. ⛓️ Blockchain
A basic blockchain supporting both:
- **Proof-of-Work**
- **Proof-of-Stake**

**Concepts:**
- Distributed systems fundamentals
- Consensus mechanisms
- Hash-linked data structures

---

### 5. 🃏 Cards
A card game simulation.

**Features:**
- Shuffle deck
- Deal cards
- Save/load game state

**Concepts:**
- Structs
- Randomization
- Serialization

---

### 6. 🧬 Check Duplicate Files
Detect duplicate files in a directory using hashing.

**Concepts:**
- Hash functions
- File traversal
- Efficient comparisons

---

### 7. 🔏 Ciphers
Classic cipher implementations:
- Caesar
- Affine
- Vigenère
- Atbash

**Concepts:**
- Modular arithmetic
- Cryptography fundamentals

---

### 8. 📊 Count-Code
Counts total code files in a directory recursively.

**Features:**
- Filter by file extension
- Ignores `.git` and `node_modules`

**Concepts:**
- Recursive directory traversal
- File system APIs

---

### 9. 🧮 Evaluating Postfix Expression
Evaluates postfix expressions using a **stack**.

**Concepts:**
- Stack data structure
- Expression parsing

---

### 10. 🌊 Fourier Transformation
Implementation of **Fourier Transformation**.

**Concepts:**
- Mathematical computation
- Signal processing basics

---

### 11. 🧩 Identicon Generator
Generates unique visual patterns from hashes.

**Concepts:**
- Hashing
- Deterministic graphics
- Grid-based rendering

---

### 12. 🪪 JWT (From Scratch)
A full **JSON Web Token** implementation.

**Concepts:**
- Base64 encoding
- HMAC / signing
- Token validation

---

### 13. 🔑 Password Generator
Generate and validate secure passwords.

**Concepts:**
- Randomness
- Security constraints
- Validation logic

---

### 14. ❓ Quiz App
A simple terminal-based quiz application.

**Concepts:**
- CLI interaction
- Input validation
- Scoring systems

---

### 15. 📈 Sort-Data
Reads large datasets from Excel files, sorts them using **Quicksort**, and saves the output.

**Concepts:**
- Sorting algorithms
- File parsing
- Performance considerations

---

### 16. 🗂️ Sort-Files
Automatically sorts files inside a directory.

**Concepts:**
- File metadata
- OS-level operations

---

### 17. 🖥️ System Info
System monitoring tools including:
- Network monitor
- Process monitor
- Resource monitor

**Concepts:**
- OS introspection
- Performance metrics

---

### 18. 🤝 TCP Handshake
Simulation of a **TCP Handshake**.

**Concepts:**
- Networking fundamentals
- Protocol design

---

### 19. 📝 Text Analyzer
Cleans and analyzes text data.

**Features:**
- Tokenization
- Frequency analysis

---

### 20. 📄 Text Summary
A simple **non-AI** text summarization tool.

**Concepts:**
- Heuristic-based summarization
- Text ranking

---

### 21. ✅ To-Do App (Linked List)
A to-do application implemented using a **Linked List**.

**Concepts:**
- Pointer-based data structures
- CRUD operations

---

## 🚀 How to Run

Each project is independent.

```bash
cd <project-folder>
go run main.go
