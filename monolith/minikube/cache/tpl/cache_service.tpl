apiVersion: v1
kind: Service
metadata:
  name: tradecache-service
spec:
  selector:
    app: tradecache
  ports:
    - protocol: TCP
      port: ${CACHE_PORT}
      targetPort: ${CACHE_PORT}
  type: LoadBalancer
