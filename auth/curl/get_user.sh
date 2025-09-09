#!/bin/bash
curl -X GET http://localhost:8080/user/:id -H "Authorization: Bearer $1"
