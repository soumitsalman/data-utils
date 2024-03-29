package datautils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// data utilities
func Filter[T any](items []T, condition func(item *T) bool) []T {
	var new_items = make([]T, 0, len(items))
	for _, item := range items {
		if condition(&item) {
			new_items = append(new_items, item)
		}
	}
	return new_items
}

func ForEach[T any](items []T, do func(item *T)) []T {
	for i := range items {
		do(&items[i])
	}
	return items
}

func Transform[T_input, T_output any](items []T_input, transformer func(item *T_input) T_output) []T_output {
	var new_items = make([]T_output, len(items))
	for i := range items {
		new_items[i] = transformer(&items[i])
	}
	return new_items
}

func Reduce[T any](items []T, reduce func(a, b T) T) T {
	var res T
	for i := range items {
		res = reduce(res, items[i])
	}
	return res
}

func In[T any](item T, list []T, condition func(a, b *T) bool) bool {
	return Index[T](item, list, condition) >= 0
}

func Any[T any](list []T, condition func(item *T) bool) bool {
	return IndexAny[T](list, condition) >= 0
}

func Index[T any](item T, list []T, compare func(a, b *T) bool) int {
	for i := range list {
		if compare(&item, &list[i]) {
			return i
		}
	}
	return -1
}

func IndexAny[T any](list []T, condition func(item *T) bool) int {
	for i := range list {
		if condition(&list[i]) {
			return i
		}
	}
	return -1
}

func SafeSlice[T any](array []T, start, noninclusive_end int) []T {
	if start < 0 {
		start = 0
	}
	if noninclusive_end < 0 {
		noninclusive_end = 0
	}
	return array[min(start, len(array)):min(noninclusive_end, len(array))]
}

func MapToArray[TKey comparable, TValue any](list map[TKey]TValue) ([]TKey, []TValue) {
	keys := make([]TKey, 0, len(list))
	values := make([]TValue, 0, len(list))
	for key, val := range list {
		keys = append(keys, key)
		values = append(values, val)
	}
	return keys, values
}

func AppendMaps[TKey comparable, TValue any](to_map, from_map map[TKey]TValue) map[TKey]TValue {
	for key, val := range from_map {
		to_map[key] = val
	}
	return to_map
}

func PrintTable[T any](contents []T, headers []string, fields func(item *T) []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader(headers)
	ForEach[T](contents, func(item *T) {
		table.Append(fields(item))
	})
	table.Render()
}
