# Тестирование кода. Unit тесты.

Unit тесты — это разновидность тестов. Основная идея unit тестов состоит в проверке корректности работы каждого отдельного "юнита" кода, чтобы убедиться, что он выполняет свою функцию правильно. Юнитами могут быть как отдельные функции, так и методы, а также классы или структуры в объектно-ориентированных языках.

### Зачем нужны unit тесты:

1. Обеспечивают повышенную уверенность в корректности кода: когда вы тестируете каждый юнит по отдельности, вы можете быть уверены, что каждая его часть работает правильно.
2. Упрощают рефакторинг: если у вас уже есть тесты, то после внесения изменений вы сразу узнаете, не сломали ли вы что-то в других частях системы.
3. Документируют код: unit тесты служат как дополнительная документация, объясняя, как именно должны работать отдельные части кода.

Однако, иногда юнит-тестирование может быть сложным, если ваш код зависит от других сложных компонентов, например, баз данных, внешних API и т. д. В таких случаях часто используются mock-объекты, которые имитируют поведение зависимых компонентов, чтобы изолировать код и сделать его тестируемым.

## Mockery library
Mockery — это библиотека на языке программирования Go, которая позволяет легко создавать mock-объекты для интерфейсов. Она значительно упрощает процесс создания моков и позволяет генерировать их автоматически на основе интерфейсов.

### Установка библиотеки Mockery:
```
go get github.com/vektra/mockery/v2
```
### Создание mock-объектов:

Для использования Mockery необходимо, чтобы ваш код содержал интерфейсы, которые вы хотите замокать.

Затем мы можем сгенерировать mock-объекты для интерфейсов с помощью команды _**go:generate**_:

```css
go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name YourInterfaceName
```

## Использование mock-объектов в тестах:

После генерации моков, вы можете использовать их в ваших тестах для имитации поведения зависимых компонентов.
Ваши тесты будут работать с моками, а не с реальными зависимостями, что позволяет вам тестировать юниты независимо.

### Пример использования библиотеки Mockery в Go:
В нашем случае у нас есть следующий интерфейс ExchangeTicker, представляющий сервис по получению цен на различные пары криптовалют:

```go
type ExchangeTicker interface {
Save(ctx context.Context, tickers []models.ExchangeTicker) error
GetTicker(ctx context.Context) GetTickerOut
GetByID(ctx context.Context, tickerID int) (models.ExchangeTicker, error)
GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error)
}
```

Мы хотим протестировать компонент Ticker, который зависит от ExchangeTicker.


```css
go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeTicker
```
Это создаст файл **_mocks/ExchangeTicker.go_**, который содержит mock-объект ExchangeTicker.

### Пример теста с использованием mock-объекта:

```go
// ticker_test.go
func newTestLogger(conf config.AppConf) *zap.Logger {
ws := bytes.NewBuffer(make([]byte, 0, 1000))
logger := logs.NewLogger(conf, zapcore.AddSync(ws))
return logger
}

func TestTicker(t *testing.T) {
// Создаем mock-объект для ExchangeTicker
exchangeMock := &mocks.ExchangeTicker{}
// Создаем аргументы для функции

ctx := context.Background()

et := []models.ExchangeTicker{models.ExchangeTicker{},models.ExchangeTicker{}}

exchangeMock.On("Save", ctx, et).Return(nil)
logger := NewTestLogger(config.Appconfig{})
    // Передаем mock-объект в myTicker
    myTicker := NewTicker(exchangeMock, logger)

    // Вызываем метод, который использует ExchangeTicker
    err := myTicker.Save(ctx, et)

    // Проверяем, что метод был вызван с нужным аргументом
    exchangeMock.AssertCalled(t, "Save", ctx, et)

    // Проверяем, что ошибки нет
    assert.NoError(t, err)
}
```
