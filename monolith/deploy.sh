#!/bin/bash
git stash
git pull
docker restart tradeapi
docker restart nginx-proxy