#!/bin/bash
curl -X POST http://localhost:8080/v1/feed_follows \
    -H "Authorization: ApiKey $1" -H "Content-Type: application/json" \
    -d "{\"feed_id\": \"${2:-'673d504f-ddef-4bc2-8f97-84ca223cafc7'}\"}"
