package generator

import (
	"fmt"
	"reflect"
	"testing"
)

type ExchangeOrderTest struct {
	ID   int
	Name string
}

func TestCollections_Filter(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			callback func(item *ExchangeOrderTest) bool
		}
		want []*ExchangeOrderTest
	}{
		{
			name: "Filtering by ID",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) bool
			}{
				callback: func(item *ExchangeOrderTest) bool {
					return item.ID == 2 || item.ID == 3
				},
			},
			want: []*ExchangeOrderTest{
				{ID: 2, Name: "Order 2"},
				{ID: 3, Name: "Order 3"},
			},
		},
		{
			name: "Filtering by Name",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) bool
			}{
				callback: func(item *ExchangeOrderTest) bool {
					return item.Name == "Order 1" || item.Name == "Order 3"
				},
			},
			want: []*ExchangeOrderTest{
				{ID: 1, Name: "Order 1"},
				{ID: 3, Name: "Order 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Filter(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Find(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			callback func(item *ExchangeOrderTest) bool
		}
		want *ExchangeOrderTest
	}{
		{
			name: "Find by ID",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) bool
			}{
				callback: func(item *ExchangeOrderTest) bool {
					return item.ID == 2
				},
			},
			want: &ExchangeOrderTest{ID: 2, Name: "Order 2"},
		},
		{
			name: "Find by Name",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) bool
			}{
				callback: func(item *ExchangeOrderTest) bool {
					return item.Name == "Order 3"
				},
			},
			want: &ExchangeOrderTest{ID: 3, Name: "Order 3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Find(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Get(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			index int
		}
		wantModel *ExchangeOrderTest
	}{
		{
			name: "Get by index",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				index int
			}{
				index: 1,
			},
			wantModel: &ExchangeOrderTest{ID: 2, Name: "Order 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotModel := tt.c.Get(tt.args.index); !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("Get() = %v, want %v", gotModel, tt.wantModel)
			}
		})
	}
}

func TestCollections_Len(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		want int
	}{
		{
			name: "Len of empty collection",
			c:    Collections[ExchangeOrderTest]{},
			want: 0,
		},
		{
			name: "Len of non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_MapToInt(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			callback func(item *ExchangeOrderTest) int
		}
		want []int
	}{
		{
			name: "MapToInt on empty collection",
			c:    Collections[ExchangeOrderTest]{},
			args: struct {
				callback func(item *ExchangeOrderTest) int
			}{
				callback: func(item *ExchangeOrderTest) int {
					return item.ID
				},
			},
			want: nil,
		},
		{
			name: "MapToInt on non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) int
			}{
				callback: func(item *ExchangeOrderTest) int {
					return item.ID * 2
				},
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MapToInt(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_MapToString(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			callback func(item *ExchangeOrderTest) string
		}
		want []string
	}{
		{
			name: "MapToString on empty collection",
			c:    Collections[ExchangeOrderTest]{},
			args: struct {
				callback func(item *ExchangeOrderTest) string
			}{
				callback: func(item *ExchangeOrderTest) string {
					return item.Name
				},
			},
			want: []string{},
		},
		{
			name: "MapToString on non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			args: struct {
				callback func(item *ExchangeOrderTest) string
			}{
				callback: func(item *ExchangeOrderTest) string {
					return fmt.Sprintf("Order %d", item.ID)
				},
			},
			want: []string{"Order 1", "Order 2", "Order 3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MapToString(tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Pop(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		want *ExchangeOrderTest
	}{
		{
			name: "Pop on non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 3, Name: "Order 3"},
				},
			},
			want: &ExchangeOrderTest{ID: 3, Name: "Order 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Push(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args ExchangeOrderTest
		want *Collections[ExchangeOrderTest]
	}{
		{
			name: "Push on empty collection",
			c:    Collections[ExchangeOrderTest]{},
			args: ExchangeOrderTest{ID: 1, Name: "Order 1"},
			want: &Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{{ID: 1, Name: "Order 1"}},
			},
		},
		{
			name: "Push on non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
				},
			},
			args: ExchangeOrderTest{ID: 2, Name: "Order 2"},
			want: &Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Push(&tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Shift(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		want *ExchangeOrderTest
	}{
		{
			name: "Shift on non-empty collection",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
			want: &ExchangeOrderTest{ID: 1, Name: "Order 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_SortByField(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			getField func(item *ExchangeOrderTest) interface{}
		}
		mode string
		want *Collections[ExchangeOrderTest]
	}{
		{
			name: "UniqByField",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 1, Name: "Order 1"},
				},
			},
			mode: "",
			args: struct {
				getField func(item *ExchangeOrderTest) interface{}
			}{
				getField: func(item *ExchangeOrderTest) interface{} {
					return item.ID
				},
			},
			want: &Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SortByField(tt.args.getField, tt.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_UniqByField(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			getField func(item *ExchangeOrderTest) interface{}
			mode     string
		}
		want *Collections[ExchangeOrderTest]
	}{
		{
			name: "Sort in ascending order",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 3, Name: "Order 3"},
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
					{ID: 2, Name: "Order 2"},
				},
			},
			args: struct {
				getField func(item *ExchangeOrderTest) interface{}
				mode     string
			}{
				getField: func(item *ExchangeOrderTest) interface{} {
					return item.ID
				},
				mode: "asc",
			},
			want: &Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 3, Name: "Order 3"},
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.UniqByField(tt.args.getField); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollections_Unshift(t *testing.T) {
	tests := []struct {
		name string
		c    Collections[ExchangeOrderTest]
		args struct {
			item *ExchangeOrderTest
		}
		want []*ExchangeOrderTest
	}{
		{
			name: "Unshift",
			c: Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
			args: struct {
				item *ExchangeOrderTest
			}{
				item: &ExchangeOrderTest{ID: 3, Name: "Order 3"},
			},
			want: []*ExchangeOrderTest{
				{ID: 3, Name: "Order 3"},
				{ID: 1, Name: "Order 1"},
				{ID: 2, Name: "Order 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Unshift(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unshift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCollection(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			items []*ExchangeOrderTest
		}
		want *Collections[ExchangeOrderTest]
	}{
		{
			name: "NewCollection",
			args: struct {
				items []*ExchangeOrderTest
			}{
				items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
			want: &Collections[ExchangeOrderTest]{
				Items: []*ExchangeOrderTest{
					{ID: 1, Name: "Order 1"},
					{ID: 2, Name: "Order 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCollection(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmptyCollection(t *testing.T) {
	tests := []struct {
		name string
		want *Collections[ExchangeOrderTest]
	}{
		{
			name: "NewEmptyCollection",
			want: &Collections[ExchangeOrderTest]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyCollection[ExchangeOrderTest](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
