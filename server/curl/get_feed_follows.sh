#!/bin/bash
curl -X GET "http://localhost:8080/v1/feed_follows" \
    -H "Authorization: ApiKey $1"
