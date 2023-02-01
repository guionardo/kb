---
title: PostgreSQL docker-compose
tags:
    - docker
    - docker-compose
    - postgresql
---

- User: postgres
- Password: postgres

Database persistence into local folder ```/home/guionardo/dev/docker/volumes/postgres```

```yaml title="docker-compose.yaml"
version: "3"
services:

  postgres:
    image: postgres
    environment:
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_USER=postgres
    ports:
      - 5432:5432
    volumes:
      - /home/guionardo/dev/docker/volumes/postgres:/var/lib/postgresql
```
