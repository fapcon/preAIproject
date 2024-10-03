# diamond-trade

## Usage
1. git clone https://studentgit.kata.academy/eazzyearn/students/mono/monolith.git
2. Настроить конфигурацию в .env файле.
3. docker-compose up --build

## Конфигурация приложения


### Конфигурация сервера:
<ul>
<li><code>ENV</code> - среда, в которой запускается сервер (production/dev/test).</li>
<li><code>DOMAIN</code> - домен, на котором запускается сервер.</li>
<li><code>SHUTDOWN_TIMEOUT</code> - время в секундах в течении которого сервер должен завершить свою работу.</li>
<li><code>SERVER_PORT</code> - порт на котором стартует сервер.</li>
</ul>

### Конфигурация JWT-токенов:
<ul>
<li><code>ACCESS_TTL</code> - время жизни Access-токена в минутах.</li>
<li><code>REFRESH_TTL</code> - время жизни Refresh-токена в днях.</li>
<li><code>ACCESS_SECRET</code> - символы для шифрования.</li>
<li><code>REFRESH_SECRET</code> - символы для шифрования.</li>
</ul>

### Конфигурация почтового сервиса:
<ul>
<li><code>EMAIL_FROM</code> - адрес почты с которой будут приходить уведомления.</li>
<li><code>EMAIL_PORT</code> - порт на котором стартует почтовой сервис.</li>
<li><code>EMAIL_HOST</code> - провайдер SMTP сервера (пр. smtp.gmail.com).</li>
<li><code>EMAIL_LOGIN</code> - логин Вашей почты.</li>
<li><code>EMAIL_PASSWORD</code> - пароль Вашей почты.</li>
<li><code>VERIFY_LINK_TTL</code> - время жизни верификационной ссылки.</li>
</ul>

### Конфигурация SQL DB:
<ul>
<li><code>DB_NET</code> - протокол, через который устанавливается соединение.</li>
<li><code>DB_DRIVER</code> - дравйвер базы данных.</li>
<li><code>DB_NAME</code> - имя базы данных.</li>
<li><code>DB_USER</code> - имя пользователя в базе данных.</li>
<li><code>DB_PASSWORD</code> - пароль пользователя в базе данных.</li>
<li><code>DB_HOST</code> - адрес по которому происходит подключение к базе данных.</li>
<li><code>MAX_CONN</code> - максимальное количество одновременных соединений.</li>
<li><code>DB_PORT</code> - порт на котором база данных принимает соединение.</li>
<li><code>DB_TIMEOUT</code> - время в минутах по истечении которого отменяется запрос.</li>
</ul>

### Конфигурация NoSQL DB:
<ul>
<li><code>MONGO_USER</code> - имя пользователя в базе данных.</li>
<li><code>MONGO_PASSWORD</code> - пароль пользователя в базе данных.</li>
<li><code>MONGO_HOST</code> - адрес по которому происходит подключение к базе данных.</li>
<li><code>MONGO_PORT</code> - порт на котором база данных принимает соединение.</li>
<li><code>MONGO_NAME</code> - имя базы данных.</li>
</ul>

### Конфигурация кэша:
<ul>
<li><code>CACHE_PASSWORD</code> - пароль от кэша.
<li><code>CACHE_ADDRESS</code> - адрес по которому происходит подключение к базе данных.</li>
<li><code>CACHE_PORT</code> - порт на котором база данных принимает соединение.</li>
</ul>

### Конфигурация RPC:
<ul>
<li><code>RPC_SHUTDOWN_TIMEOUT</code> - время в секундах в течении которого RPC-сервер должен завершить свою работу.</li>
<li><code>RPC_PORT</code> - порт на котором RPC-сервер принимает соединение.</li>
<li><code>RPC_ADDRESS</code> - адрес по которому происходит подключение к RPC-серверу.</li>
</ul>

## Загрузка библиотеки из приватного репозитория studentgit.kata.academy:

### Инструкция:

