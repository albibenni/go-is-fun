#!/bin/bash
curl -X DELETE "http://localhost:8080/v1/feed_follows/${2:-'0d7e6dfc-83c9-4c7e-8b39-0a10f8493fd0'}" \
    -H "Authorization: ApiKey $1"
