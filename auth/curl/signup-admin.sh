#!/bin/bash
curl -X POST http://localhost:8080/signup \
    -H "Content-Type: application/json" \
    -d "{
        \"first_name\": \"${1:-Alice}\",
        \"last_name\": \"${2:-Smith}\",
        \"password\": \"${3:-admin123456}\",
        \"email\": \"${4:-alice.smith@company.com}\",
        \"phone\": \"${5:-+1987654321}\",
        \"role\": \"${6:-ADMIN}\"
    }"