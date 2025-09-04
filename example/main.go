package main

import (
	"fmt"

	"github.com/NovaDAndrew/ezarr"
)

func main() {
	fmt.Println("=== List Example ===")

	list := ezarr.New(1, 2, 3, "hello", true)
	fmt.Println("List:", list)

	list.Append(42)
	fmt.Println("After append(42):", list)

	list.Insert(0, "first")
	fmt.Println("After insert(0, \"first\"):", list)

	list.Remove(3)
	fmt.Println("After remove(3):", list)

	numList := ezarr.New(5, 3, 1, 4, 2)
	fmt.Println("Number list:", numList)
	numList.Sort()
	fmt.Println("After sort:", numList)
	numList.Reverse()
	fmt.Println("After reverse:", numList)

	fmt.Println("\n=== Dict Example ===")

	dict, _ := ezarr.NewDict("name", "Puer", "age", 18, "city", "Murom", "country", "Russia")
	fmt.Println("Dict:", dict)

	name, _ := dict.Get("name")
	fmt.Println("Name:", name)

	country := dict.GetDefault("country", "Unknown")
	fmt.Println("Country (with default):", country)

	dict.Set("country", "USA")
	fmt.Println("After setting country:", dict)

	fmt.Println("Contains 'name':", dict.Contains("name"))
	fmt.Println("Contains 'email':", dict.Contains("email"))

	fmt.Println("Keys:", dict.GetKeys())
	fmt.Println("Values:", dict.GetValues())
	fmt.Println("Items:", dict.GetItems())

	otherDict, _ := ezarr.NewDict("email", "puer@example.com", "age", 18)
	dict.Update(otherDict)
	fmt.Println("After update from other dict:", dict)
}
