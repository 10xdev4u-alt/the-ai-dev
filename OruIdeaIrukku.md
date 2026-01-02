# 🌐 PGit - Peer-to-Peer Git

> **Decentralized Git Collaboration Without Central Servers**

---

## 📖 Project Description

### One-Liner
```
PGit is a decentralized, peer-to-peer Git collaboration tool that enables 
developers to clone, push, and pull repositories directly between machines 
without relying on centralized platforms like GitHub, GitLab, or Bitbucket.
```

---

## 🎯 Problem Statement

### Current Reality
```
┌─────────┐         ┌─────────────┐         ┌─────────┐
│  Dev A  │ ──────▶ │   GitHub    │ ◀────── │  Dev B  │
└─────────┘         │  (Central)  │         └─────────┘
                    └─────────────┘
                          │
                    Single Point of
                    - Failure
                    - Censorship  
                    - Control
                    - Cost
```

### Problems We Solve

| Problem | Description |
|---------|-------------|
| **Single Point of Failure** | GitHub down = Work stops globally |
| **Censorship Risk** | Repos can be removed without notice |
| **Privacy Concerns** | All code visible to platform owner |
| **Internet Dependency** | No offline/LAN collaboration |
| **Cost** | Private repos, LFS, teams = $$$  |
| **Vendor Lock-in** | Migration is painful |

---

## 💡 Our Solution

### PGit Approach
```
┌─────────┐                           ┌─────────┐
│  Dev A  │ ◀───── Direct P2P ──────▶ │  Dev B  │
└─────────┘         Connection        └─────────┘
     ▲                                      ▲
     │                                      │
     └──────────▶ ┌─────────┐ ◀─────────────┘
                  │  Dev C  │
                  └─────────┘
                  
         No Central Server Required!
```

### What is PGit?

**PGit (Peer Git)** is a command-line tool that extends Git with peer-to-peer networking capabilities. It allows developers to:

- **Discover** other developers sharing repositories
- **Clone** repositories directly from peers
- **Push** changes to connected peers
- **Pull** updates from the network
- **Collaborate** without any central infrastructure

---

## ✨ Features

### Core Features (MVP)

| Feature | Description | Status |
|---------|-------------|--------|
| 🔍 **Peer Discovery** | Find peers sharing same repo via DHT | 🔄 Planned |
| 📥 **P2P Clone** | Clone repo directly from peer's machine | 🔄 Planned |
| 📤 **P2P Push** | Push commits to connected peers | 🔄 Planned |
| 📥 **P2P Pull** | Pull latest changes from peers | 🔄 Planned |
| 🔗 **NAT Traversal** | Works behind firewalls/routers | 🔄 Planned |
| 🔐 **Encryption** | All transfers encrypted | 🔄 Planned |

### Future Features (Post-MVP)

| Feature | Description |
|---------|-------------|
| 📋 **Issue Tracking** | Decentralized issues/discussions |
| 🔀 **Merge Requests** | P2P code review workflow |
| 👥 **Team Management** | Permissions and access control |
| 🌐 **Web UI** | Browser-based interface |
| 📱 **Mobile Sync** | Access repos from mobile |

---

## 🏗️ Architecture

### High-Level Overview

```
┌────────────────────────────────────────────────────────────┐
│                        CLI Layer                           │
│   ┌──────────────────────────────────────────────────┐    │
│   │  pgit clone <peer-id>/repo                       │    │
│   │  pgit push                                        │    │
│   │  pgit pull                                        │    │
│   │  pgit peers list                                  │    │
│   │  pgit share <repo>                                │    │
│   └──────────────────────────────────────────────────┘    │
├────────────────────────────────────────────────────────────┤
│                     Core Engine                            │
│   ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │
│   │ Git Manager │  │   Sync      │  │   Protocol      │   │
│   │             │  │   Engine    │  │   Handler       │   │
│   │ - Objects   │  │             │  │                 │   │
│   │ - Refs      │  │ - Diff      │  │ - Pack/Unpack   │   │
│   │ - Pack      │  │ - Merge     │  │ - Validation    │   │
│   └─────────────┘  └─────────────┘  └─────────────────┘   │
├────────────────────────────────────────────────────────────┤
│                   Network Layer (LibP2P)                   │
│   ┌───────────┐ ┌───────────┐ ┌───────────┐ ┌──────────┐  │
│   │    DHT    │ │  Gossip   │ │   NAT     │ │  Relay   │  │
│   │ Discovery │ │  PubSub   │ │ Traversal │ │  Server  │  │
│   └───────────┘ └───────────┘ └───────────┘ └──────────┘  │
├────────────────────────────────────────────────────────────┤
│                    Storage Layer                           │
│   ┌─────────────────────┐  ┌────────────────────────┐     │
│   │   Git Object Store  │  │   Peer/Config Store    │     │
│   │   (.git directory)  │  │   (BadgerDB/SQLite)    │     │
│   └─────────────────────┘  └────────────────────────┘     │
└────────────────────────────────────────────────────────────┘
```

