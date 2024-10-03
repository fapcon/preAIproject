apiVersion: v1
kind: Service
metadata:
  name: tradeapi-service
spec:
  selector:
    app: tradeapi
  type: NodePort
  ports:
    - protocol: TCP
      port: ${SERVER_PORT}
      targetPort: ${SERVER_PORT}
      nodePort: 30000
