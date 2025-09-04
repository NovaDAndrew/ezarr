package ezarr

import (
	"strings"
	"testing"
)

func TestNewDict(t *testing.T) {
	dict, err := NewDict("name", "Puer", "age", 18)
	if err != nil {
		t.Errorf("NewDict returned error: %v", err)
	}

	if dict.Len() != 2 {
		t.Errorf("Expected length 2, got %d", dict.Len())
	}

	name, err := dict.Get("name")
	if err != nil || name != "Puer" {
		t.Errorf("Expected name to be 'Puer', got %v", name)
	}

	age, err := dict.Get("age")
	if err != nil || age != 18 {
		t.Errorf("Expected age to be 18, got %v", age)
	}

	_, err = NewDict("name", "Puer", "age")
	if err == nil {
		t.Error("Expected error for odd number of arguments, got nil")
	}
}

func TestFromKeys(t *testing.T) {
	keys := []interface{}{"a", "b", "c"}
	dict := FromKeys(keys, 0)

	if dict.Len() != 3 {
		t.Errorf("Expected length 3, got %d", dict.Len())
	}

	for _, key := range keys {
		val, err := dict.Get(key)
		if err != nil || val != 0 {
			t.Errorf("Expected value 0 for key %v, got %v", key, val)
		}
	}
}

func TestGet(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	name, err := dict.Get("name")
	if err != nil || name != "Puer" {
		t.Errorf("Expected name to be 'Puer', got %v, error: %v", name, err)
	}

	_, err = dict.Get("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent key, got nil")
	}
}

func TestGetDefault(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	name := dict.GetDefault("name", "Unknown")
	if name != "Puer" {
		t.Errorf("Expected name to be 'Puer', got %v", name)
	}

	country := dict.GetDefault("country", "Unknown")
	if country != "Unknown" {
		t.Errorf("Expected default value 'Unknown', got %v", country)
	}
}

func TestSet(t *testing.T) {
	dict, _ := NewDict()

	dict.Set("name", "Puer")
	name, err := dict.Get("name")
	if err != nil || name != "Puer" {
		t.Errorf("Expected name to be 'Puer', got %v", name)
	}

	dict.Set("name", "Jane")
	name, _ = dict.Get("name")
	if name != "Jane" {
		t.Errorf("Expected name to be 'Jane', got %v", name)
	}
}

func TestDelete(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	err := dict.Delete("name")
	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}

	if dict.Len() != 1 {
		t.Errorf("Expected length 1, got %d", dict.Len())
	}

	_, err = dict.Get("name")
	if err == nil {
		t.Error("Expected error for deleted key, got nil")
	}

	err = dict.Delete("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent key, got nil")
	}
}

func TestGetKeys(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	keys := dict.GetKeys()
	if keys.Len() != 2 {
		t.Errorf("Expected 2 keys, got %d", keys.Len())
	}

	hasName := false
	hasAge := false
	for _, k := range keys.Elements {
		if k == "name" {
			hasName = true
		}
		if k == "age" {
			hasAge = true
		}
	}

	if !hasName || !hasAge {
		t.Errorf("Keys should contain 'name' and 'age', got %v", keys)
	}
}

func TestGetValues(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	values := dict.GetValues()
	if values.Len() != 2 {
		t.Errorf("Expected 2 values, got %d", values.Len())
	}

	hasPuer := false
	has18 := false
	for _, v := range values.Elements {
		if v == "Puer" {
			hasPuer = true
		}
		if v == 18 {
			has18 = true
		}
	}

	if !hasPuer || !has18 {
		t.Errorf("Values should contain 'Puer' and 18, got %v", values)
	}
}

func TestGetItems(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	items := dict.GetItems()
	if items.Len() != 2 {
		t.Errorf("Expected 2 items, got %d", items.Len())
	}

	hasNamePuer := false
	hasAge18 := false
	for _, item := range items.Elements {
		pair, ok := item.([]interface{})
		if !ok || len(pair) != 2 {
			t.Errorf("Expected item to be a pair, got %v", item)
			continue
		}

		if pair[0] == "name" && pair[1] == "Puer" {
			hasNamePuer = true
		}
		if pair[0] == "age" && pair[1] == 18 {
			hasAge18 = true
		}
	}

	if !hasNamePuer || !hasAge18 {
		t.Errorf("Items should contain ['name', 'Puer'] and ['age', 18], got %v", items)
	}
}

func TestDictLen(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	if dict.Len() != 2 {
		t.Errorf("Expected length 2, got %d", dict.Len())
	}

	dict.Set("country", "USA")
	if dict.Len() != 3 {
		t.Errorf("Expected length 3, got %d", dict.Len())
	}

	dict.Delete("name")
	if dict.Len() != 2 {
		t.Errorf("Expected length 2, got %d", dict.Len())
	}
}

