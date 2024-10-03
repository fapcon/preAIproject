package generator

import "sort"

type Value interface{}

type Collections[T Value] struct {
	Items []*T
}

func NewCollection[T Value](items []*T) *Collections[T] {
	var collection Collections[T]

	collection.Items = items

	return &collection
}

func NewEmptyCollection[T Value]() *Collections[T] {
	return &Collections[T]{}
}

func (c *Collections[T]) Find(callback func(item *T) bool) *T {
	for _, v := range c.Items {
		if callback(v) {
			return v
		}
	}
	return nil
}

func (c *Collections[T]) Filter(callback func(item *T) bool) []*T {
	var newItems []*T

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return newItems
}

func (c *Collections[T]) MapToInt(callback func(item *T) int) []int {
	var items []int

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *Collections[T]) MapToString(callback func(item *T) string) []string {
	items := []string{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *Collections[T]) Push(item *T) *Collections[T] {
	newItems := append(c.Items, item)

	return &Collections[T]{Items: newItems}
}

func (c *Collections[T]) Shift() *T {
	item := c.Items[0]
	c.Items = c.Items[1:]

	return item
}

func (c *Collections[T]) Pop() *T {
	item := c.Items[len(c.Items)-1]
	c.Items = c.Items[:len(c.Items)-1]

	return item
}

func (c *Collections[T]) Unshift(item *T) []*T {
	newItems := append([]*T{item}, c.Items...)

	return newItems
}

func (c *Collections[T]) Len() int {
	return len(c.Items)
}

func (c *Collections[T]) Get(index int) (model *T) {
	if index >= 0 && len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

func (c *Collections[T]) UniqByField(getField func(item *T) interface{}) *Collections[T] {
	collection := &Collections[T]{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *T) bool {
			return getField(model) == getField(item)
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *Collections[T]) SortByField(getField func(item *T) interface{}, mode string) *Collections[T] {
	collection := &Collections[T]{Items: c.Items}

	switch mode {
	case "desc":
		sort.Slice(collection.Items, func(i, j int) bool {
			value1 := getField(collection.Items[i])
			value2 := getField(collection.Items[j])

			switch v1 := value1.(type) {
			case int:
				return v1 > value2.(int)
			case string:
				return v1 > value2.(string)
			case float64:
				return v1 > value2.(float64)
			default:
				// Обработка ошибки или возврат false, если тип не поддерживается
				return false
			}
		})
	default:
		sort.Slice(collection.Items, func(i, j int) bool {
			value1 := getField(collection.Items[i])
			value2 := getField(collection.Items[j])

			switch v1 := value1.(type) {
			case int:
				return v1 < value2.(int)
			case string:
				return v1 < value2.(string)
			case float64:
				return v1 < value2.(float64)
			default:
				// Обработка ошибки или возврат false, если тип не поддерживается
				return false
			}
		})
	}

	return collection
}
