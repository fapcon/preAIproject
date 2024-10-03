#!/bin/bash

#Раскомментировать строку ниже для запуска minikube
minikube start

#Создание образа api внутри minikube
eval $(minikube docker-env)
docker build  --no-cache -t tradeapi:latest ../

for folder in */; do
    if [ -f "$folder/create.sh" ]; then
        echo "Запускается скрипт create.sh в папке: $folder"
        chmod +x "$folder/create.sh"

        cd "$folder"
        "./create.sh"
        cd ..
    else
        echo "Файл create.sh не найден в папке: $folder"
    fi
done
