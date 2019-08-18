# Chewbacca

Example app written using Go and React that follows the Standard Package Layout.

- The domain is isolated in the root `app` package and does not depend on any other package.
- Sub-packages such as `http` are adapters between domain and our implementation.
- Finally the `main` package act as an adapter and ties everything together.

This means for example that we could swap out PostgreSQL for MySQL without affecting other dependencies.

### Prerequisites

- Docker CE 18.09.x
- Docker compose 1.23.x
- Go 1.10.x
- Node 10.15.x
- Yarn 1.17.x

### Run the app

```bash
make start -j3
```

Browse to [localhost:3000](http://localhost:3000) :rainbow:
