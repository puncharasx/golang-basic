package main

import (
	"fmt"
	"gobasic/customer"
	"unicode/utf8"
)

type Person struct {
	Name    string
	Surname string
	Age     int

	// Can't decalation in struct
	// but you can create method
	// see below
}

// receiver
func (p Person) Hello() string {
    return "Hello" + p.name
}

func main() {
	// Struct instance
	s := Person{
		Name:    "Puncharas",
		Surname: "Phosap",
		Age:     12,
	}

	println(s.Hello())

	o := customer.Customer{}
	o.Setcustomer("Kuy")
	println("Recive = ", o.Getcustomer())

	// " Default value = Zero value "
	// Declaration
	var x int = 10
	var y = 10
	z := 10 // Short Declaration
	_ = x
	_ = y
	_ = z

	point := 50
	// if else
	if point >= 50 {

	} else {

	}

	// array - slide
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = append(arr, 4)
	arrLength := len(arr)

	arrUp := arr[1:]
	arrSide := arr[1:5] // Index + 1

	fmt.Printf("%v \n -- %v", arrUp, arrSide) // Array 2 upside

	fmt.Printf("%#v\n %v", arr, arrLength) // %#v tell length array

	// count charactor
	name := "Hello"
	utf8.RuneCountInString(name)

	countries := map[string]string{}
	countries["th"] = "Thailand"

	// b, o := 10,20     Tuple Decalartion

	// Map return 2 value
	// Value 1 = Value to Varible country , Value 2 = Bool to ok
	counrty, ok := countries["th"]
	if ok == true {
		println(counrty, ok)
	} else {
		println("No value")
	}

	// For loop
	values := []int{10, 20, 30, 40, 50, 60}
	for i := 0; i < len(values); i++ {
		println(values[i])
	}

	for _, v := range values {
		println(v)
	}

	c, _, _ := sum(10, 20)
	println(c)

	// Anonymous Function
	r := func(a, b int) int {
		return a * b
	}
	r(2, 3)

	// Pointer

	var p int
	p = 10

	var t *int
	t = &p
	*t = 40

	println(&p)
	println(*t)

	var point1 int
	println(point1)
	pointer(&point1)
	println("Point1 :", point1)

}

// function type => (int,int) int
func sum(z, x int) (int, string, bool) { // tuple (int, string)
	return z + x, "Hello", true
}

func sub(a, b int) int {
	return a - b
}

func hello(name string) string {
	return "hello " + name
}

func pointer(result *int) {
	a := 10
	b := 20
	*result = a * b // * For add value
}

func omaiwa(a *int) {
    

    &a

}
