package main

import (
	"reflect"
	"testing"
)

func TestNewExchangeOrderCollection(t *testing.T) {
	type args struct {
		items []*ExchangeOrder
	}
	testCase1 := struct {
		name string
		args args
		want *ExchangeOrderCollection
	}{
		name: "EmptyItems",
		args: args{
			items: []*ExchangeOrder{},
		},
		want: &ExchangeOrderCollection{
			Items: []*ExchangeOrder{},
		},
	}
	order := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	testCase2 := struct {
		name string
		args args
		want *ExchangeOrderCollection
	}{
		name: "SingleItem",
		args: args{
			items: []*ExchangeOrder{order},
		},
		want: &ExchangeOrderCollection{
			Items: []*ExchangeOrder{order},
		},
	}
	order1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	order2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	order3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	testCase3 := struct {
		name string
		args args
		want *ExchangeOrderCollection
	}{
		name: "MultipleItems",
		args: args{
			items: []*ExchangeOrder{order1, order2, order3},
		},
		want: &ExchangeOrderCollection{
			Items: []*ExchangeOrder{order1, order2, order3},
		},
	}
	tests := []struct {
		name string
		args args
		want *ExchangeOrderCollection
	}{
		testCase1,
		testCase2,
		testCase3,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExchangeOrderCollection(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExchangeOrderCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Filter(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		callback SearchCallbackExchangeOrder
	}
	// Prepare test data
	order1 := &ExchangeOrder{ID: 1, Pair: "BTC/USD", Status: 1}
	order2 := &ExchangeOrder{ID: 2, Pair: "ETH/USD", Status: 2}
	order3 := &ExchangeOrder{ID: 3, Pair: "BTC/USD", Status: 1}
	orders := []*ExchangeOrder{order1, order2, order3}

	// Define the callback function for filtering
	callback := func(order *ExchangeOrder) bool {
		return order.Status == 1 && order.Pair == "BTC/USD"
	}

	// Define the expected result
	expectedResult := &ExchangeOrderCollection{
		Items: []*ExchangeOrder{order1, order3},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name:   "Filter by status and pair",
			fields: fields{Items: orders},
			args:   args{callback: callback},
			want:   expectedResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Filter(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Find(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		callback SearchCallbackExchangeOrder
	}
	order1 := &ExchangeOrder{ID: 1, UserID: 100, Pair: "BTC/USD", Status: 1}
	order2 := &ExchangeOrder{ID: 2, UserID: 200, Pair: "ETH/USD", Status: 2}
	order3 := &ExchangeOrder{ID: 3, UserID: 100, Pair: "BTC/ETH", Status: 1}
	order4 := &ExchangeOrder{ID: 4, UserID: 300, Pair: "LTC/USD", Status: 3}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrder
	}{
		{
			name: "Find existing order by ID",
			fields: fields{
				Items: []*ExchangeOrder{order1, order2, order3, order4},
			},
			args: args{
				callback: func(order *ExchangeOrder) bool {
					return order.ID == 2
				},
			},
			want: order2,
		},
		{
			name: "Find existing order by UserID",
			fields: fields{
				Items: []*ExchangeOrder{order1, order2, order3, order4},
			},
			args: args{
				callback: func(order *ExchangeOrder) bool {
					return order.UserID == 100
				},
			},
			want: order1,
		},
		{
			name: "Find non-existing order",
			fields: fields{
				Items: []*ExchangeOrder{order1, order2, order3},
			},
			args: args{
				callback: func(order *ExchangeOrder) bool {
					return order.ID == 4
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Find(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_ForEach(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		callback func(item *ExchangeOrder)
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001, Status: 1}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002, Status: 2}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003, Status: 1}
	var calledCount int
	callback := func(item *ExchangeOrder) {
		calledCount++
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Callback called for each item in the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				callback: callback,
			},
		},
		{
			name: "Callback not called when the collection is empty",
			fields: fields{
				Items: []*ExchangeOrder{},
			},
			args: args{
				callback: callback,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			c.ForEach(tt.args.callback)
		})
	}
}

func TestExchangeOrderCollection_Get(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		index int
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantModel *ExchangeOrder
	}{
		{
			name: "Valid index, should return the correct ExchangeOrder",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				index: 1,
			},
			wantModel: exchangeOrder2,
		},
		{
			name: "Index out of range, should return nil",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				index: 5,
			},
			wantModel: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if gotModel := c.Get(tt.args.index); !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("Get() = %v, want %v", gotModel, tt.wantModel)
			}
		})
	}
}

