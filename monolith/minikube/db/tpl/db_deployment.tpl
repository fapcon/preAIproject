apiVersion: apps/v1
kind: Deployment
metadata:
  name: tradedb-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tradedb
  template:
    metadata:
      labels:
        app: tradedb
    spec:
      containers:
        - name: tradedb
          image: postgres:14.4-alpine
          env:
            - name: POSTGRES_PASSWORD
              value: ${DB_PASSWORD}
            - name: POSTGRES_USER
              value: ${DB_USER}
            - name: POSTGRES_SSLMODE
              value: disable
          ports:
            - containerPort: ${DB_PORT}

