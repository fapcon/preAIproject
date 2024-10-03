## Деплой мастер узла

1. Инициализация мастер-узла. 
```shell
kubeadm init --pod-network-cidr=10.244.0.0/16 --cri-socket=unix:///var/run/cri-dockerd.sock --ignore-preflight-errors=NumCPU,Mem
```

- `pod-netword-cidr` - настройка виртуальных IP для CNI плагина Flannel
- `cri-socket` - указывается сокет cri-dockerd
- `ignore-preflight-errors` - ингорирование ошибок в случае, если машина не соответствует минимальным требованиям Kubernetes

По итогам инициализации, kubeadm сгенерирует join команду для подключения рабочих узлов в кластер. 

2. Создание админ-пользователя на машине мастер-узла для доступа к кластеру
```shell
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

Так же можно скопировать конфиг админа в переменную среду
```shell
export KUBECONFIG=/etc/kubernetes/admin.conf
```

4. Установка бинарных файлов для работы CNI Network
```
mkdir -p /opt/cni/bin
curl -O -L https://github.com/containernetworking/plugins/releases/download/v1.2.0/cni-plugins-linux-amd64-v1.2.0.tgz
tar -C /opt/cni/bin -xzf cni-plugins-linux-amd64-v1.2.0.tgz
```

3. Установка CNI-плагина Flannel. 
```shell
kubectl apply -f https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml
```

## Подключение рабочего узла

Для подключения можно воспользоваться сгенерированной join командой на мастер-узле, добавив флаг `--cri-socket=unix:///var/run/cri-dockerd.sock` для выбора конкретного cri-сокета при подключении

```shell
kubeadm join <ip-address>:<port>\
    --token=<token-from-step-2> \
    --discovery-token-ca-cert-hash sha256:<ca-hash-from-step-1> \
    --cri-socket=unix:///var/run/cri-dockerd.sock
```

Либо сгенерировать новую join команду и так же воспользоваться командой, добавив флаг `--cri-socket=unix:///var/run/cri-dockerd.sock`
```shell
kubeadm token create --print-join-command
```

После подключения рабочего узла, можно проверить работоспособность подключённых узлов, выполнив на мастер-узле команду:
```shell
kubectl get node
```

Пример вывода статус Ready:
```shell
NAME           STATUS   ROLES           AGE     VERSION
azeroth        Ready    control-plane   2d12h   v1.27.4
hollow-pezid   Ready    <none>          2d12h   v1.27.4
```

Так же можно проверить работоспособность подов в узлах, выполнив на мастер-узле команду:
```shell
kubectl get pods -A
```

Вывод:
```shell
NAMESPACE              NAME                                         READY   STATUS    RESTARTS       AGE
kube-flannel           kube-flannel-ds-4m5mz                        1/1     Running   0              2d12h
kube-flannel           kube-flannel-ds-hlvgc                        1/1     Running   0              2d12h
kube-system            coredns-5d78c9869d-4tpfn                     1/1     Running   0              2d12h
kube-system            coredns-5d78c9869d-ltk8w                     1/1     Running   0              2d12h
kube-system            etcd-azeroth                                 1/1     Running   0              2d12h
kube-system            kube-apiserver-azeroth                       1/1     Running   0              2d12h
kube-system            kube-controller-manager-azeroth              1/1     Running   0              2d12h
kube-system            kube-proxy-dr5kh                             1/1     Running   0              2d12h
kube-system            kube-proxy-v58pk                             1/1     Running   0              2d12h
kube-system            kube-scheduler-azeroth                       1/1     Running   0              2d12h
```

Kubernetes генерирует дополнительные поды на рабочие узлы для связи с ними.

## Подключение Dashboard
1. Установка конфига Kubernetes-dashboard. Если используются стандартные настройки мастер-ноды(taint не сняты), то доска будет автоматически установлена на рабочий узел.
```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
```

Проверка установки:
```shell
kubectl get all -n kubernetes-dashboard
```

Вывод:
```
NAME                                             READY   STATUS    RESTARTS   AGE
pod/dashboard-metrics-scraper-5cb4f4bb9c-xkf9g   1/1     Running   0          2d12h
pod/kubernetes-dashboard-6967859bff-p22xr        1/1     Running   0          2d12h

NAME                                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
service/dashboard-metrics-scraper   ClusterIP   10.103.103.173   <none>        8000/TCP        2d12h
service/kubernetes-dashboard        NodePort    10.109.52.81     <none>        443:30800/TCP   2d12h

NAME                                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/dashboard-metrics-scraper   1/1     1            1           2d12h
deployment.apps/kubernetes-dashboard        1/1     1            1           2d12h

NAME                                                   DESIRED   CURRENT   READY   AGE
replicaset.apps/dashboard-metrics-scraper-5cb4f4bb9c   1         1         1       2d12h
replicaset.apps/kubernetes-dashboard-6967859bff        1         1         1       2d12h
```

