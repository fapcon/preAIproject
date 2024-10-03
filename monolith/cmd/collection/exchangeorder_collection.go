package main

import "sort"

type ExchangeOrderCollection struct {
	Items []*ExchangeOrder
}

type SearchCallbackExchangeOrder func(item *ExchangeOrder) bool

//grizzly:replaceName New{{.Name}}Collection
func NewExchangeOrderCollection(items []*ExchangeOrder) *ExchangeOrderCollection {
	var collection ExchangeOrderCollection

	collection.Items = items

	return &collection
}

//grizzly:replaceName NewEmpty{{.Name}}Collection
func NewEmptyExchangeOrderCollection() *ExchangeOrderCollection {
	return &ExchangeOrderCollection{}
}

func (c *ExchangeOrderCollection) Find(callback SearchCallbackExchangeOrder) *ExchangeOrder {
	for _, v := range c.Items {
		if callback(v) == true {
			return v
		}
	}

	return nil
}

func (c *ExchangeOrderCollection) Filter(callback SearchCallbackExchangeOrder) *ExchangeOrderCollection {
	var newItems []*ExchangeOrder

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return &ExchangeOrderCollection{Items: newItems}
}

func (c *ExchangeOrderCollection) MapToInt(callback func(item *ExchangeOrder) int) []int {
	items := []int{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *ExchangeOrderCollection) MapToString(callback func(item *ExchangeOrder) string) []string {
	items := []string{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *ExchangeOrderCollection) Push(item *ExchangeOrder) *ExchangeOrderCollection {
	newItems := append(c.Items, item)

	return &ExchangeOrderCollection{Items: newItems}
}

func (c *ExchangeOrderCollection) Shift() *ExchangeOrder {
	item := c.Items[0]
	c.Items = c.Items[1:]

	return item
}

func (c *ExchangeOrderCollection) Pop() *ExchangeOrder {
	item := c.Items[len(c.Items)-1]
	c.Items = c.Items[:len(c.Items)-1]

	return item
}

func (c *ExchangeOrderCollection) Unshift(item *ExchangeOrder) *ExchangeOrderCollection {
	newItems := append([]*ExchangeOrder{item}, c.Items...)

	return &ExchangeOrderCollection{Items: newItems}
}

func (c *ExchangeOrderCollection) Len() int {
	return len(c.Items)
}

func (c *ExchangeOrderCollection) Get(index int) (model *ExchangeOrder) {
	if index >= 0 && len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

func (c *ExchangeOrderCollection) UniqByMessage() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Message == item.Message
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByOrderID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.OrderID == item.OrderID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByOrderType() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.OrderType == item.OrderType
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByOrderTypeMsg() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.OrderTypeMsg == item.OrderTypeMsg
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByQuantity() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Quantity == item.Quantity
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByUnitedOrders() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.UnitedOrders == item.UnitedOrders
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByAmount() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Amount == item.Amount
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByPrice() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Price == item.Price
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByApiKeyID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.ApiKeyID == item.ApiKeyID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByUUID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.UUID == item.UUID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByPair() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Pair == item.Pair
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByStatus() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Status == item.Status
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqBySideMsg() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.SideMsg == item.SideMsg
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByStatusMsg() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.StatusMsg == item.StatusMsg
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqBySumBuy() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.SumBuy == item.SumBuy
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.ID == item.ID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByUserID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.UserID == item.UserID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqByExchangeID() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.ExchangeID == item.ExchangeID
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *ExchangeOrderCollection) UniqBySide() *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *ExchangeOrder) bool {
			return model.Side == item.Side
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

type byMessageAsc []*ExchangeOrder

func (a byMessageAsc) Len() int           { return len(a) }
func (a byMessageAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMessageAsc) Less(i, j int) bool { return a[i].Message < a[j].Message }

type byMessageDesc []*ExchangeOrder

func (a byMessageDesc) Len() int           { return len(a) }
func (a byMessageDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMessageDesc) Less(i, j int) bool { return a[i].Message > a[j].Message }

func (c *ExchangeOrderCollection) SortByMessage(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byMessageDesc(collection.Items))
	} else {
		sort.Sort(byMessageAsc(collection.Items))
	}

	return collection
}

