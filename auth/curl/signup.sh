#!/bin/bash
curl -X POST http://localhost:8080/signup \
    -H "Content-Type: application/json" \
    -d "{
        \"first_name\": \"${1:-John}\",
        \"last_name\": \"${2:-Doe}\",
        \"password\": \"${3:-password123}\",
        \"email\": \"${4:-john.doe@example.com}\",
        \"phone\": \"${5:-+1234567890}\",
        \"role\": \"${6:-USER}\"
    }"
