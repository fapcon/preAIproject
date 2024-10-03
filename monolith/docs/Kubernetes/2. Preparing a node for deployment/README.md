Следующие этапы необходимо проделать как на мастер-узле, так и на рабочих узлах.

## 1. Установка Docker
1. Обновление индексов пакетов и установка необходимых пакетов
```sh
sudo apt-get update && sudo apt-get install ca-certificates curl gnupg
```

2. Добавление GPG ключей официального Docker
```sh
sudo install -m 0755 -d /etc/apt/keyrings
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
$ sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

3. Настройка репозитория
```sh
 echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

4. Установка Docker и необходимых пакетов для него
```sh
 sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

Для проверки можно воспользоваться командой `docker -v`

## 2. Установка компонентов Kubernetes
1. Обновление индексов пакетов и установка необходимых пакетов
```shell
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
```

2. Установка GPG ключей и настройка репозитория
```shell
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
```

3. Установка Kubernetes и блокировка автоматической установки обновления/удаления
```shell
sudo apt-get update && sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

Для проверки можно воспользоватсья командой `kubectl version --short`

## 3. Установка Cri-dockerd
1. Обновление и установка необходимых пакетов
```shell
sudo apt update && sudo apt install git wget curl
```

2. Получаем последнюю версию выпуска
```shell
VER=$(curl -s https://api.github.com/repos/Mirantis/cri-dockerd/releases/latest|grep tag_name | cut -d '"' -f 4|sed 's/v//g')
echo $VER
```

3. Скачиваем и распаковываем архив с cri-dockerd
Для процессоров Intel x64
```shell
wget https://github.com/Mirantis/cri-dockerd/releases/download/v${VER}/cri-dockerd-${VER}.amd64.tgz tar xvf cri-dockerd-${VER}.amd64.tgz
```
Для процессоров ARM x64
```shell
wget https://github.com/Mirantis/cri-dockerd/releases/download/v${VER}/cri-dockerd-${VER}.arm64.tgz cri-dockerd-${VER}.arm64.tgz
```

4. Перемещаем cri-dockerd для запуска программы и проверяем версию
```shell
sudo mv cri-dockerd/cri-dockerd /usr/local/bin/
cri-dockerd --version
```

5. Скачиваем сервисную часть и сокет, настраиваем менеджер систем и служб ОС
```shell
wget https://raw.githubusercontent.com/Mirantis/cri-dockerd/master/packaging/systemd/cri-docker.service
wget https://raw.githubusercontent.com/Mirantis/cri-dockerd/master/packaging/systemd/cri-docker.socket
sudo mv cri-docker.socket cri-docker.service /etc/systemd/system/
sudo sed -i -e 's,/usr/bin/cri-dockerd,/usr/local/bin/cri-dockerd,' /etc/systemd/system/cri-docker.service
```

6. Презапускаем daemon и включаем сервис и сокет
```shell
sudo systemctl daemon-reload 
sudo systemctl enable cri-docker.service 
sudo systemctl enable --now cri-docker.socket
```

7. Проверяем статус сокета
```shell
systemctl status cri-docker.socket
```

Вывод:
```shell
● cri-docker.socket - CRI Docker Socket for the API
   Loaded: loaded (/etc/systemd/system/cri-docker.socket; enabled; vendor preset: disabled)
   Active: active (listening) since Fri 2023-03-10 10:02:13 UTC; 4s ago
   Listen: /run/cri-dockerd.sock (Stream)
    Tasks: 0 (limit: 23036)
   Memory: 4.0K
   CGroup: /system.slice/cri-docker.socket

Mar 10 10:02:13 rocky8.mylab.io systemd[1]: Starting CRI Docker Socket for the API.
Mar 10 10:02:13 rocky8.mylab.io systemd[1]: Listening on CRI Docker Socket for the API.
```

Можно так же протестировать загрузку образов контейнеров Kubernetes
```shell
sudo kubeadm config images pull --cri-socket /run/cri-dockerd.sock
```