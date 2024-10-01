#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose -v turso $DATABASE_URL up