### Component Details

#### 1. CLI Layer
```go
// Commands structure
pgit
├── clone   <peer-address>/repo-name    # Clone from peer
├── push    [peer-id]                   # Push to peers
├── pull    [peer-id]                   # Pull from peers
├── share   <repo-path>                 # Start sharing repo
├── peers
│   ├── list                            # List connected peers
│   ├── add     <peer-address>          # Add trusted peer
│   └── remove  <peer-id>               # Remove peer
├── repo
│   ├── list                            # List shared repos
│   └── info    <repo-name>             # Repo details
└── daemon                              # Run background service
```

#### 2. Network Layer (LibP2P)
```
┌─────────────────────────────────────────────────────┐
│                    LibP2P Host                      │
├─────────────────────────────────────────────────────┤
│  Transport    │  TCP, QUIC, WebSocket              │
│  Security     │  TLS 1.3, Noise Protocol           │
│  Multiplexing │  yamux, mplex                      │
│  Discovery    │  Kademlia DHT, mDNS (LAN)          │
│  PubSub       │  GossipSub (repo announcements)    │
│  NAT          │  AutoNAT, Hole Punching, Relay     │
└─────────────────────────────────────────────────────┘
```

#### 3. Protocol Design
```
┌──────────────────────────────────────────────────────────┐
│                  PGit Wire Protocol                      │
├──────────────────────────────────────────────────────────┤
│  Message Types:                                          │
│  ┌────────────────────────────────────────────────────┐  │
│  │  HANDSHAKE      - Initial peer authentication     │  │
│  │  REPO_LIST      - Request available repos         │  │
│  │  REF_LIST       - Request refs (branches/tags)    │  │
│  │  WANT           - Request specific objects        │  │
│  │  HAVE           - Declare available objects       │  │
│  │  PACK           - Send packfile data              │  │
│  │  ACK/NAK        - Acknowledgments                 │  │
│  └────────────────────────────────────────────────────┘  │
├──────────────────────────────────────────────────────────┤
│  Frame Format:                                           │
│  ┌──────┬──────┬─────────┬─────────────────┐            │
│  │ Len  │ Type │ RepoID  │     Payload     │            │
│  │ 4B   │ 1B   │  32B    │    Variable     │            │
│  └──────┴──────┴─────────┴─────────────────┘            │
└──────────────────────────────────────────────────────────┘
```

---

## 🔄 Workflow Examples

### Scenario 1: Sharing a Repository
```
Developer A wants to share project with team:

┌─────────────────────────────────────────────────────────┐
│ $ cd my-project                                         │
│ $ pgit share .                                          │
│                                                         │
│ ✓ Repository initialized for P2P sharing               │
│ ✓ Your Peer ID: 12D3KooWRfZ...                         │
│ ✓ Share URL: pgit://12D3KooWRfZ.../my-project          │
│                                                         │
│ Share this with your team members!                      │
└─────────────────────────────────────────────────────────┘
```

### Scenario 2: Cloning from Peer
```
Developer B wants to clone:

┌─────────────────────────────────────────────────────────┐
│ $ pgit clone pgit://12D3KooWRfZ.../my-project           │
│                                                         │
│ ⟳ Connecting to peer...                                │
│ ✓ Connected to 12D3KooWRfZ...                          │
│ ⟳ Fetching repository data...                          │
│ ████████████████████████ 100% (2.3 MB)                  │
│ ✓ Repository cloned successfully!                       │
│                                                         │
│ $ cd my-project                                         │
│ $ pgit peers list                                       │
│ ┌─────────────────────────────────────────┐             │
│ │ PEER ID              STATUS    LATENCY  │             │
│ │ 12D3KooWRfZ...       Online    23ms     │             │
│ └─────────────────────────────────────────┘             │
└─────────────────────────────────────────────────────────┘
```

