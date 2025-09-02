#!/bin/bash
curl -X POST http://localhost:8080/v1/feeds/feed \
    -H "Authorization: ApiKey $1" -H "Content-Type: application/json" \
    -d "{\"name\": \"${2:-John feed}\", \"url\": \"${3:-url-example}\"}"
