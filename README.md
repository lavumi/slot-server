# Slot Server
web base slot game server


## Table of Contents

- [Frameworks](#frameworks)
    - [Server](#server)
    - [Database](#database)
    - [Web Front](#web-front)
- [Folder Structure](#folder-structure)
- [Installation](#installation)
- [Usage](#usage)

## Frameworks

Listed below are the versions of the frameworks used in this project.

### Server

- Golang: v1.19
- Gin: v1.9.1
- Protocol Buffer: v25.1

### Database

- MongoDb: v6.0.8

### Web Front

- Materialize: v1.0.0


### Folder Structure
```text
/slot-server/
├── cmd/
│   ├── main.go                 # run web, slot both
│   ├── web.go                  # REST api server start point
│   └── slot.go                 # gRPC slot server start point
├── internal/
│   ├── db/                     # database connections (mongodb, redis)
│   ├── server/                 # Server Service
│   │   ├── configs/          
│   │   ├── controllers/          
│   │   ├── forms/          
│   │   ├── middleware/          
│   │   ├── model/               
│   │   ├── rpc/                    # slot gRPC client
│   │   └── server.go               # server start point
│   └── slot/                   # Slot Service
│       ├── api/                    # slot api (made from protobuf)
│       ├── game/                   # Slot main logics
│       │     └── foodie/               # Foodie Reels (Slot00)
│       ├── model/                  # models for slot
│       ├── module/                 # modules
│       ├── proto/                  # slot protobuf
│       └── manager.go              # manager
└── web/                        # Web Front static pages
    ├── config/                     # configs for client ( symbol data and paytable )
    ├── script/                     # client code
    └── ...                         # Slot web view
```




## Installation

```bash
env GOOS=linux go build -o build/web-server cmd/web.go
env GOOS=linux go build -o build/slot-server cmd/slot.go
docker build -f Web.Dockerfile -t web-server:0.0.5 .
docker build -f Slot.Dockerfile -t slot-server:latest .
docker compose up -d
```
## Usage
