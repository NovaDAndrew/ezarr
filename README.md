# ezarr

Python-like data structures for Go.

## Installation

```bash
go get github.com/NovaDAndrew/ezarr
```

## Usage

```go
import "github.com/NovaDAndrew/ezarr"
```

## Features

### List

A dynamic array similar to Python's list:

```go
// Create a new list
list := ezarr.New(1, 2, 3, "hello", true)

// Basic operations
list.Append(42)                // Add element to the end
list.Insert(0, "first")        // Insert at specific position
list.Remove(3)                 // Remove element at index
value := list.Pop(1)           // Remove and return element

// Searching and counting
index := list.Index("hello")   // Find index of element
count := list.Count(2)         // Count occurrences

// Sorting and reversing
list.Sort()                    // Sort elements (must be comparable)
list.Reverse()                 // Reverse order

// Slicing and copying
sliced := list.Slice(1, 3)     // Get sublist
copied := list.Copy()          // Create a copy

// Other operations
len := list.Len()              // Get length
list.Clear()                   // Remove all elements
```

### Dict

A key-value store similar to Python's dictionary:

```go
// Create a new dictionary
dict, _ := ezarr.NewDict("name", "Puer", "age", 18)

// Get and set values
name, _ := dict.Get("name")                      // Get value by key
country := dict.GetDefault("country", "Unknown") // Get with default value
dict.Set("city", "New York")                     // Set key-value pair

// Check key existence
exists := dict.Contains("name")          // Check if key exists

// Get collections
keys := dict.GetKeys()                   // Get all keys as List
values := dict.GetValues()               // Get all values as List
items := dict.GetItems()                 // Get all key-value pairs as List

// Modify dictionary
dict.Update(otherDict)                   // Update with another dictionary
dict.Delete("age")                       // Remove key-value pair
value, _ := dict.Pop("name")             // Remove and return value
key, val, _ := dict.PopItem()            // Remove and return last pair

// Create new dictionaries
merged := dict.Merge(otherDict)          // Merge with another dict (non-destructive)
dictFromKeys := ezarr.FromKeys(keys, 0)  // Create dict from keys with default value

// Filter dictionary
filtered := dict.Filter(func(key, value interface{}) bool {
    // Return true to include the key-value pair
    return true
})

// Other operations
len := dict.Len()                        // Get length
dict.Clear()                             // Remove all elements
```

## License

MIT