2. Для доступа к доске по IP узла, необходимо заменить в конфиге тип порта с **CusterIP** на **NodePort**

```
kubectl edit service/kubernetes-dashboard -n kubernetes-dashboard
```

Как должен выглядеть в итоге конфиг:
```yaml
ports: 
port: 443 
protocol: TCP 
targetPort: 8443 
selector: 
k8s-app: kubernetes-dashboard 
sessionAffinity: None 
type: NodePort
```

После настройки, для вступления в силу изменений, нужно удалить под доски. Он автоматически создаст новый под с изменениями. Вписать точное название доски из команды `kubectl get pods -A`
```shell
kubectl delete pod kubernetes-dashboard-<dashboard id> -n kubernetes-dashboard
```

Проверить, что тип(NodePort) и порт поменялись командой `kubectl get svc -A`
```
NAMESPACE              NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
default                kubernetes                  ClusterIP   10.96.0.1        <none>        443/TCP                  2d13h
kube-system            kube-dns                    ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   2d13h
kubernetes-dashboard   dashboard-metrics-scraper   ClusterIP   10.103.103.173   <none>        8000/TCP                 2d12h
kubernetes-dashboard   kubernetes-dashboard        NodePort    10.109.52.81     <none>        443:30800/TCP            2d12h   
```

3. Создание сервис-аккаунта для доски под именем dashboard
```
kubectl create serviceaccount dashboard -n kubernetes-dashboard
```

4. Выдача прав пользователю dashboard
```
kubectl create clusterrolebinding dashboard-admin -n kubernetes-dashboard  --clusterrole=cluster-admin  --serviceaccount=default:dashboard
```

Так же необходимо отредактировать конфиг clusterrolebinding для изменения пространства имён
```
kubectl edit clusterrolebinding
```

В конфиге нужно найти следующие строки, поменять последнюю строку на kubernetes-dashboard и выйти с сохранением файла.
```yaml
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRoleBinding
  metadata:
    creationTimestamp: "2023-07-23T19:17:31Z"
    name: dashboard-admin
    resourceVersion: "1617"
    uid: f7245349-f3f1-499c-a78e-a34d14734c53
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: cluster-admin
  subjects:
  - kind: ServiceAccount
    name: dashboard
    namespace: kubernetes-dashboard
```

6. Выдача токена пользователю dashboard и войти по адресу `http://<node-ip>:<node port>`, где
- node-ip - адрес рабочей машины
- node-port - адрес порта, указанный в `kubectl get svc -A`. Диапазон возможных портов `30000-32767`
```
kubectl create token -n kubernetes-dashboard dashboard
```

Можно указать флаг `--duration=<time>` для указания времени жизни токена. По стандарту время жизни 1 час

## Доступ к мастер-узлу удалённо с локальной машины

1. Загрузка Kubectl на локальную машину
```shell
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

2. Загрузка и проверка checksum
```shell
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
bash echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```

3. Установка Kubectl
```bash
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

4. Проверка на установку проверкой версии
```bash
kubectl version --client --short
```

5. Установка админ-пользователя на локальную машину
```
scp -r <user>@<master-node-ip>:/home/<remote-username>/.kube /home/<local-username>/  
```

Пример
```
scp -r root@0.0.0.0:/root/.kube /home/<local-username>/
```

6. Экспорт конфига в переменную среду
```shell
export KUBECONFIG=${HOME}/.kube/config 
```

Теперь можно использовать kubectl-команды к удалённому мастер-узлу.
```
kubectl get pods -A
```

## Удаление
### Docker
Удаление программ
```
docker image prune -a
systemctl restart docker

sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
```

Удаление связанных файлов
```
sudo rm -rf /var/lib/docker /etc/docker /var/run/docker.sock
sudo rm -rf /var/lib/containerd
sudo rm -f /etc/apparmor.d/docker
```

### Kubernetes
Удаление из кластера, удаление конфигов текущего узла
```
kubeadm reset
```

Удаление программ
```
apt remove -y kubeadm kubectl kubelet kubernetes-cni 
apt purge -y kube*
```

Удаление всех связанных файлов
```
sudo rm -rf ~/.kube
sudo rm -rf /etc/cni /etc/kubernetes /var/lib/dockershim /var/lib/etcd /var/lib/kubelet /var/lib/etcd2/ /var/run/kubernetes ~/.kube/* 
sudo rm -f /etc/systemd/system/etcd*
```

Удаление автоматически установленных пакетов
```
apt autoremove -y
```

Очистка межсетевого экрана
```
iptables -F && iptables -X
iptables -t nat -F && iptables -t nat -X
iptables -t raw -F && iptables -t raw -X
iptables -t mangle -F && iptables -t mangle -X
```