#!/bin/bash
curl -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d "{
        \"email\": \"${1:-alice.smith@company.com}\",
        \"password\": \"${2:-admin123456}\"
    }"