package ezarr

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type List struct {
	Elements []interface{}
}

func New(elements ...interface{}) *List {
	return &List{Elements: elements}
}

func (l *List) Append(element interface{}) *List {
	l.Elements = append(l.Elements, element)
	return l
}

func (l *List) Extend(other *List) *List {
	l.Elements = append(l.Elements, other.Elements...)
	return l
}

func (l *List) Insert(index int, element interface{}) *List {
	if index < 0 {
		index = len(l.Elements) + index
		if index < 0 {
			index = 0
		}
	}
	if index > len(l.Elements) {
		index = len(l.Elements)
	}

	l.Elements = append(l.Elements[:index], append([]interface{}{element}, l.Elements[index:]...)...)
	return l
}

func (l *List) Remove(element interface{}) error {
	index := l.Index(element)
	if index == -1 {
		return fmt.Errorf("element %v not found in list", element)
	}
	l.Elements = append(l.Elements[:index], l.Elements[index+1:]...)
	return nil
}

func (l *List) Pop(index int) (interface{}, error) {
	if len(l.Elements) == 0 {
		return nil, fmt.Errorf("cannot pop from empty list")
	}

	if index < 0 {
		index = len(l.Elements) + index
	}

	if index < 0 || index >= len(l.Elements) {
		return nil, fmt.Errorf("index %d out of range", index)
	}

	element := l.Elements[index]
	l.Elements = append(l.Elements[:index], l.Elements[index+1:]...)
	return element, nil
}

func (l *List) Index(element interface{}) int {
	for i, e := range l.Elements {
		if reflect.DeepEqual(e, element) {
			return i
		}
	}
	return -1
}

func (l *List) Count(element interface{}) int {
	count := 0
	for _, e := range l.Elements {
		if reflect.DeepEqual(e, element) {
			count++
		}
	}
	return count
}

func (l *List) Sort() error {
	if len(l.Elements) == 0 {
		return nil
	}

	firstType := reflect.TypeOf(l.Elements[0])
	for _, e := range l.Elements {
		if reflect.TypeOf(e) != firstType {
			return fmt.Errorf("cannot sort list with mixed types")
		}
	}

	switch l.Elements[0].(type) {
	case int:
		sort.Slice(l.Elements, func(i, j int) bool {
			return l.Elements[i].(int) < l.Elements[j].(int)
		})
	case float64:
		sort.Slice(l.Elements, func(i, j int) bool {
			return l.Elements[i].(float64) < l.Elements[j].(float64)
		})
	case string:
		sort.Slice(l.Elements, func(i, j int) bool {
			return l.Elements[i].(string) < l.Elements[j].(string)
		})
	default:
		return fmt.Errorf("cannot sort elements of type %s", firstType.String())
	}

	return nil
}

func (l *List) Reverse() *List {
	for i, j := 0, len(l.Elements)-1; i < j; i, j = i+1, j-1 {
		l.Elements[i], l.Elements[j] = l.Elements[j], l.Elements[i]
	}
	return l
}

func (l *List) Slice(start, end int) *List {
	len := len(l.Elements)

	if start < 0 {
		start = len + start
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end = len + end
		if end < 0 {
			end = 0
		}
	}

	if start > len {
		start = len
	}
	if end > len {
		end = len
	}
	if start > end {
		start, end = end, start
	}

	return &List{Elements: append([]interface{}{}, l.Elements[start:end]...)}
}

func (l *List) Copy() *List {
	newElements := make([]interface{}, len(l.Elements))
	copy(newElements, l.Elements)
	return &List{Elements: newElements}
}

func (l *List) Len() int {
	return len(l.Elements)
}

func (l *List) Clear() *List {
	l.Elements = []interface{}{}
	return l
}

func (l *List) String() string {
	strElems := make([]string, len(l.Elements))
	for i, e := range l.Elements {
		strElems[i] = fmt.Sprintf("%v", e)
	}
	return "[" + strings.Join(strElems, ", ") + "]"
}
