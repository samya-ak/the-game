# the-game

# Getting Started
## Prerequisites
- Docker
- go1.19

## Major Dependencies
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [gqlgen](https://github.com/99designs/gqlgen)
- [testify](https://github.com/stretchr/testify)
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

## Setup
Run the app for first time
``` shell
git clone git@github.com:samya-ak/the-game.git
make -C the-game/src start-local-docker
```

Available make commands
``` shell
# build image and start containers while removing unused containers
make start-local-docker

# stop containers
make stop-local-docker

# combination of above two commands, stop before starting again
make stop-start

# generates models, boilerplate resolvers using gqlgen from graphql schema
make schema
```

## Suggestions for improvements
- Integrate tracing to track request lifecycle
- Implement authentication and authorization for enhanced security
- Integrate Metrics for request-response latency, failure rate, status code frequency etc.
- Export logs to remote dashboard
- Implement Migration management and versioning
- Integrate CI/CD to ship app quickly and efficiently
- Implement test coverage dashboards

## License
MIT
