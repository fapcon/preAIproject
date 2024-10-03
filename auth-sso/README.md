# SSO Сервис

## Что в себя включает:

- **Auth**: аутентификация и авторизация.
- **Permissions**: права пользователя.
- **User Info**: информация о пользователях.

## Архитектура:

- **Transport** - отвечает за взаимодействие с внешним миром (gRPC Server / Routing запросов)
- **Service** - бизнес логика(Auth, Permissions)
- **Data** - репозитории и хранилища данных

## Стек технологий:

- [Go](https://go.dev/)
- [gRPC](https://grpc.io/)
- [Protobuf](https://protobuf.dev/)
- [JWT](https://jwt.io/)

## Инструкция для разработчика:
### Запуск и деплой:

- Локально: запустить local.sh и main.go в папке cmd/sso
- Деплой: запустить deploy.sh

### Для разработки:

- Установить protoc плагин для GO:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- Обновить PATH, чтобы использовать protoc:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

- Установить buf для генерации proto:

```bash
brew install bufbuild/buf/buf
```
- Запустить генерацию прото файлов
```bash
buf generate
```

gRPC: :8082
Можно протестировать с помощью Evans CLI:
```bash
evans -r repl -p 8082
```

## Пользователи для проверки oAuth авторизации
### Google
``` 
Почта: testpolzovatel308@gmail.com
Пароль: solo228322
```
http://localhost:8080/docs 
swagger документация
