# Slot Server
A web-based slot game server built with Go, featuring:
- Separate REST API and gRPC slot game servers
- MongoDB for data persistence
- Real-time game logic processing
- Web-based client interface


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
│   ├── auth/                   # Authorization Service
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
# Web Server build
docker build --target web-server -t web-server:latest .

# Slot Server build
docker build --target slot-server -t slot-server:latest .

docker compose up -d
```


## Usage

### API Endpoints
- `POST /api/auth/guest`: User authentication
- `POST /api/game/{slotId}/spin`: Process a spin request


