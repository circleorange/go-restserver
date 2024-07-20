#!/bin/bash

curl -X POST localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "Name": "Power ON",
    "Description": "Machines are running",
    "Location": "Galway",
    "DateTime": "2025-01-01T15:30:00.000Z"
  }'