func TestExchangeOrderCollection_Len(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Returns the correct length of the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			want: 3,
		},
		{
			name: "Returns 0 for an empty collection",
			fields: fields{
				Items: []*ExchangeOrder{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_MapToInt(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		callback func(item *ExchangeOrder) int
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "Returns the correct length of the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				callback: func(item *ExchangeOrder) int {
					return int(item.OrderID)
				},
			},
			want: []int{1001, 1002, 1003},
		},
		{
			name: "Returns 0 for an empty collection",
			fields: fields{
				Items: []*ExchangeOrder{},
			},
			args: args{
				callback: func(item *ExchangeOrder) int {
					return int(item.OrderID)
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.MapToInt(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_MapToString(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		callback func(item *ExchangeOrder) string
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001, Pair: "BTC/USD"}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002, Pair: "ETH/BTC"}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003, Pair: "XRP/USD"}

	callback := func(item *ExchangeOrder) string {
		return item.Pair
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "MapToString applies the callback function to each item",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				callback: callback,
			},
			want: []string{"BTC/USD", "ETH/BTC", "XRP/USD"},
		},
		{
			name: "MapToString returns an empty slice for an empty collection",
			fields: fields{
				Items: []*ExchangeOrder{},
			},
			args: args{
				callback: callback,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.MapToString(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Pop(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name   string
		fields fields
		want   *ExchangeOrder
	}{
		{
			name: "Pop removes and returns the last item in the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			want: exchangeOrder3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Push(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		item *ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name: "Push adds an item to the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1},
			},
			args: args{
				item: exchangeOrder2,
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Push(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Shift(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	exchangeOrder3 := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name   string
		fields fields
		want   *ExchangeOrder
	}{
		{
			name: "Shift removes and returns the first item in the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			want: exchangeOrder1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_SortByAmount(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		mode string
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, Amount: 10.0}
	exchangeOrder2 := &ExchangeOrder{ID: 2, Amount: 5.0}
	exchangeOrder3 := &ExchangeOrder{ID: 3, Amount: 7.5}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name: "SortByAmount sorts the collection in ascending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "asc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder2, exchangeOrder3, exchangeOrder1},
			},
		},
		{
			name: "SortByAmount sorts the collection in descending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "desc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder3, exchangeOrder2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.SortByAmount(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_SortByApiKeyID(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		mode string
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, ApiKeyID: 100}
	exchangeOrder2 := &ExchangeOrder{ID: 2, ApiKeyID: 200}
	exchangeOrder3 := &ExchangeOrder{ID: 3, ApiKeyID: 150}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name: "SortByApiKeyID sorts the collection in ascending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "asc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder3, exchangeOrder2},
			},
		},
		{
			name: "SortByApiKeyID sorts the collection in descending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "desc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder2, exchangeOrder3, exchangeOrder1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.SortByApiKeyID(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByApiKeyID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_SortByExchangeID(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		mode string
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, ExchangeID: 100}
	exchangeOrder2 := &ExchangeOrder{ID: 2, ExchangeID: 200}
	exchangeOrder3 := &ExchangeOrder{ID: 3, ExchangeID: 150}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name: "SortByExchangeID sorts the collection in ascending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "asc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder3, exchangeOrder2},
			},
		},
		{
			name: "SortByExchangeID sorts the collection in descending order",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			args: args{
				mode: "desc",
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder2, exchangeOrder3, exchangeOrder1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.SortByExchangeID(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByExchangeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_UniqByAmount(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, Amount: 10.0}
	exchangeOrder2 := &ExchangeOrder{ID: 2, Amount: 5.0}
	exchangeOrder3 := &ExchangeOrder{ID: 3, Amount: 10.0}
	exchangeOrder4 := &ExchangeOrder{ID: 4, Amount: 7.5}
	tests := []struct {
		name   string
		fields fields
		want   *ExchangeOrderCollection
	}{
		{
			name: "UniqByAmount removes duplicate items with the same amount",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3, exchangeOrder4},
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder4},
			},
		},
		{
			name: "UniqByAmount returns the same collection if there are no duplicate amounts",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder4},
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.UniqByAmount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqByAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_UniqByApiKeyID(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, ApiKeyID: 100}
	exchangeOrder2 := &ExchangeOrder{ID: 2, ApiKeyID: 200}
	exchangeOrder3 := &ExchangeOrder{ID: 3, ApiKeyID: 100}
	tests := []struct {
		name   string
		fields fields
		want   *ExchangeOrderCollection
	}{
		{
			name: "UniqByApiKeyID removes duplicate items with the same ApiKeyID",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2, exchangeOrder3},
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2},
			},
		},
		{
			name: "UniqByApiKeyID returns the same collection if there are no duplicate ApiKeyIDs",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2},
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.UniqByApiKeyID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqByApiKeyID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeOrderCollection_Unshift(t *testing.T) {
	type fields struct {
		Items []*ExchangeOrder
	}
	type args struct {
		item *ExchangeOrder
	}
	exchangeOrder1 := &ExchangeOrder{ID: 1, UUID: "abc123", OrderID: 1001}
	exchangeOrder2 := &ExchangeOrder{ID: 2, UUID: "def456", OrderID: 1002}
	newExchangeOrder := &ExchangeOrder{ID: 3, UUID: "ghi789", OrderID: 1003}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ExchangeOrderCollection
	}{
		{
			name: "Unshift adds an item to the beginning of the collection",
			fields: fields{
				Items: []*ExchangeOrder{exchangeOrder1, exchangeOrder2},
			},
			args: args{
				item: newExchangeOrder,
			},
			want: &ExchangeOrderCollection{
				Items: []*ExchangeOrder{newExchangeOrder, exchangeOrder1, exchangeOrder2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExchangeOrderCollection{
				Items: tt.fields.Items,
			}
			if got := c.Unshift(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unshift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmptyExchangeOrderCollection(t *testing.T) {
	tests := []struct {
		name string
		want *ExchangeOrderCollection
	}{
		{
			name: "Create NewEmptyExchangeOrderCollection",
			want: &ExchangeOrderCollection{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyExchangeOrderCollection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyExchangeOrderCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