### Scenario 3: Push/Pull Workflow
```
Developer B makes changes and pushes:

┌─────────────────────────────────────────────────────────┐
│ $ git add .                                             │
│ $ git commit -m "Add new feature"                       │
│ $ pgit push                                             │
│                                                         │
│ ⟳ Finding online peers...                              │
│ ✓ Found 2 peers                                         │
│ ⟳ Pushing to 12D3KooWRfZ... ✓                          │
│ ⟳ Pushing to 12D3KooWXyZ... ✓                          │
│ ✓ Pushed to 2/2 peers                                   │
└─────────────────────────────────────────────────────────┘

Developer A pulls changes:

┌─────────────────────────────────────────────────────────┐
│ $ pgit pull                                             │
│                                                         │
│ ⟳ Checking for updates...                              │
│ ✓ Found 1 new commit from 12D3KooWXyZ...               │
│ ⟳ Fetching changes...                                  │
│ ✓ Updated main: abc123 → def456                        │
└─────────────────────────────────────────────────────────┘
```

---

## 🛠️ Tech Stack

| Component | Technology | Why? |
|-----------|------------|------|
| **Language** | Go 1.21+ | Performance, concurrency, single binary |
| **P2P Network** | go-libp2p | Battle-tested, modular, NAT traversal |
| **Git Operations** | go-git | Pure Go, no CGO dependency |
| **Local Storage** | BadgerDB | Fast embedded KV store |
| **CLI Framework** | Cobra | Industry standard for Go CLIs |
| **Logging** | Zap | High-performance structured logging |
| **Config** | Viper | Flexible configuration management |

---

## 📁 Project Structure

```
pgit/
├── cmd/
│   └── pgit/
│       └── main.go                 # Entry point
│
├── internal/
│   ├── cli/                        # CLI commands
│   │   ├── root.go
│   │   ├── clone.go
│   │   ├── push.go
│   │   ├── pull.go
│   │   ├── share.go
│   │   ├── peers.go
│   │   └── daemon.go
│   │
│   ├── core/                       # Core business logic
│   │   ├── engine.go               # Main engine
│   │   ├── sync.go                 # Sync logic
│   │   └── config.go               # Configuration
│   │
│   ├── git/                        # Git operations
│   │   ├── repository.go           # Repo management
│   │   ├── objects.go              # Git objects
│   │   ├── packfile.go             # Pack/Unpack
│   │   └── refs.go                 # References
│   │
│   ├── p2p/                        # P2P networking
│   │   ├── host.go                 # LibP2P host
│   │   ├── discovery.go            # Peer discovery
│   │   ├── protocol.go             # Wire protocol
│   │   ├── handlers.go             # Message handlers
│   │   └── nat.go                  # NAT traversal
│   │
│   └── storage/                    # Storage layer
│       ├── store.go                # Storage interface
│       ├── badger.go               # BadgerDB impl
│       └── peers.go                # Peer storage
│
├── pkg/
│   ├── types/                      # Shared types
│   │   ├── messages.go
│   │   ├── repo.go
│   │   └── peer.go
│   │
│   └── utils/                      # Utilities
│       ├── crypto.go
│       └── helpers.go
│
├── configs/
│   └── default.yaml                # Default config
│
├── scripts/
│   ├── build.sh
│   └── test.sh
│
├── docs/
│   ├── ARCHITECTURE.md
│   ├── PROTOCOL.md
│   └── CONTRIBUTING.md
│
├── test/
│   ├── integration/
│   └── e2e/
│
├── .github/
│   └── workflows/
│       └── ci.yml
│
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── LICENSE
```

---

## 📅 Development Roadmap

### Phase 1: Foundation (Week 1-3)
```
┌─────────────────────────────────────────────────────────┐
│ Week 1: Project Setup                                   │
│ ├── □ Initialize Go module                              │
│ ├── □ Setup project structure                           │
│ ├── □ Configure CI/CD (GitHub Actions)                  │
│ ├── □ Setup linting (golangci-lint)                     │
│ └── □ Basic CLI skeleton with Cobra                     │
│                                                         │
│ Week 2: LibP2P Integration                              │
│ ├── □ Create LibP2P host                                │
│ ├── □ Implement peer identity (key generation)          │
│ ├── □ Setup DHT for peer discovery                      │
│ ├── □ Add mDNS for LAN discovery                        │
│ └── □ Test peer-to-peer connection                      │
│                                                         │
│ Week 3: Basic Protocol                                  │
│ ├── □ Define message types (protobuf)                   │
│ ├── □ Implement handshake protocol                      │
│ ├── □ Create stream handlers                            │
│ └── □ Test message exchange                             │
└─────────────────────────────────────────────────────────┘
```

