#!/bin/bash

curl -X POST localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "Email": "user1@example.com",
    "Password": "secretpassword"
  }'
