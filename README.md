# Chewbacca

Example app written using Go and React that follows the Standard Package Layout application design.

* The domain is isolated in the root `app` package and does not depend on any other package.
* Sub-packages such as `http` are adapters between domain and our implementation.
* Finally the `main` package also act as an adapter and ties everything together.

This means for example that we could swap out PostgreSQL for MySQL without affecting other dependencies.

### Prerequisites

* Docker CE >= 18.09.x
* Docker compose >= 1.23.x

### Build the app

```bash
make build
```

### Run the app

```bash
make start
```

Browse to [localhost:8080](http://localhost:8080) :rainbow:
