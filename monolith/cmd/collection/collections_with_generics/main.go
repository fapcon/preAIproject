package main

import (
	"fmt"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/cmd/collection/collections_with_generics/generator"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/cmd/collection/collections_with_generics/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

func main() {
	col1 := generator.NewCollection([]*models.ExchangeOrderTest{
		{
			ID:           7,
			UUID:         "uuid7",
			OrderID:      2,
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
			OrderID:      4,
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
		{
			ID:           1,
			UUID:         "uuid1",
			OrderID:      6,
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
			OrderID:      8,
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
	})
	col2 := generator.NewCollection([]*models.StrategyTest{
		{
			ID:         1,
			ExchangeID: 123,
			UserID:     987,
			APIKey:     "my-api-key",
			SecretKey:  "my-secret-key",
		},
		{
			ID:         2,
			ExchangeID: 789,
			UserID:     987,
			APIKey:     "another-api-key",
			SecretKey:  "another-secret-key",
		},
		{
			ID:         3,
			ExchangeID: 456,
			UserID:     789,
			APIKey:     "third-api-key",
			SecretKey:  "third-secret-key",
		},
	})

	fmt.Println(col1)
	fmt.Println(col2)

	col2 = col2.Push(&models.StrategyTest{
		ID:         4,
		ExchangeID: 4,
		UserID:     456,
		APIKey:     "my-api-key",
		SecretKey:  "my-secret-key",
	})

	fmt.Println(col2.Len())

	fmt.Println(col2.MapToString(func(item *models.StrategyTest) string {
		return item.APIKey
	}))

	fmt.Println(col2.MapToInt(func(item *models.StrategyTest) int {
		return item.ID
	}))

	fmt.Println(col2.Find(func(item *models.StrategyTest) bool {
		return item.ID == 1
	}))

	col3 := col2.Filter(func(item *models.StrategyTest) bool {
		return item.UserID == 987
	})
	for _, item := range col3 {
		value := *item
		fmt.Println(value)
	}

	fmt.Println(col2.Get(1))

	fmt.Println(col2.Shift())

	fmt.Println(col2.Pop())

	col1 = col1.Push(&models.ExchangeOrderTest{
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

	col4 := col1.UniqByField(func(item *models.ExchangeOrderTest) interface{} {
		return item.ID
	})
	for _, item := range col4.Items {
		value := *item
		fmt.Println(value)
	}
	fmt.Println(col1.Len())

	col5 := col1.SortByField(func(item *models.ExchangeOrderTest) interface{} {
		return int(item.OrderID)
	}, "")
	for _, item := range col5.Items {
		value := *item
		fmt.Println(value)
	}
}
