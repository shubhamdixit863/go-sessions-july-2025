package main

import (
	"fmt"
	_ "fmt"
)

var (
	i = 9
	g = "string"
	h bool
)

//c := 90  // should always be inside the function

const Pi = 3.1415
const RATE_LIMIT = 5

func hi() (string, int) {
	return "hello", 2
}
func main() {

	// How to create variables in go
	//var c int // 8 bits ,16 bits 32 bits 64 bits
	//
	//c = 90
	//fmt.Println(c)
	//
	//var d string = "jsjsjs"
	//
	//fmt.Println("String data", d)
	//
	//var b bool
	//fmt.Println("bool data", b)

	// short hand variable declaration
	//
	//c := 90
	//fmt.Println(c)
	//
	//var x int      // declaration only, zero value
	//var y int = 42 // explicit type + initializer
	//var z = 100       // type inferred as int

	// Multiple declarations

	//var a ,b int
	//var c, d = 1, 2 // c  infereed as int and d as string
	//fmt.Println(c)
	//fmt.Println(d)
	//g = "another string"
	//fmt.Println(g)
	//fmt.Println(Pi)

	// blank identifier
	_, g := hi()

	fmt.Println(g)

	// data types in go

	var flag bool = true
	fmt.Println(flag)

	// Numeric types
	// signed ---> int ,int8 ,int16
	// unsigned
	var k int8 // -128 to 127     2^7 to 2^7 -1
	k = 2      //
	fmt.Println(k)
	var j int16 // -32768 to 32767
	//j = 34789
	fmt.Println(k, j)

	// floating point numbers
	// float32 and float64
	//var f float64 = 3.90900

	// complex  i (-1) 1+2i
	c := complex(1, 2)
	fmt.Println(c)

	// String  // immuatble in go

	//	var f string = "string \n"
	//
	//	var t = `
	//jddjjdjd
	//jddjdjdjdj
	//dkdkkdkd
	//
	//`

	l, m, n := Arithmetic(2, 3) // Arithmetic(2,3)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	AlSum(9, 9, 7, 6, 7, 1, 2, 3, 0, 1781, 0, 90)
}

// polymorphism ,abstraction .inheritance  ,//encapsulation  -->oops princple
// function programming  (go does composition)

func Arithmetic(a, b int) (int, int, int) {
	sum := a + b
	sub := a - b
	mul := a * b
	return sum, sub, mul
}

func AlSum(b ...int) int { // variadic functions
	fmt.Println(b)
	return 9
}
