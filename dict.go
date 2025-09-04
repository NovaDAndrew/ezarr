package ezarr

import (
	"fmt"
	"reflect"
	"strings"
)

type Dict struct {
	Keys   []interface{}
	Values []interface{}
}

func NewDict(pairs ...interface{}) (*Dict, error) {
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("number of arguments must be even")
	}

	d := &Dict{
		Keys:   make([]interface{}, len(pairs)/2),
		Values: make([]interface{}, len(pairs)/2),
	}

	for i := 0; i < len(pairs); i += 2 {
		index := i / 2
		d.Keys[index] = pairs[i]
		d.Values[index] = pairs[i+1]
	}

	return d, nil
}

func (d *Dict) Get(key interface{}) (interface{}, error) {
	index := d.findIndex(key)
	if index == -1 {
		return nil, fmt.Errorf("key %v not found", key)
	}
	return d.Values[index], nil
}

func (d *Dict) GetDefault(key, defaultValue interface{}) interface{} {
	index := d.findIndex(key)
	if index == -1 {
		return defaultValue
	}
	return d.Values[index]
}

func (d *Dict) Set(key, value interface{}) *Dict {
	index := d.findIndex(key)
	if index != -1 {
		d.Values[index] = value
	} else {
		d.Keys = append(d.Keys, key)
		d.Values = append(d.Values, value)
	}
	return d
}

func (d *Dict) Delete(key interface{}) error {
	index := d.findIndex(key)
	if index == -1 {
		return fmt.Errorf("key %v not found", key)
	}

	d.Keys = append(d.Keys[:index], d.Keys[index+1:]...)
	d.Values = append(d.Values[:index], d.Values[index+1:]...)
	return nil
}

func (d *Dict) GetKeys() *List {
	keys := make([]interface{}, len(d.Keys))
	copy(keys, d.Keys)
	return &List{Elements: keys}
}

func (d *Dict) GetValues() *List {
	values := make([]interface{}, len(d.Values))
	copy(values, d.Values)
	return &List{Elements: values}
}

func (d *Dict) GetItems() *List {
	items := make([]interface{}, len(d.Keys))
	for i := range d.Keys {
		pair := []interface{}{d.Keys[i], d.Values[i]}
		items[i] = pair
	}
	return &List{Elements: items}
}

func (d *Dict) Len() int {
	return len(d.Keys)
}

func (d *Dict) Clear() *Dict {
	d.Keys = []interface{}{}
	d.Values = []interface{}{}
	return d
}

func (d *Dict) String() string {
	pairs := make([]string, len(d.Keys))
	for i := range d.Keys {
		pairs[i] = fmt.Sprintf("%v: %v", d.Keys[i], d.Values[i])
	}
	return "{" + strings.Join(pairs, ", ") + "}"
}

func (d *Dict) Contains(key interface{}) bool {
	return d.findIndex(key) != -1
}

func (d *Dict) Update(other *Dict) *Dict {
	for i, key := range other.Keys {
		d.Set(key, other.Values[i])
	}
	return d
}

func (d *Dict) findIndex(key interface{}) int {
	for i, k := range d.Keys {
		if reflect.DeepEqual(k, key) {
			return i
		}
	}
	return -1
}