1. Добавить в GOPRIVATE приватный репозиторий:
```
go env -w GOPRIVATE="studentgit.kata.academy/*"
```
2. Добавить в .gitconfig приватный репозиторий.
```
  git config --global url.ssh://git@studentgit.kata.academy/.insteadOf https://studentgit.kata.academy/
```
- Проверить
```
cat ~/.gitconfig
```
```
name@name:~/go/src/name monolith$ cat ~/.gitconfig
[user]
name = ...
email = ...@...
[url "ssh://git@studentgit.kata.academy/"]
    insteadOf = https://studentgit.kata.academy/
```
3. При необходимости удалить проблемный кэш: rm -rf /tmp/gopath/pkg/mod/cache/vcs/...
   (либо удалить кеши в go/pkg/mod/cache/vcs;
   либо go clean --modcache;
   либо go mod clean)
```
go: downloading studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git v<Версия>
go: studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>: verifying module: gitlab.com/golight/orm.git@v<Версия>: reading https://sum.golang.org/lookup/studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>: 404 Not Found
    server response:
    not found: studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>: invalid version: git ls-remote -q origin in /tmp/gopath/pkg/mod/cache/vcs/6356c466db30609d610ed5d4646f40a414199532989f0bf16ae8fba4b80976ce: exit status 128:
```
4. Использовать команду go get с .git
```
go get studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>
```
```
...:~/go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith$ go get studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>
    go: downloading studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git v<Версия>
    go: studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>: parsing go.mod:
        module declares its path as: gitlab.com/golight/orm
        but was required as: studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git
```

### Docker-compose
```
volumes:
- /$GOPATH/pkg/mod/cache/download/studentgit.kata.academy:/go/pkg/mod/cache/download/studentgit.kata.academy
```
"/$GOPATH/pkg/mod/cache/download/studentgit.kata.academy" том, который используется для хранения пакетов.
Он монтируется в директорию "/go/pkg/mod/cache/download/studentgit.kata.academy" внутри контейнера.
Это решает проблему с загрузкой пакетов из приватного репозитория при работе с docker.

### В случае если данный вариант не помог, тогда необходимо:
```
1) создать новый ssh названием ключа id_gitkata
2) в папке .ssh создать config и прописать в нем
   Host studentgit.kata.academy
   User <Твой_Юзернейм>
   IdentityFile ~/.ssh/id_gitkata
3) вставить его в аккаунт gitlaba kata
4) выполнить в терминале eval "$(ssh-agent -s)"
5) выполнить в терминале ssh-add ~/.ssh/id_gitkata
6) выполнить команду ssh -T git@studentgit.kata.academy  и проверить, чтобы было Welcome to GitLab, @<Твой_Юзернейм>!
7) удалить кеши в go/pkg/mod/cache/vcs
8) go env -w GOPRIVATE="studentgit.kata.academy/*"
9) git config --global url.ssh://git@studentgit.kata.academy/.insteadOf https://studentgit.kata.academy/
10) go get studentgit.kata.academy/eazzyearn/students/<Внешний репозиторий>.git@v<Версия>
```

## Обновление basemod docker-образа

### Создание локального образа

1. В произвольном каталоге создаем Docker-файл следующего содержимого:
```
FROM golang:1.19.2-alpine
ADD ./download /go/pkg/mod/cache/download
```
2. В этом же каталоге создаем каталог \"download\":
```bash
mkdir download
```
3. Копируем в данный католог все необходимые пекеты из локального кэша, например:
```bash
cp -r $GOPATH/pkg/mod/cache/download/studentgit.kata.academy ./download
```
4. Создаем локальный docker-image (создается с тегом \"latest\"):
```bash
docker build . -t eazzygroup/basemod
```
5. Проверяем наличие образа:
```bash
docker images | grep eazzygroup/basemod
```
6. Тестируем локально запуск сервера через docker-compose с данным контейнером.
При возникновении проблем можно исследовать содержимое образа на наличие всех необходимых данных
с помощью утилиты [dive](https://github.com/wagoodman/dive):
```bash
dive eazzygroup/basemod
```

### Загружаем образ на Docker-Hub
1. Логинемся на Docker-Hub:
```bash
docker login
```
2. Пушим образ (пушится с тегом \"latest\"):
```bash
docker push eazzygroup/basemod
```
3. Для повышения надежности дополнительно создаем и пушим образ с указанием версии,
чтобы при следующем создании и пуше некорректного образа с тегом latest можно было откатиться:
```bash
docker build . -t eazzygroup/basemod:0.0.2
docker push eazzygroup/basemod:0.0.2
```
