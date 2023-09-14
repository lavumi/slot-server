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

### Database

- MongoDb: v6.0.8

### Web Front

- Materialize: v1.0.0


### Folder Structure
```text
/slot-server/
├── cmd/
│   ├── main.go                 # server start point
│   └── test.go                 # for slot unit test...
├── configs/                    # slot par sheet and slot list
├── internal/
│   ├── auth/                   # Authorization Service
│   ├── database/               # Mongo db data connections
│   ├── server/                 # Server Service
│   │   ├── middleware/          
│   │   ├── model/               
│   │   ├── router/              
│   │   └── server.go           # server start point
│   └── slot/                   # Slot Service
│       ├── config/             # Config struct
│       ├── model/              # models for slot
│       ├── module/             # modules
│       ├── slot00/             # slot 0
│       └── manager.go          # manager
└── web/                        # Not used. We had plans to set up RAID 1 later.
    ├── config/                 # configs for client ( symbol data and paytable )
    ├── script/                 # client code
    └── ...                     # Slot web view
```




## Installation

todo improve

```bash
env GOOS=linux go build -o slot-server cmd/main.go
docker build -t slot-server:latest .
docker save -o slot-server.tar slot-server
rm slot-server
```
## Usage

```bash
go run cmd/main.go
```



[//]: # (- at remote)
[//]: # (```bash)
[//]: # ( docker load -i slot-server.tar)
[//]: # (```)