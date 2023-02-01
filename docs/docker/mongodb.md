---
title: MongoDB docker-compose
tags:
    - docker
    - docker-compose
    - mongodb
---

No auth

```yaml title="docker-compose.yaml"
version: '3.1'

services:
  mongo:
    image: mongo
    environment:
      MONGO_DATA_DIR: /data/db
      MONGO_LOG_DIR: /dev/null
    ports:
      - "27017:27017"
```
