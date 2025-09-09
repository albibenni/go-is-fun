#!/bin/bash
curl -X GET http://localhost:8080/users -H "Authorization: Bearer $1"
