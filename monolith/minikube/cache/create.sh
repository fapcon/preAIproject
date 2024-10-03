#!/bin/bash

service="cache"

set -a
source ../../.env
set +a

# Замена переменных окружения в шаблоне и создание манифестов
envsubst < ./tpl/${service}_deployment.tpl > ${service}_deployment.yaml
envsubst < ./tpl/${service}_service.tpl > ${service}_service.yaml

kubectl apply -f ${service}_deployment.yaml
kubectl apply -f ${service}_service.yaml

echo "Манифесты ${service} успешно обновлены."