func TestDictClear(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	dict.Clear()
	if dict.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", dict.Len())
	}
}

func TestDictString(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	str := dict.String()
	if !strings.Contains(str, "name: Puer") || !strings.Contains(str, "age: 18") {
		t.Errorf("String representation should contain 'name: Puer' and 'age: 18', got %s", str)
	}
}

func TestContains(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	if !dict.Contains("name") {
		t.Error("Expected Contains('name') to be true")
	}

	if dict.Contains("country") {
		t.Error("Expected Contains('country') to be false")
	}
}

func TestUpdate(t *testing.T) {
	dict1, _ := NewDict("name", "Puer", "age", 18)
	dict2, _ := NewDict("country", "USA", "age", 31)

	dict1.Update(dict2)

	if dict1.Len() != 3 {
		t.Errorf("Expected length 3, got %d", dict1.Len())
	}

	age, _ := dict1.Get("age")
	if age != 31 {
		t.Errorf("Expected age to be updated to 31, got %v", age)
	}

	country, err := dict1.Get("country")
	if err != nil || country != "USA" {
		t.Errorf("Expected country to be 'USA', got %v", country)
	}
}

func TestMerge(t *testing.T) {
	dict1, _ := NewDict("name", "Puer", "age", 18)
	dict2, _ := NewDict("country", "USA", "age", 31)

	merged := dict1.Merge(dict2)

	age1, _ := dict1.Get("age")
	if age1 != 18 {
		t.Errorf("Original dict1 should be unchanged, age should be 18, got %v", age1)
	}

	if merged.Len() != 3 {
		t.Errorf("Expected merged length 3, got %d", merged.Len())
	}

	age, _ := merged.Get("age")
	if age != 31 {
		t.Errorf("Expected merged age to be 31, got %v", age)
	}

	country, err := merged.Get("country")
	if err != nil || country != "USA" {
		t.Errorf("Expected merged country to be 'USA', got %v", country)
	}
}

func TestDictPop(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	value, err := dict.Pop("name")
	if err != nil || value != "Puer" {
		t.Errorf("Expected Pop('name') to return 'Puer', got %v, error: %v", value, err)
	}

	if dict.Len() != 1 {
		t.Errorf("Expected length 1 after Pop, got %d", dict.Len())
	}

	if dict.Contains("name") {
		t.Error("Expected 'name' to be removed after Pop")
	}

	_, err = dict.Pop("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent key, got nil")
	}
}

func TestPopItem(t *testing.T) {
	dict, _ := NewDict("name", "Puer", "age", 18)

	key, value, err := dict.PopItem()
	if err != nil {
		t.Errorf("PopItem returned error: %v", err)
	}

	if key != "name" && key != "age" {
		t.Errorf("Expected key to be 'name' or 'age', got %v", key)
	}

	if (key == "name" && value != "Puer") || (key == "age" && value != 18) {
		t.Errorf("Value %v doesn't match key %v", value, key)
	}

	if dict.Len() != 1 {
		t.Errorf("Expected length 1 after PopItem, got %d", dict.Len())
	}

	if dict.Contains(key) {
		t.Errorf("Expected key %v to be removed after PopItem", key)
	}

	_, _, err = dict.PopItem()
	if err != nil {
		t.Errorf("Second PopItem returned error: %v", err)
	}

	_, _, err = dict.PopItem()
	if err == nil {
		t.Error("Expected error for empty dictionary, got nil")
	}
}

func TestFilter(t *testing.T) {
	dict, _ := NewDict("a", 1, "b", 2, "c", 3, "d", 4)

	filtered := dict.Filter(func(key, value interface{}) bool {
		val, ok := value.(int)
		return ok && val%2 == 0
	})

	if filtered.Len() != 2 {
		t.Errorf("Expected filtered length 2, got %d", filtered.Len())
	}

	hasB2 := false
	hasD4 := false

	items := filtered.GetItems()
	for _, item := range items.Elements {
		pair := item.([]interface{})
		if pair[0] == "b" && pair[1] == 2 {
			hasB2 = true
		}
		if pair[0] == "d" && pair[1] == 4 {
			hasD4 = true
		}
	}

	if !hasB2 || !hasD4 {
		t.Errorf("Filtered dict should contain b:2 and d:4, got %v", filtered)
	}

	if dict.Len() != 4 {
		t.Errorf("Original dict should be unchanged, length should be 4, got %d", dict.Len())
	}
}
