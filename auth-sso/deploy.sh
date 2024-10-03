#!/bin/bash

git stash
git pull
docker-compose up --force-recreate --build