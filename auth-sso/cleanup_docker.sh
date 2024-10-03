#!/bin/bash
docker container prune -f
docker volume prune -f
docker image prune -f
docker system prune -f
