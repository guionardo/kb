---
title: MongoDB and PostgreSQL in portainer stack
tags: 
    - docker
    - docker-compose
    - portainer
    - database
    - mongodb
    - postgresql
---

Running databases as a [Portainer](https://portainer.io) stack.

```yaml
version: "2"
services:
  mongodb:
    image: mongo:4.4.19-focal
    environment:
     PUID: 1000
     PGID: 1000
     MONGO_DATA_DIR: /data/db
    volumes:
      - mongodb_data:/data/db
    ports:
      - 27017:27017
    restart: unless-stopped
    network_mode: host
    
  postgres:
    image: postgres
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_USER: secret
    ports:
      - 5432:5432
    volumes:
      - postgres_data:/var/lib/postgresql
    restart: unless-stopped
    network_mode: host

volumes:
  mongodb_data:
  postgres_data:
```
