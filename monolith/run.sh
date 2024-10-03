#!/bin/bash

docker-compose -f docker-compose-db-local.yml up -d
docker-compose up -d