### Phase 2: Git Integration (Week 4-6)
```
┌─────────────────────────────────────────────────────────┐
│ Week 4: Git Basics                                      │
│ ├── □ Integrate go-git library                          │
│ ├── □ Read local repository                             │
│ ├── □ List refs (branches, tags)                        │
│ └── □ Read git objects                                  │
│                                                         │
│ Week 5: Pack Protocol                                   │
│ ├── □ Generate packfiles                                │
│ ├── □ Parse packfiles                                   │
│ ├── □ Object negotiation (have/want)                    │
│ └── □ Delta compression                                 │
│                                                         │
│ Week 6: Clone Implementation                            │
│ ├── □ Implement REPO_LIST request                       │
│ ├── □ Implement REF_LIST request                        │
│ ├── □ Fetch and apply packfile                          │
│ └── □ Test full clone workflow                          │
└─────────────────────────────────────────────────────────┘
```

### Phase 3: Sync Operations (Week 7-9)
```
┌─────────────────────────────────────────────────────────┐
│ Week 7: Push Implementation                             │
│ ├── □ Detect local changes                              │
│ ├── □ Generate minimal packfile                         │
│ ├── □ Send to connected peers                           │
│ └── □ Handle acknowledgments                            │
│                                                         │
│ Week 8: Pull Implementation                             │
│ ├── □ Query peers for updates                           │
│ ├── □ Fetch missing objects                             │
│ ├── □ Update local refs                                 │
│ └── □ Handle conflicts (report to user)                 │
│                                                         │
│ Week 9: Multi-Peer Sync                                 │
│ ├── □ Sync from multiple peers                          │
│ ├── □ Peer selection strategy                           │
│ ├── □ Parallel fetching                                 │
│ └── □ Consistency verification                          │
└─────────────────────────────────────────────────────────┘
```

### Phase 4: Polish & Release (Week 10-12)
```
┌─────────────────────────────────────────────────────────┐
│ Week 10: NAT & Reliability                              │
│ ├── □ AutoNAT implementation                            │
│ ├── □ Relay server support                              │
│ ├── □ Hole punching                                     │
│ └── □ Connection retry logic                            │
│                                                         │
│ Week 11: Testing & Docs                                 │
│ ├── □ Unit tests (>70% coverage)                        │
│ ├── □ Integration tests                                 │
│ ├── □ End-to-end tests                                  │
│ ├── □ Documentation                                     │
│ └── □ Usage examples                                    │
│                                                         │
│ Week 12: Release                                        │
│ ├── □ Performance optimization                          │
│ ├── □ Security audit                                    │
│ ├── □ Build releases (Linux, macOS, Windows)            │
│ ├── □ Write announcement blog post                      │
│ └── □ Release v0.1.0                                    │
└─────────────────────────────────────────────────────────┘
```

---

## 🎯 Success Metrics

| Metric | Target |
|--------|--------|
| Clone Speed | Within 2x of Git over HTTPS |
| Connection Time | < 5 seconds to first peer |
| NAT Success Rate | > 80% direct connection |
| Memory Usage | < 100MB for typical repo |
| Binary Size | < 20MB |

---

## 🤝 Use Cases

### 1. Team Collaboration (No Server)
```
Small team, no budget for GitHub Teams
→ Use PGit for free private collaboration
```

### 2. Offline/LAN Development
```
Hackathon, Workshop, Remote location
→ Share code via local network, no internet needed
```

### 3. Censorship Resistance
```
Open source project in restricted region
→ Code cannot be taken down
```

### 4. Privacy-First Development
```
Sensitive project, NDA requirements  
→ Code never touches third-party servers
```

---

## 🚀 Getting Started (Future)

```bash
# Install
curl -sSL https://pgit.dev/install.sh | bash

# Or using Go
go install github.com/yourusername/pgit@latest

# Share your repo
cd your-project
pgit share .

# Clone from peer
pgit clone pgit://12D3KooW.../repo-name

# Push changes
pgit push

# Pull updates
pgit pull
```

---

## 📜 License

MIT License - Free for everyone!

---

**Bro, idhu full project description! Ready to start coding ah?