type byOrderIDAsc []*ExchangeOrder

func (a byOrderIDAsc) Len() int           { return len(a) }
func (a byOrderIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderIDAsc) Less(i, j int) bool { return a[i].OrderID < a[j].OrderID }

type byOrderIDDesc []*ExchangeOrder

func (a byOrderIDDesc) Len() int           { return len(a) }
func (a byOrderIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderIDDesc) Less(i, j int) bool { return a[i].OrderID > a[j].OrderID }

func (c *ExchangeOrderCollection) SortByOrderID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byOrderIDDesc(collection.Items))
	} else {
		sort.Sort(byOrderIDAsc(collection.Items))
	}

	return collection
}

type byOrderTypeAsc []*ExchangeOrder

func (a byOrderTypeAsc) Len() int           { return len(a) }
func (a byOrderTypeAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderTypeAsc) Less(i, j int) bool { return a[i].OrderType < a[j].OrderType }

type byOrderTypeDesc []*ExchangeOrder

func (a byOrderTypeDesc) Len() int           { return len(a) }
func (a byOrderTypeDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderTypeDesc) Less(i, j int) bool { return a[i].OrderType > a[j].OrderType }

func (c *ExchangeOrderCollection) SortByOrderType(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byOrderTypeDesc(collection.Items))
	} else {
		sort.Sort(byOrderTypeAsc(collection.Items))
	}

	return collection
}

type byOrderTypeMsgAsc []*ExchangeOrder

func (a byOrderTypeMsgAsc) Len() int           { return len(a) }
func (a byOrderTypeMsgAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderTypeMsgAsc) Less(i, j int) bool { return a[i].OrderTypeMsg < a[j].OrderTypeMsg }

type byOrderTypeMsgDesc []*ExchangeOrder

func (a byOrderTypeMsgDesc) Len() int           { return len(a) }
func (a byOrderTypeMsgDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOrderTypeMsgDesc) Less(i, j int) bool { return a[i].OrderTypeMsg > a[j].OrderTypeMsg }

func (c *ExchangeOrderCollection) SortByOrderTypeMsg(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byOrderTypeMsgDesc(collection.Items))
	} else {
		sort.Sort(byOrderTypeMsgAsc(collection.Items))
	}

	return collection
}

type byQuantityAsc []*ExchangeOrder

func (a byQuantityAsc) Len() int           { return len(a) }
func (a byQuantityAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byQuantityAsc) Less(i, j int) bool { return a[i].Quantity < a[j].Quantity }

type byQuantityDesc []*ExchangeOrder

func (a byQuantityDesc) Len() int           { return len(a) }
func (a byQuantityDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byQuantityDesc) Less(i, j int) bool { return a[i].Quantity > a[j].Quantity }

func (c *ExchangeOrderCollection) SortByQuantity(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byQuantityDesc(collection.Items))
	} else {
		sort.Sort(byQuantityAsc(collection.Items))
	}

	return collection
}

type byUnitedOrdersAsc []*ExchangeOrder

func (a byUnitedOrdersAsc) Len() int           { return len(a) }
func (a byUnitedOrdersAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUnitedOrdersAsc) Less(i, j int) bool { return a[i].UnitedOrders < a[j].UnitedOrders }

type byUnitedOrdersDesc []*ExchangeOrder

func (a byUnitedOrdersDesc) Len() int           { return len(a) }
func (a byUnitedOrdersDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUnitedOrdersDesc) Less(i, j int) bool { return a[i].UnitedOrders > a[j].UnitedOrders }

func (c *ExchangeOrderCollection) SortByUnitedOrders(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byUnitedOrdersDesc(collection.Items))
	} else {
		sort.Sort(byUnitedOrdersAsc(collection.Items))
	}

	return collection
}

type byAmountAsc []*ExchangeOrder

func (a byAmountAsc) Len() int           { return len(a) }
func (a byAmountAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAmountAsc) Less(i, j int) bool { return a[i].Amount < a[j].Amount }

type byAmountDesc []*ExchangeOrder

func (a byAmountDesc) Len() int           { return len(a) }
func (a byAmountDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAmountDesc) Less(i, j int) bool { return a[i].Amount > a[j].Amount }

