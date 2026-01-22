# LogLens

**LogLens** is a professional-grade, explainable log analysis engine written in Go.  
It focuses on **deterministic detection**, **temporal correlation**, and **audit-ready reporting**, designed for security analysis, auditing, forensic investigations, and root-cause analysis (RCA).

Unlike black-box or ML-heavy approaches, LogLens prioritizes **transparency, reproducibility, and explainability**.

---

## ✨ Key Features

- 🔍 **Rule-based detection (YAML)**
  - Statistical thresholds
  - Regex-based content inspection
  - Temporal correlation (sliding windows)

- ⏱️ **Temporal correlation**
  - Detect bursts and patterns within time windows
  - Deterministic sliding-window implementation

- 🧠 **Explainable findings**
  - Every alert is traceable to:
    - a rule
    - a log line
    - a concrete condition

- 🧹 **Finding deduplication**
  - Prevents alert flooding
  - Preserves individual evidences when required

- 📄 **Automated Markdown reporting**
  - Audit-ready reports
  - Clear executive summary + detailed findings

- 📦 **Structured JSON output**
  - Easy integration with pipelines, SIEMs, or scripts

- 🧪 **Tested core logic**
  - Parser
  - Regex rules
  - Deduplication behavior

---

## 🏗️ Architecture Overview

```
Log file
   ↓
Event normalization
   ↓
Parser
   ↓
Statistics
   ↓
Rules engine
   ├─ Threshold rules
   ├─ Regex rules
   └─ Temporal correlation
   ↓
Finding aggregation & deduplication
   ↓
Output (Console / JSON / Markdown report)
```

---

## 📁 Project Structure

```
loglens/
├── cmd/loglens/
├── internal/
│   ├── engine/
│   ├── event/
│   ├── parser/
│   ├── rules/
│   ├── finding/
│   └── report/
├── configs/
├── samples/
├── .vscode/
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### Requirements

- Go 1.21+

### Build

```bash
go build ./cmd/loglens
```

### Run

```bash
go run ./cmd/loglens -f samples/sample.log
```

---

## ⚙️ Usage

### Basic analysis

```bash
loglens -f app.log
```

### JSON output

```bash
loglens -f app.log --json
```

### Generate Markdown report

```bash
loglens -f app.log --report report.md
```

---

## 🧩 Rule Examples

### Threshold rule

```yaml
- id: excessive_errors
  description: Many errors detected
  level: ERROR
  threshold: 2
  severity: high
```

### Regex rule

```yaml
- id: connection_failure
  description: Connection failure detected
  severity: high
  pattern: "(failed to connect|timeout)"
```

### Temporal correlation rule

```yaml
- id: burst_errors
  description: Error burst in short time window
  level: ERROR
  threshold: 3
  window_seconds: 60
  severity: high
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 🎯 Design Principles

- Explainability first
- Deterministic behavior
- Auditability over opacity
- Incremental extensibility

---

## 📜 License

MIT

---

## ✍️ Author

Developed by **Filipe**
