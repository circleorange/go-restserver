#!/bin/bash

curl -X PUT localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -d '{
    "Name": "Power OFF",
    "Description": "Machines are powered OFF",
    "Location": "Dublin",
    "DateTime": "2025-01-01T15:30:00.000Z"
  }'
