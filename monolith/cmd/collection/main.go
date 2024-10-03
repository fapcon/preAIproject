//go:generate grizzly generate main.go
package main

import (
	"fmt"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

//grizzly:generate
type ExchangeOrder struct {
	ID           int64          `json:"id" mapper:"id"`
	UUID         string         `json:"uuid" mapper:"uuid"`
	OrderID      int64          `json:"order_id" mapper:"order_id"`
	UserID       int            `json:"user_id" mapper:"user_id"`
	ExchangeID   int            `json:"exchange_id" mapper:"exchange_id"`
	UnitedOrders int            `json:"united_orders" mapper:"united_orders"`
	OrderType    int            `json:"order_type" mapper:"order_type"`
	OrderTypeMsg string         `json:"order_type_msg" mapper:"order_type_msg"`
	Pair         string         `json:"pair" mapper:"pair"`
	Amount       float64        `json:"amount" mapper:"amount"`
	Quantity     float64        `json:"quantity" mapper:"quantity"`
	Price        float64        `json:"price" mapper:"price"`
	Side         int            `json:"side" mapper:"side"`
	SideMsg      string         `json:"side_msg" mapper:"side_msg"`
	Message      string         `json:"message" mapper:"message"`
	Status       int            `json:"status" mapper:"status"`
	StatusMsg    string         `json:"status_msg" mapper:"status_msg"`
	SumBuy       int            `json:"sumBuy" mapper:"sumBuy"`
	ApiKeyID     int            `json:"api_key_id" mapper:"api_key_id"`
	CreatedAt    time.Time      `json:"created_at" mapper:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" mapper:"updated_at"`
	DeletedAt    types.NullTime `json:"deleted_at" mapper:"deleted_at"`
}

func main() {
	orders := []*ExchangeOrder{
		{
			ID:           1,
			UUID:         "uuid1",
			OrderID:      123,
			UserID:       456,
			ExchangeID:   789,
			UnitedOrders: 0,
			OrderType:    1,
			OrderTypeMsg: "order type 1",
			Pair:         "BTC/USD",
			Amount:       10.5,
			Quantity:     2.5,
			Price:        4.2,
			Side:         1,
			SideMsg:      "side 1",
			Message:      "message 1",
			Status:       0,
			StatusMsg:    "status 1",
			SumBuy:       100,
			ApiKeyID:     1000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           2,
			UUID:         "uuid2",
			OrderID:      456,
			UserID:       789,
			ExchangeID:   123,
			UnitedOrders: 1,
			OrderType:    2,
			OrderTypeMsg: "order type 2",
			Pair:         "ETH/BTC",
			Amount:       20.7,
			Quantity:     3.8,
			Price:        5.9,
			Side:         2,
			SideMsg:      "side 2",
			Message:      "message 2",
			Status:       1,
			StatusMsg:    "status 2",
			SumBuy:       200,
			ApiKeyID:     2000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           3,
			UUID:         "uuid3",
			OrderID:      789,
			UserID:       123,
			ExchangeID:   456,
			UnitedOrders: 0,
			OrderType:    1,
			OrderTypeMsg: "order type 1",
			Pair:         "ETH/USD",
			Amount:       15.2,
			Quantity:     2.2,
			Price:        6.5,
			Side:         1,
			SideMsg:      "side 1",
			Message:      "message 3",
			Status:       1,
			StatusMsg:    "status 2",
			SumBuy:       150,
			ApiKeyID:     3000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           4,
			UUID:         "uuid4",
			OrderID:      987,
			UserID:       654,
			ExchangeID:   321,
			UnitedOrders: 1,
			OrderType:    2,
			OrderTypeMsg: "order type 2",
			Pair:         "LTC/BTC",
			Amount:       8.9,
			Quantity:     1.7,
			Price:        5.2,
			Side:         2,
			SideMsg:      "side 2",
			Message:      "message 4",
			Status:       0,
			StatusMsg:    "status 1",
			SumBuy:       80,
			ApiKeyID:     4000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           5,
			UUID:         "uuid5",
			OrderID:      654,
			UserID:       321,
			ExchangeID:   987,
			UnitedOrders: 0,
			OrderType:    1,
			OrderTypeMsg: "order type 1",
			Pair:         "XRP/USD",
			Amount:       5.6,
			Quantity:     1.2,
			Price:        4.7,
			Side:         1,
			SideMsg:      "side 1",
			Message:      "message 5",
			Status:       1,
			StatusMsg:    "status 2",
			SumBuy:       60,
			ApiKeyID:     5000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           6,
			UUID:         "uuid6",
			OrderID:      321,
			UserID:       987,
			ExchangeID:   654,
			UnitedOrders: 1,
			OrderType:    2,
			OrderTypeMsg: "order type 2",
			Pair:         "XMR/BTC",
			Amount:       12.3,
			Quantity:     2.9,
			Price:        3.8,
			Side:         2,
			SideMsg:      "side 2",
			Message:      "message 6",
			Status:       0,
			StatusMsg:    "status 1",
			SumBuy:       120,
			ApiKeyID:     6000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           7,
			UUID:         "uuid7",
			OrderID:      123,
			UserID:       456,
			ExchangeID:   789,
			UnitedOrders: 0,
			OrderType:    1,
			OrderTypeMsg: "order type 1",
			Pair:         "BTC/USD",
			Amount:       15.6,
			Quantity:     3.1,
			Price:        5.0,
			Side:         1,
			SideMsg:      "side 1",
			Message:      "message 7",
			Status:       1,
			StatusMsg:    "status 2",
			SumBuy:       150,
			ApiKeyID:     7000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
		{
			ID:           8,
			UUID:         "uuid8",
			OrderID:      456,
			UserID:       789,
			ExchangeID:   123,
			UnitedOrders: 1,
			OrderType:    2,
			OrderTypeMsg: "order type 2",
			Pair:         "ETH/BTC",
			Amount:       18.9,
			Quantity:     2.6,
			Price:        4.7,
			Side:         2,
			SideMsg:      "side 2",
			Message:      "message 8",
			Status:       0,
			StatusMsg:    "status 1",
			SumBuy:       180,
			ApiKeyID:     8000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    types.NullTime{},
		},
	}

	// Создание коллекции ExchangeOrderCollection на основе списка orders.
	collection := NewExchangeOrderCollection(orders)

	// Добавление нового элемента в коллекцию.
	collection = collection.Push(&ExchangeOrder{
		ID:           8,
		UUID:         "uuid9",
		OrderID:      123,
		UserID:       456,
		ExchangeID:   789,
		UnitedOrders: 0,
		OrderType:    1,
		OrderTypeMsg: "order type 9",
		Pair:         "BTC/USD",
		Amount:       10.5,
		Quantity:     2.5,
		Price:        4.2,
		Side:         1,
		SideMsg:      "side 9",
		Message:      "message 9",
		Status:       0,
		StatusMsg:    "status 9",
		SumBuy:       100,
		ApiKeyID:     1000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    types.NullTime{},
	})

	// Получение количества элементов в коллекции.
	fmt.Println(collection.Len()) // == 9

	// Поиск элемента в коллекции, удовлетворяющего заданному условию.
	fmt.Println(collection.Find(func(item *ExchangeOrder) bool {
		return item.ID == 7
	}))

	// Преобразование коллекции в список целых чисел на основе свойства SumBuy каждого элемента.
	fmt.Println(collection.MapToInt(func(item *ExchangeOrder) int {
		return item.SumBuy
	}))

	// Преобразование коллекции в список строк на основе свойства Pair каждого элемента.
	fmt.Println(collection.MapToString(func(item *ExchangeOrder) string {
		return item.Pair
	}))

	// Получение количества уникальных элементов в коллекции по идентификатору.
	fmt.Println(collection.UniqByID().Len()) // == 8
}
