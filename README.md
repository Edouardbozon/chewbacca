# Chewbacca

Example app written using Go and React.

### Prerequisites

* Node 10.15.x
* Go 1.10.x
* Docker CE 18.09.x

### Start

Build server image

```
docker build -t chewbacca-server ./packages/server/
```

Build client image

```
docker build -t chewbacca-client ./packages/client/
```

Start everything

```bash
docker-compose up
```

Browse to [localhost:8080](http://localhost:8080)
