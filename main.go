package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Print the address of the array
	fmt.Printf("Address of the array: %p\n", &numbers)

	// Slices
	slice := numbers[1:3]
	fmt.Println("Slice:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	slice[0] = 10
	slice = append(slice, 20)
	slice = append(slice, 20)
	slice = append(slice, 20)
	slice = append(slice, 20)
	fmt.Println("Modified Slice:", slice)
	fmt.Println("Modified Array:", numbers)
	fmt.Println("Slice capacity:", cap(slice))
}

func init() {
	fmt.Println("hello world")
}

/*
	1. slice from an existing array
	2. slice from a slice
	3. slice literal
	4. make function with len
	5. make function with len and capacity
	6. empty slice or nil slice
	7. slice underlying array rule => 1024 -> 100% increase, 1024 < 25% increase
*/
