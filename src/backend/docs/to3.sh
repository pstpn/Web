#!/bin/bash

old_version_v1=$(cat docs/v1/v1_swagger.json)
old_version_v2=$(cat docs/v2/v2_swagger.json)

curl -X 'POST' \
    'https://converter.swagger.io/api/convert' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d "$old_version_v1" > docs/v1/openapi3.json

curl -X 'POST' \
    'https://converter.swagger.io/api/convert' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d "$old_version_v2" > docs/v2/openapi3.json