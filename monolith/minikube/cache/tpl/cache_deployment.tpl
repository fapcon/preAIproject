apiVersion: apps/v1
kind: Deployment
metadata:
  name: tradecache-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tradecache
  template:
    metadata:
      labels:
        app: tradecache
    spec:
      containers:
        - name: tradecache
          image: redis:7.0.2-alpine
          command: ["redis-server", "--appendonly", "yes", "--requirepass", "${CACHE_PASSWORD}"]
          ports:
            - containerPort: 6379
          env:
            - name: CACHE_PASSWORD
              value: ${CACHE_PASSWORD}
          volumeMounts:
            - name: cache-volume
              mountPath: /data
      volumes:
        - name: cache-volume
          hostPath:
            path: /path/to/cache/directory