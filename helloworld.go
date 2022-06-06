package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// Single Line Comment
	/*
	   Multiline Line Comment
	*/

	// Data Types in Go
	/*
		int
		float32
		string
		bool
	*/

	// Declaring a variable
	var something int = 32
	// Variable type is inferred when using :=
	somethingElse := 64

	// Decalaring without initial value
	var aString string
	var aInt int
	var aBool bool

	// := can only be used inside functions

	// Multiple Variables declaration
	var a, b, c, d = 1, 2, 3, 4
	// Declaring Varibales in a Block
	var (
		x string = "x"
		y int    = 1
		z bool   = true
	)

	// const keyword
	const myConst string = "constant"
	const untypedConst = "untyped constant"

	fmt.Println(something, somethingElse)
	fmt.Println(aString, aInt, aBool)
	fmt.Println(a, b, c, d)
	fmt.Println(x, y, z)
	fmt.Println(myConst, untypedConst)

	fmt.Print(myConst, "\n")
	fmt.Println(myConst)
	fmt.Printf("my constant %s\n", untypedConst)

	// Arrays
	// Using var
	// var array_name = [Length]datatype{values} - defined length
	// var array_name = [...]datatype{values} - infered length
	var myArr = [3]int{10, 20, 30}
	var mySecondArray = [2]string{"hello", "world"}

	// Using :=
	// array_name = [Length]datatype{values} - defined length
	// array_name = [...]datatype{values} - infered length
	var myThirdArray = [2]string{"hello", "world"}

	fmt.Println(myArr)
	fmt.Println(mySecondArray)
	fmt.Println(myThirdArray)
	fmt.Println(myArr[0])
	fmt.Println(len(myArr))

	// Slices - similiar to Array but more powerful
	var mySlice = []int{10, 20, 30}
	var mySecondSlice = []string{"hello", "world"}

	fmt.Print(mySlice, "\n")
	fmt.Print(mySecondSlice, "\n")

	// len - length of slice
	// cap - capacity of slice
	fmt.Println(len(mySlice), cap(mySlice))

	// apending elements to a slice
	mySlice = append(mySlice, 7, 8, 9)
	fmt.Println(mySlice)

	// appending one slice to another slice
	mySlice = append(mySlice, mySlice...)
	fmt.Println(mySlice)

	// Operators and Conditionals

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for idx, value := range mySlice {
		fmt.Printf("%v\t%v\n", idx, value)
	}

	myFunc()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t%v\n", i, factorial(i))
	}

	// structs
	type Person struct {
		name string
		age  int
	}

	var person Person
	person.name = "NAME"
	person.age = 45
	fmt.Println(person)

	// Maps
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	

}

// functions
func myFunc() {
	fmt.Println("My Function")
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
