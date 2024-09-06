#!/bin/bash

old_version=$(cat docs/swagger.json)

curl -X 'POST' \
    'https://converter.swagger.io/api/convert' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d "$old_version" > docs/openapi3.json