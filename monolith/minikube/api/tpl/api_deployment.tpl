apiVersion: apps/v1
kind: Deployment
metadata:
  name: tradeapi-deployment
  labels:
    app: tradeapi
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tradeapi
  template:
    metadata:
      labels:
        app: tradeapi
    spec:
      containers:
        - name: tradeapi
          image: tradeapi:latest
          imagePullPolicy: Never
          command: ["/bin/sh", "-c", "go mod tidy && go run ./cmd/api"]
          ports:
            - containerPort: ${SERVER_PORT}
          env:
            - name: DB_HOST
              value: tradedb-service
            - name: DB_PASSWORD
              value: ${DB_PASSWORD}
            - name: DB_USER
              value: ${DB_USER}
            - name: CACHE_ADDRESS
              value: tradecache-service
            - name: CACHE_PASSWORD
              value: ${CACHE_PASSWORD}
            - name: VIRTUAL_HOST
              value: earnapi.eazzygroup.org
            - name: LETSENCRYPT_HOST
              value: earnapi.eazzygroup.org
            - name: VIRTUAL_PORT
              value: "${SERVER_PORT}"
          readinessProbe:
              tcpSocket:
                port: ${SERVER_PORT}
              initialDelaySeconds: 60
              periodSeconds: 15

          livenessProbe:
              tcpSocket:
                port: 8080
              initialDelaySeconds: 60
              periodSeconds: 15
              failureThreshold: 4
status: {}