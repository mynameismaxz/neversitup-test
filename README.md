```bash
├── README.md
├── bin
├── cmd
│   └── api
│       └── main.go
├── docs
│   └── api.yaml
├── go.mod
├── go.sum
├── internal
│   ├── dto
│   │   ├── requests.go
│   │   └── responses.go
│   ├── files
│   │   ├── delivery
│   │   │   └── http
│   │   │       ├── handlers.go
│   │   │       └── routes.go
│   │   ├── delivery.go
│   │   ├── mongo_repository.go
│   │   ├── repository
│   │   │   └── mongo_repository.go
│   │   ├── usecase
│   │   │   └── usecase.go
│   │   └── usecase.go
│   ├── middleware
│   │   ├── middleware.go
│   │   └── trace_middleware.go
│   ├── models
│   │   ├── files.go
│   │   └── token.go
│   ├── server
│   │   ├── handlers.go
│   │   └── server.go
│   └── token
│       ├── delivery
│       │   └── http
│       │       ├── handlers.go
│       │       └── routes.go
│       ├── delivery.go
│       ├── mongo_repository.go
│       ├── repository
│       │   └── mongo_repository.go
│       ├── usecase
│       │   └── usecase.go
│       └── usecase.go
├── makefile
├── pkg
│   └── database
│       └── mongo.go
├── tmp
└── zarf
    ├── compose
    │   ├── data
    │   │   ├── WiredTiger
    │   │   ├── WiredTiger.lock
    │   │   ├── WiredTiger.turtle
    │   │   ├── WiredTiger.wt
    │   │   ├── WiredTigerHS.wt
    │   │   ├── _mdb_catalog.wt
    │   │   ├── collection-0--3799931987700645826.wt
    │   │   ├── collection-0-5866647955042024691.wt
    │   │   ├── collection-2--3799931987700645826.wt
    │   │   ├── collection-4--3799931987700645826.wt
    │   │   ├── collection-7--3799931987700645826.wt
    │   │   ├── diagnostic.data
    │   │   │   ├── metrics.2023-11-22T10-26-47Z-00000
    │   │   │   ├── metrics.2023-11-22T10-26-52Z-00000
    │   │   │   ├── metrics.2023-11-22T10-32-46Z-00000
    │   │   │   ├── metrics.2023-11-23T02-41-37Z-00000
    │   │   │   ├── metrics.2023-11-26T12-15-50Z-00000
    │   │   │   ├── metrics.2023-11-28T23-15-48Z-00000
    │   │   │   └── metrics.interim
    │   │   ├── index-1--3799931987700645826.wt
    │   │   ├── index-1-5866647955042024691.wt
    │   │   ├── index-3--3799931987700645826.wt
    │   │   ├── index-5--3799931987700645826.wt
    │   │   ├── index-6--3799931987700645826.wt
    │   │   ├── index-8--3799931987700645826.wt
    │   │   ├── index-9--3799931987700645826.wt
    │   │   ├── journal
    │   │   │   ├── WiredTigerLog.0000000004
    │   │   │   ├── WiredTigerPreplog.0000000001
    │   │   │   └── WiredTigerPreplog.0000000002
    │   │   ├── mongod.lock
    │   │   ├── sizeStorer.wt
    │   │   └── storage.bson
    │   └── docker-compose.yml
    └── docker
        └── Dockerfile.api
```