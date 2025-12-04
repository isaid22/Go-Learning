package main

import "fmt"

func main() {
	var intArr [5]int32
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3
	intArr[3] = 4
	intArr[4] = 5

	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])
		fmt.Println(&intArr[i])
	}

	var intSlice []int32 = []int32{10, 20, 30}
	fmt.Printf("the length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n After appending, the length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{100, 200}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // specify length and capacity
	fmt.Printf("the length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))

	// map
	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["Adam"] = 25
	myMap["Sarah"] = 90
	fmt.Println(myMap["Adam"])

	var age, ok = myMap["Bob"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Bob not found")
	}

	// Declare a map where keys are strings and values can be any type
	data := make(map[string]interface{})

	// Assign values of different types
	data["name"] = "Alice"
	data["age"] = 30
	data["isStudent"] = true
	data["grades"] = []int{90, 85, 92}

	// Accessing and type asserting values
	fmt.Println("Name:", data["name"].(string))
	fmt.Println("Age:", data["age"].(int))
	fmt.Println("Is Student:", data["isStudent"].(bool))
	fmt.Println("Grades:", data["grades"].([]int))

	// After condition, semicolon marks end of initialization and the beginning of the condition if ok is true, then print Occupation's value.
	if val, ok := data["occupation"].(string); ok {
		fmt.Println("Occupation:", val)
	} else {
		fmt.Println("Occupation not found or not a string.")
	}

}
