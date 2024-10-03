apiVersion: v1
kind: Service
metadata:
  name: tradedb-service
spec:
  selector:
    app: tradedb
  ports:
    - protocol: TCP
      port: ${DB_PORT}
      targetPort: ${DB_PORT}
  type: LoadBalancer
