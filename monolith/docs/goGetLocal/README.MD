# Загрузка библиотеки из приватного репозитория studentgit.kata.academy для локальной разработки

## Инструкция:

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
igor@igor:~/go/src/ыгвщ кmonolith$ cat ~/.gitconfig
[user]
name = ...
email = ...@...
[url "ssh://git@studentgit.kata.academy/"]
    insteadOf = https://studentgit.kata.academy/
```
3. При необходимости удалить проблемный кэш: rm -rf /tmp/gopath/pkg/mod/cache/vcs/...
```
go: downloading gitlab.com/golight/orm.git v0.0.1
go: gitlab.com/golight/orm.git@v0.0.1: verifying module: gitlab.com/golight/orm.git@v0.0.1: reading https://sum.golang.org/lookup/gitlab.com/golight/orm.git@v0.0.1: 404 Not Found
    server response:
    not found: gitlab.com/golight/orm.git@v0.0.1: invalid version: git ls-remote -q origin in /tmp/gopath/pkg/mod/cache/vcs/6356c466db30609d610ed5d4646f40a414199532989f0bf16ae8fba4b80976ce: exit status 128:
```
4. Использовать команду go get с .git 
```
go get gitlab.com/golight/orm.git@v0.0.1
```
```
...:~/go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith$ go get gitlab.com/golight/orm.git@v0.0.1
    go: downloading gitlab.com/golight/orm.git v0.0.1
    go: gitlab.com/golight/orm.git@v0.0.1: parsing go.mod:
        module declares its path as: gitlab.com/golight/orm
        but was required as: gitlab.com/golight/orm.git
```

## Docker-compose
```
volumes:
- /$GOPATH/pkg/mod/cache/download/studentgit.kata.academy:/go/pkg/mod/cache/download/studentgit.kata.academy
```
"/$GOPATH/pkg/mod/cache/download/studentgit.kata.academy" том, который используется для хранения пакетов.
Он монтируется в директорию "/go/pkg/mod/cache/download/studentgit.kata.academy" внутри контейнера.
Это решает проблему с загрузкой пакетов из приватного репозитория при работе с docker.