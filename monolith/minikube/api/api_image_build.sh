#!/bin/bash

sudo docker build --no-cache  -t tradeapi:latest ../../

#Переключения контекста докера на виртуальную машину minikube
eval $(minikube docker-env)

#build образа api
docker build --no-cache  -t tradeapi:latest ../../
