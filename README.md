# HashMap 
 A simple implmentation of a <b>HashMap</b> that allows dynamic types for keys and values in Golang.

# Download and install
```
$ go get github.com/kjunmin/HashMap
```

# Usage
```
package main

import (
    "fmt"
    "github.com/kjunmin/HashMap"
)

func main() {
	hashMap := Init(64, HashFunc)

	hashMap.Insert(1, "Test")
	hashMap.Insert(2, 777)
	hashMap.Insert("Hello", "World")
	hashMap.Insert("Slices", []int{5, 6, 7, 8})
	fmt.Println(hashMap.Get(1))
	// output: Test
	fmt.Println(hashMap.Get("Hello"))
	// output: World
	fmt.Println(hashMap.Get("Slices"))
	//output: [5, 6, 7, 8]
	fmt.Println(hashMap.Get(3))
	// output: Key not found
	fmt.Println(hashMap.Count())
	// output: 4
	hashMap.Erase("Slices")
	fmt.Println(hashMap.Get("Slices"))
	// output: Key not found
	fmt.Println(hashMap.Count())
	// output: 3
}
```