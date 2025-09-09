#!/bin/bash
curl -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d "{
        \"email\": \"${1:-john.doe@example.com}\",
        \"password\": \"${2:-password123}\"
    }"