func (c *ExchangeOrderCollection) SortByAmount(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byAmountDesc(collection.Items))
	} else {
		sort.Sort(byAmountAsc(collection.Items))
	}

	return collection
}

type byPriceAsc []*ExchangeOrder

func (a byPriceAsc) Len() int           { return len(a) }
func (a byPriceAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPriceAsc) Less(i, j int) bool { return a[i].Price < a[j].Price }

type byPriceDesc []*ExchangeOrder

func (a byPriceDesc) Len() int           { return len(a) }
func (a byPriceDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPriceDesc) Less(i, j int) bool { return a[i].Price > a[j].Price }

func (c *ExchangeOrderCollection) SortByPrice(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byPriceDesc(collection.Items))
	} else {
		sort.Sort(byPriceAsc(collection.Items))
	}

	return collection
}

type byApiKeyIDAsc []*ExchangeOrder

func (a byApiKeyIDAsc) Len() int           { return len(a) }
func (a byApiKeyIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byApiKeyIDAsc) Less(i, j int) bool { return a[i].ApiKeyID < a[j].ApiKeyID }

type byApiKeyIDDesc []*ExchangeOrder

func (a byApiKeyIDDesc) Len() int           { return len(a) }
func (a byApiKeyIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byApiKeyIDDesc) Less(i, j int) bool { return a[i].ApiKeyID > a[j].ApiKeyID }

func (c *ExchangeOrderCollection) SortByApiKeyID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byApiKeyIDDesc(collection.Items))
	} else {
		sort.Sort(byApiKeyIDAsc(collection.Items))
	}

	return collection
}

type byUUIDAsc []*ExchangeOrder

func (a byUUIDAsc) Len() int           { return len(a) }
func (a byUUIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUUIDAsc) Less(i, j int) bool { return a[i].UUID < a[j].UUID }

type byUUIDDesc []*ExchangeOrder

func (a byUUIDDesc) Len() int           { return len(a) }
func (a byUUIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUUIDDesc) Less(i, j int) bool { return a[i].UUID > a[j].UUID }

func (c *ExchangeOrderCollection) SortByUUID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byUUIDDesc(collection.Items))
	} else {
		sort.Sort(byUUIDAsc(collection.Items))
	}

	return collection
}

type byPairAsc []*ExchangeOrder

func (a byPairAsc) Len() int           { return len(a) }
func (a byPairAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPairAsc) Less(i, j int) bool { return a[i].Pair < a[j].Pair }

type byPairDesc []*ExchangeOrder

func (a byPairDesc) Len() int           { return len(a) }
func (a byPairDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPairDesc) Less(i, j int) bool { return a[i].Pair > a[j].Pair }

func (c *ExchangeOrderCollection) SortByPair(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byPairDesc(collection.Items))
	} else {
		sort.Sort(byPairAsc(collection.Items))
	}

	return collection
}

type byStatusAsc []*ExchangeOrder

func (a byStatusAsc) Len() int           { return len(a) }
func (a byStatusAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStatusAsc) Less(i, j int) bool { return a[i].Status < a[j].Status }

type byStatusDesc []*ExchangeOrder

func (a byStatusDesc) Len() int           { return len(a) }
func (a byStatusDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStatusDesc) Less(i, j int) bool { return a[i].Status > a[j].Status }

func (c *ExchangeOrderCollection) SortByStatus(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byStatusDesc(collection.Items))
	} else {
		sort.Sort(byStatusAsc(collection.Items))
	}

	return collection
}

type bySideMsgAsc []*ExchangeOrder

func (a bySideMsgAsc) Len() int           { return len(a) }
func (a bySideMsgAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySideMsgAsc) Less(i, j int) bool { return a[i].SideMsg < a[j].SideMsg }

type bySideMsgDesc []*ExchangeOrder

func (a bySideMsgDesc) Len() int           { return len(a) }
func (a bySideMsgDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySideMsgDesc) Less(i, j int) bool { return a[i].SideMsg > a[j].SideMsg }

func (c *ExchangeOrderCollection) SortBySideMsg(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(bySideMsgDesc(collection.Items))
	} else {
		sort.Sort(bySideMsgAsc(collection.Items))
	}

	return collection
}

type byStatusMsgAsc []*ExchangeOrder

