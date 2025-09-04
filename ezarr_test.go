package ezarr

import (
	"reflect"
	"testing"
)

// Test | New verifies the creation of a new list with initial elements
func TestNew(t *testing.T) {
	list := New(1, 2, 3)
	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	expected := []interface{}{1, 2, 3}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}
}

// Test | Append verifies adding elements to the end of a list
func TestAppend(t *testing.T) {
	list := New(1, 2)
	list.Append(3)

	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	if list.Elements[2] != 3 {
		t.Errorf("Expected 3 at index 2, got %v", list.Elements[2])
	}
}

// Test | Extend verifies adding all elements from one list to another
func TestExtend(t *testing.T) {
	list1 := New(1, 2)
	list2 := New(3, 4)

	list1.Extend(list2)

	if list1.Len() != 4 {
		t.Errorf("Expected length 4, got %d", list1.Len())
	}

	expected := []interface{}{1, 2, 3, 4}
	for i, v := range expected {
		if list1.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list1.Elements[i])
		}
	}
}

// Test | Insert verifies inserting elements at specific positions, including with negative indices
func TestInsert(t *testing.T) {
	list := New(1, 3)

	// Insert in the middle
	list.Insert(1, 2)
	expected := []interface{}{1, 2, 3}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}

	// Insert at the beginning
	list.Insert(0, 0)
	expected = []interface{}{0, 1, 2, 3}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}

	// Insert at the end
	list.Insert(4, 4)
	expected = []interface{}{0, 1, 2, 3, 4}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}

	// Insert with negative index
	list.Insert(-1, 3.5)
	expected = []interface{}{0, 1, 2, 3, 3.5, 4}
	for i, v := range expected {
		if !reflect.DeepEqual(list.Elements[i], v) {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}
}

// Test | Remove verifies removing elements from a list
func TestRemove(t *testing.T) {
	list := New(1, 2, 3, 2)

	// Remove existing element
	err := list.Remove(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []interface{}{1, 3, 2}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}

	// Remove non-existent element
	err = list.Remove(4)
	if err == nil {
		t.Error("Expected error when removing non-existent element, got nil")
	}
}

// Test | Pop verifies removing and returning elements at specific indices
func TestPop(t *testing.T) {
	list := New(1, 2, 3)

	// Pop with positive index
	val, err := list.Pop(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected popped value 2, got %v", val)
	}
	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	// Pop with negative index
	val, err = list.Pop(-1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected popped value 3, got %v", val)
	}
	if list.Len() != 1 {
		t.Errorf("Expected length 1, got %d", list.Len())
	}

	// Pop with invalid index
	_, err = list.Pop(5)
	if err == nil {
		t.Error("Expected error when popping with invalid index, got nil")
	}

	// Pop from empty list
	list.Clear()
	_, err = list.Pop(0)
	if err == nil {
		t.Error("Expected error when popping from empty list, got nil")
	}
}

// Test | Index verifies finding the position of elements in a list
func TestIndex(t *testing.T) {
	list := New(1, 2, 3, 2)

	// Find existing element
	index := list.Index(2)
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}

	// Find non-existent element
	index = list.Index(4)
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

// Test | Count verifies counting occurrences of elements in a list
func TestCount(t *testing.T) {
	list := New(1, 2, 3, 2, 2)

	// Count existing element
	count := list.Count(2)
	if count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}

	// Count non-existent element
	count = list.Count(4)
	if count != 0 {
		t.Errorf("Expected count 0, got %d", count)
	}
}

// Test | Sort verifies sorting lists of different types
func TestSort(t *testing.T) {
	// Sort integers
	list := New(3, 1, 4, 2)
	err := list.Sort()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []interface{}{1, 2, 3, 4}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}

	// Sort strings
	strList := New("banana", "apple", "cherry")
	err = strList.Sort()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedStr := []interface{}{"apple", "banana", "cherry"}
	for i, v := range expectedStr {
		if strList.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, strList.Elements[i])
		}
	}

	// Sort mixed types
	mixedList := New(1, "apple")
	err = mixedList.Sort()
	if err == nil {
		t.Error("Expected error when sorting mixed types, got nil")
	}
}

// Test | Reverse verifies reversing the order of elements in a list
func TestReverse(t *testing.T) {
	list := New(1, 2, 3, 4)
	list.Reverse()

	expected := []interface{}{4, 3, 2, 1}
	for i, v := range expected {
		if list.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, list.Elements[i])
		}
	}
}

// Test | Slice verifies creating new lists from slices of existing lists
func TestSlice(t *testing.T) {
	list := New(0, 1, 2, 3, 4)

	// Normal slice
	slice := list.Slice(1, 4)
	expected := []interface{}{1, 2, 3}
	for i, v := range expected {
		if slice.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, slice.Elements[i])
		}
	}

	// Slice with negative indices
	slice = list.Slice(-3, -1)
	expected = []interface{}{2, 3}
	for i, v := range expected {
		if slice.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, slice.Elements[i])
		}
	}

	// Slice with out of bounds indices
	slice = list.Slice(-10, 10)
	expected = []interface{}{0, 1, 2, 3, 4}
	for i, v := range expected {
		if slice.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, slice.Elements[i])
		}
	}
}

// Test | Copy verifies creating independent copies of lists
func TestCopy(t *testing.T) {
	list := New(1, 2, 3)
	copy := list.Copy()

	// Check that copy contains the same elements
	for i, v := range list.Elements {
		if copy.Elements[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, copy.Elements[i])
		}
	}

	// Check that modifying original doesn't affect copy
	list.Append(4)
	if copy.Len() != 3 {
		t.Errorf("Expected copy length 3, got %d", copy.Len())
	}

	// Check that modifying copy doesn't affect original
	copy.Append(5)
	if list.Len() != 4 {
		t.Errorf("Expected original length 4, got %d", list.Len())
	}
}

// Test | Clear verifies emptying a list
func TestClear(t *testing.T) {
	list := New(1, 2, 3)
	list.Clear()

	if list.Len() != 0 {
		t.Errorf("Expected length 0, got %d", list.Len())
	}
}

// Test | String verifies string representation of a list
func TestString(t *testing.T) {
	list := New(1, "two", 3.0)
	str := list.String()
	expected := "[1, two, 3]"

	if str != expected {
		t.Errorf("Expected string representation '%s', got '%s'", expected, str)
	}
}
