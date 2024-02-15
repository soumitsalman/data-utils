# data-utils
Utility function for on-compute data processing of generic arrays and maps

## Examples
Run the following to get the package

```
go get github.com/soumitsalman/data-utils
```

Use the following import
import (    
    // "utils" or whatever alias you want to use
    utils "github.com/soumitsalman/data-utils"
)

**Transform[T_input, T_output any](items []T_input, transformer func(item *T_input) T_output) []T_output**
This is equivalent to Python `map`. It can be used to convert an array of `T_input` to array of `T_output` or even extract a specific field from a struct.

```
type Message struct {
    Source string
    Id string
    PrettyName string
    // ... other fields
}

var messages []Message
// ... some instantiation of messages

sources := utils.Transform[Message, string](messages, func(item *Message) string { 
    return item.Source 
})
```

**Filter[T any](items []T, condition func(item *T) bool) []T**
This is equivalent to Python `filter`. It can be used to pick a sub-array based on the condition function.

```
unknown_source := utils.Filter[Message](engagements, func(item *Message) bool {
	return item.Source == ""
})
```

**Reduce[T any](items []T, reduce func(a, b T) T) T**
This is equivalent to Python `reduce`. It can be used to add or combine an array of type T into one T value.

```
sum_result := utils.Reduce[int](int, func(a, b int) int { 
    return a + b 
})
```

**Any[T any](list []T, condition func(item *T) bool) bool**
Looks through `list` to see if there is any item that matches the `condition`. If so it will return `true`. If there is no such item, it returns `false`.

```
utils.Any[Message](existing_cats, func(item *Message) bool 
{ 
    return item.Source == "My Trunk" && item.Id == "pile driver" 
})
```

**IndexAny[T any](list []T, condition func(item *T) bool) int**
Similar to Any, it looks through `list` to see if there is any item that matches the `condition`. If so it will return the index of the item in the array. If there is no such item, it returns -1.

```
utils.IndexAny[Message](existing_cats, func(item *Message) bool 
{ 
    return item.Source == "My Trunk" && item.Id == "pile driver" 
})
```

**ForEach[T any](items []T, do func(item *T)) []T**
A for-loop wrapper on `items` array. This can be used to `do` any function with and on each item in the array. `item *T` allows passing reference to the item in the array.

```
sources := make([]string, 0, len(messages))

utils.ForEach[Message](contents, func(item *Message) {
	sources = append(sources, item.Source)
    item.PrettyName = fmt.Sprintf("%s@%s", item.Id, item.Sources)
    // ... do other stuff
})
```

## Other functions in this package
* `In[T any](item T, list []T, condition func(a, b *T) bool) bool` --> Similar to Any except it tries to match `item` within `list` using `condition` function.
* `Index[T any](item T, list []T, compare func(a, b *T) bool) int` --> Similar to IndexAny except it tries to match `item` within `list` using `compare` function.
* `SafeSlice[T any](array []T, start, noninclusive_end int) []T` --> similar to `somearray[start:noninclusive_end]` but it normalizes the `start` and `noninclusive_end` to keep it contained within `0` to `len(somearray)`.
* `MapToArray[TKey comparable, TValue any](list map[TKey]TValue) ([]TKey, []TValue)` --> Returns two arrays: one of the keys and one of the values.
* `AppendMaps[TKey comparable, TValue any](to_map, from_map map[TKey]TValue) map[TKey]TValue` --> Inserts the values from `from_map` to `to_map`.
* `DateToString(time_val float64) string --> Represents a floating point value of time into `YYYY-MM-DD` format.
* `TruncateTextWithEllipsis(text string, max_len int) string` --> If `text` is longer than `max_len` then it crops the text to `max_len` and adds `...`. Or else it returns the text unmodified.