func (a byStatusMsgAsc) Len() int           { return len(a) }
func (a byStatusMsgAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStatusMsgAsc) Less(i, j int) bool { return a[i].StatusMsg < a[j].StatusMsg }

type byStatusMsgDesc []*ExchangeOrder

func (a byStatusMsgDesc) Len() int           { return len(a) }
func (a byStatusMsgDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStatusMsgDesc) Less(i, j int) bool { return a[i].StatusMsg > a[j].StatusMsg }

func (c *ExchangeOrderCollection) SortByStatusMsg(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byStatusMsgDesc(collection.Items))
	} else {
		sort.Sort(byStatusMsgAsc(collection.Items))
	}

	return collection
}

type bySumBuyAsc []*ExchangeOrder

func (a bySumBuyAsc) Len() int           { return len(a) }
func (a bySumBuyAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySumBuyAsc) Less(i, j int) bool { return a[i].SumBuy < a[j].SumBuy }

type bySumBuyDesc []*ExchangeOrder

func (a bySumBuyDesc) Len() int           { return len(a) }
func (a bySumBuyDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySumBuyDesc) Less(i, j int) bool { return a[i].SumBuy > a[j].SumBuy }

func (c *ExchangeOrderCollection) SortBySumBuy(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(bySumBuyDesc(collection.Items))
	} else {
		sort.Sort(bySumBuyAsc(collection.Items))
	}

	return collection
}

type byIDAsc []*ExchangeOrder

func (a byIDAsc) Len() int           { return len(a) }
func (a byIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIDAsc) Less(i, j int) bool { return a[i].ID < a[j].ID }

type byIDDesc []*ExchangeOrder

func (a byIDDesc) Len() int           { return len(a) }
func (a byIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIDDesc) Less(i, j int) bool { return a[i].ID > a[j].ID }

func (c *ExchangeOrderCollection) SortByID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byIDDesc(collection.Items))
	} else {
		sort.Sort(byIDAsc(collection.Items))
	}

	return collection
}

type byUserIDAsc []*ExchangeOrder

func (a byUserIDAsc) Len() int           { return len(a) }
func (a byUserIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUserIDAsc) Less(i, j int) bool { return a[i].UserID < a[j].UserID }

type byUserIDDesc []*ExchangeOrder

func (a byUserIDDesc) Len() int           { return len(a) }
func (a byUserIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUserIDDesc) Less(i, j int) bool { return a[i].UserID > a[j].UserID }

func (c *ExchangeOrderCollection) SortByUserID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byUserIDDesc(collection.Items))
	} else {
		sort.Sort(byUserIDAsc(collection.Items))
	}

	return collection
}

type byExchangeIDAsc []*ExchangeOrder

func (a byExchangeIDAsc) Len() int           { return len(a) }
func (a byExchangeIDAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byExchangeIDAsc) Less(i, j int) bool { return a[i].ExchangeID < a[j].ExchangeID }

type byExchangeIDDesc []*ExchangeOrder

func (a byExchangeIDDesc) Len() int           { return len(a) }
func (a byExchangeIDDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byExchangeIDDesc) Less(i, j int) bool { return a[i].ExchangeID > a[j].ExchangeID }

func (c *ExchangeOrderCollection) SortByExchangeID(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byExchangeIDDesc(collection.Items))
	} else {
		sort.Sort(byExchangeIDAsc(collection.Items))
	}

	return collection
}

type bySideAsc []*ExchangeOrder

func (a bySideAsc) Len() int           { return len(a) }
func (a bySideAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySideAsc) Less(i, j int) bool { return a[i].Side < a[j].Side }

type bySideDesc []*ExchangeOrder

func (a bySideDesc) Len() int           { return len(a) }
func (a bySideDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySideDesc) Less(i, j int) bool { return a[i].Side > a[j].Side }

func (c *ExchangeOrderCollection) SortBySide(mode string) *ExchangeOrderCollection {
	collection := &ExchangeOrderCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(bySideDesc(collection.Items))
	} else {
		sort.Sort(bySideAsc(collection.Items))
	}

	return collection
}

func (c *ExchangeOrderCollection) ForEach(callback func(item *ExchangeOrder)) {
	for _, i := range c.Items {
		callback(i)
	}
}
