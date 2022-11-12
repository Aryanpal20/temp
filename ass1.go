// 1. program for hello world
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("helllo world")
// }

// 2. variable declaration with initial value
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var student1 string = "Rohan"
// 	var student2 = "Mohan"
// 	x := 2
// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

// 3. variable declaration without initial values
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var student1 string
// 	var student2 string
// 	var x int
// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

// 4. value assignment after declartion
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var student string
// 	student = "Rohan"
// 	fmt.Println(student)
// }

// 5. this will show declaring variable outside of  a function
// we can use only var keyword outside the function
// package main

// import (
// 	"fmt"
// )

// var x int
// var y int = 2
// var z = 3

// func main() {
// 	x = 1
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// 6. if use this one for declaring variable outside the function it show error
//  package main
// import ("fmt")
// x := 1
// func main(){
// 	fmt.Println(x)
// }
// 7. declare multiple variable with var keyword
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var x, y, z int = 1, 2, 3
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// 8. declare multiple variable with both methods
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b = 1, 2
// 	x, y := 4, 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// 9. variable declartion in a block
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var (
// 		a, b, c = 1, 2, 3
// 	)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// 10. declaring the typed constrants
// package main

// import (
// 	"fmt"
// )

// const a int = 1

// func main() {
// 	fmt.Println(a)
// }

// 11. declaring the untyped constrants
// package main

// import (
// 	"fmt"
// )

// const a = 1

// func main() {
// 	fmt.Println(a)
// }

// 12. multiple constrant declartion
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	const (
// 		A = 1
// 		B = 2
// 		C = 3
// 	)
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(C)
// }

// 13. different type of data types
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int = 2
// 	var b float64 = 3.1
// 	var c string = "hlw"
// 	var d bool = true
// 	fmt.Println("boolean", d)
// 	fmt.Println("integer", a)
// 	fmt.Println("string", c)
// 	fmt.Println("float", b)
// }

// 14. two arrays with defined length
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	var arr2 = [2]string{"hello", "world"}
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// 15. two arrays with inferred length
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = []int{1, 2, 3}
// 	var arr2 = []int{4, 5, 6, 7}
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// 16. access the elememt in the array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [4]string{"rohan", "mohan", "sohan"}
// 	fmt.Println(arr1[2])
// 	fmt.Println(arr1[0])
// }

// 17. chnage element in the array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [4]int{1, 2, 3, 4}
// 	arr1[2] = 2
// 	fmt.Println(arr1)
// }

// 18. array initialize
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [3]int{1}
// 	var arr2 = [6]string{"hlo", "hi"}
// 	var arr3 = [3]bool{}
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// }

// 19. initialize a specific element of an array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [4]int{1: 20, 2: 30}
// 	fmt.Println(arr1)
// 	fmt.Println("")
// 	// initialize a specific element of an array by give them
// 	var arr2 = [4]int{1: 20, 2: 30}
// 	arr2[0] = 10
// 	arr2[3] = 40
// 	fmt.Println(arr2)
// }

// 20. find the length of an array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = []int{1, 2, 3, 4, 5}
// 	var arr2 = []string{"rohan", "mohan", "sohan"}
// 	fmt.Println(len(arr1))
// 	fmt.Println(len(arr2))
// }

// 21. we test the two values to find out which one is greater
// package main

// import (
// 	"fmt"
// )

// var a = 4
// var b = 5

// func main() {
// 	if a > b {
// 		fmt.Println("a is greater than b")
// 	} else {
// 		fmt.Println("b is greater than a")
// 	}
// }

//              // 22. we test the two values to find greater or equal
// package main

// import (
// 	"fmt"
// )

// var a = 6
// var b = 6

// func main() {
// 	if a > b {

// 		fmt.Println("a is greater than b")
// 	} else if b > a {

// 		fmt.Println("b is greater than a")
// 	} else {
// 		fmt.Println("both are equal")
// 	}
// }expression

// 23. we test the three values to find which one is greater
// package main

// import (
// 	"fmt"
// )

// var a = 3
// var b = 4
// var c = 5

// func main() {
// 	if a > b && a > c {
// 		fmt.Println("a is greater than b and c")
// 	} else if b > c && b > a {
// 		fmt.Println("b is greater than a and c")
// 	} else {
// 		fmt.Println("c is greater than a and b")
// 	}
// }

// 24. using nested if
// package main

// import (
// 	"fmt"
// )

// var a = 10

// func main() {
// 	if a >= 30 {
// 		fmt.Println("a is greater than 30")
// 		if a >= 40 {
// 			fmt.Println("a is also greater than 40")
// 		}
// 	} else {
// 		fmt.Println("a is less than 30 and 40")
// 	}
// }

// 25. print 2 table
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("2 * ", i, " = ", 2*i)
// 	}
// }

// 26. this code will skip the value of 5
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i == 5 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// 27. this code will break the value after 5
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i == 5 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

// 28. write a code to take inout from user
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var first string
// 	fmt.Println("enter your first name :- ", first)
// 	fmt.Scanln(&first)

// 	var last string
// 	fmt.Println("enter your last name :- ", last)
// 	fmt.Scanln(&last)
// 	fmt.Println("your full name is :-")
// 	fmt.Println(first + " " + last)
// }

// 29. second largest element in array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = [4]int{1, 5, 3, 4}
// 	var max1 = 0
// 	var max2 = 0

// 	for i := 0; i <= 3; i++ {
// 		if max1 < arr[i] {
// 			max2 = max1
// 			max1 = arr[i]
// 		} else if max2 < arr[i] {
// 			max2 = arr[i]
// 		}
// 	}
// 	fmt.Println(max2)
// }

// 30. largest element in array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = [4]int{1, 5, 3, 4}
// 	var max = arr[0]
// 	for i := 0; i <= 3; i++ {
// 		if max < arr[i] {
// 			max = arr[i]
// 		}
// 	}
// 	fmt.Println((max))
// }

//------------// 31. Go program to find the frequency of an integer array
///////// package main

///////// import (
//////// 	"fmt"
//////// )

/////// func main() {
//////// 	arr := []int{90, 70, 30, 30, 10, 80, 40, 50, 40, 30}
/////// 	freq := make(map[int]int)
////// 	for _, num := range arr {
////// 		freq[num] = freq[num] + 1
/////// 	}
////// 	fmt.Println("Frequency of the Array is : ", freq)
////// }package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// )

// func appendCategory(a []string, b []string) []string {

// 	check := make(map[string]int)
// 	d := append(a, b...)
// 	res := make([]string, 0)
// 	for _, val := range d {
// 		check[val] = 1
// 	}

// 	for letter, _ := range check {
// 		res = append(res, letter)package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// }

// func main() {
// 	a := []string{"x", "y", "z"}
// 	b := []string{"x", "p", "q"}
// 	c := appendCategory(a, b)
// 	fmt.Println(c)

//-----------// }

// package mainpackage main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// 	var d = 0

// 	for i := 0; i <= 5; i++{
// 		if arr[i] == arr1[i]{
// 			d
// 		}
// 	}
// }
//  ---------------------------------------------------------------
//package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )
//package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// }

//  34. count of 100 bt tens
// package main

// import (package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// func main() {
// 	for i := 0; i <= 100; i += 10 {
// 		fmt.Println(i)
// 	}
// }

//  35.  nested loop
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [3]string{"student1", "student2", "student3"}
// 	var arr2 = [3]string{"teacher1", "teacher2", "teacher3"}
// 	for i := 0; i <= len(arr1); i++ {
// 		for j := 0; j <= len(arr2); j++ {

// 		}
// 	}
// 	fmt.Println(arr1, arr2)
// }

//// 36. inner loop executed with each time with outer loop
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [3]string{"student1", "student2", "student3"}
// 	var arr2 = [3]string{"teacher1", "teacher2", "teacher3"}
// 	for i := 0; i < len(arr1); i++ {
// 		for j := 0; j < len(arr2); j++ {
// 			fmt.Println(arr1[i], arr2[j])
// 		}
// 	}
// }

//------- 37. The task is to shift all the zeroes appearing in the array to the end of the array. In Golang
//Input: 1 0 7 0 3

//Output: 1 7 3 0 0

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a = [5]int{1, 0, 7, 0, 3}
// 	var b = 0
// 	var c = len(a)
// 	for i := 0; i < c; i++ {
// 		if a[i] != 0 {
// 			a[b], a[i] = a[i], a[b]
// 			b++
// 		}
// 	}
// 	fmt.Println("shifting the zeroes")
// 	for i := 0; i < c; i++ {
// 		fmt.Printf("%v ", a[i])
// 	}
// }
//----------------------------------------------------------------------------------

// 38.  -----Given an array of n elements, your task is to find out the average of the array.
//---- Approach:

//------------- Accept the size of the array.
// ---------Accept the elements of the array.
// ---------Store the sum of the elements using for loop.
// ----------Calculate Average = (sum/size of array)
// ----------Print the average.\

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 2, 3, 4}
// 	sarr := 4
// 	s := 0
// 	for i := 0; i < sarr; i++ {

// 		s = s + arr[i]
// 	}
// 	avg := (float64(s)) / (float64(sarr))
// 	fmt.Println(s)
// 	fmt.Println(avg)
// }

//  39. Golang program to demonstrate the example of an Array

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = [4]int{1, 2, 3, 4}
// 	fmt.Println("the array is :-")
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Printf("%d ", arr[i])
// 	}
// }

// 40. Golang program to create an array without specifying its size.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println("the arrays element are :-")package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }

// package main

// import (
// 	"fmt"
// )

// func myMessage() {
// 	fmt.Println("this is my first function program")
// }
// func main() {
// 	myMessage()
// 	// here we can call many times
// 	myMessage()
// 	myMessage()
// 	myMessage()
// }

// 42. function with parameter

// package main
// import (
// 	"fmt"
// )

// func information(name string, mark int, subject string) {
// 	fmt.Println("my name is ", name)
// 	fmt.Println("my mark is ", mark)
// 	fmt.Println("the subject is ", subject)
// }
// func main() {
// 	information("rohan", 78, "maths")
// 	information("mohan", 80, "maths")
// }

// 43. myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int)

// package main

// import (
// 	"fmt"
// )

// func integer(x int, y int) (result int) {
// 	result = x + y
// 	return result
// }
// func main() {
// 	fmt.Println(integer(2, 4))
// 	fmt.Println(integer(10, 20))
// }

// 44. we store the return value in a variable

// package main

// import (
// 	"fmt"
// )

// func valueinvariable(x int, y int) (result int) {
// 	result = x + y
// 	return
// }
// func main() {
// 	var total int
// 	total = valueinvariable(2, 6)
// 	fmt.Println(total)
// }

//  45. to create a slice

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice := []string{"apple", "mango"}
// 	fmt.Println(myslice)
// 	fmt.Println(len(myslice))
// }

//   46. create a slice by slicing an array:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = [3]int{3, 4, 5}
// 	myslice := arr[1:3]

// 	fmt.Println(myslice)
// 	fmt.Println(len(myslice))
// 	fmt.Println(cap(myslice))
// }

// 47. The make() function can also be used to create a slice.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice := make([]int, 4)
// 	myslice = append(myslice, 1, 2, 3, 4)
// 	myslice = append(myslice[4:])// we can use this code for remove first 4 zero
// 	fmt.Println(myslice)
// }

// 48.  to access the first and third elements in the slice:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(myslice[0])
// 	fmt.Println(myslice[2])
// }

// 49. change a specific slice element by referring to the index number.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice := []int{1, 2, 3, 4, 5}
// 	myslice[3] = 5
// 	fmt.Println(myslice)
// }

// 50. how to append one slice to another slice:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)
// 	myslice4 := []int{7, 8, 9}
// 	myslice5 := append(myslice3, myslice4...)
// 	fmt.Println(myslice5)
// }
// 51.  how to calculate the sum of all array elements

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [5]int{2, 3, 4, 5, 6}
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += arr[i]
// 	}
// 	fmt.Println("the sum of array is :- ", sum)
// }

//  52. how to compare two arrays using equal to (==) operator

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr1 := [3]int{1, 2, 3}
// 	arr2 := [3]int{4, 2, 3}
// 	arr3 := [3]int{1, 2, 3}

// 	if arr1 == arr2 {
// 		fmt.Println("arr1 and arr2 are similar")
// 	} else {
// 		fmt.Println("arr1 and arr2 are not similar")
// 	}
// 	if arr1 == arr3 {
// 		fmt.Println("arr1 and arr3 are similar")
// 	} else {
// 		fmt.Println("arr1 and arr3 are not similar")
// 	}
// 	if arr2 == arr3 {
// 		fmt.Println("arr2 and arr3 are similar")
// 	} else {
// 		fmt.Println("arr2 and arr3 are not similar")
// 	}

// }

// 53. how to print even numbers of the array

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	for i := 0; i < 5; i++ {
// 		if arr[i]%2 == 0 {
// 			fmt.Println("the even number of an array", arr[i])
// 		}
// 	}
// }

//  54. how to merge two arrays into third array

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr [10]int
// 	arr1 := []int{0, 1, 2, 3, 4, 5}
// 	arr2 := []int{6, 7, 8, 9}
// 	for i := 0; i <= 5; i++ {
// 		arr[i] = arr1[i]
// 	}
// 	var j int = 6
// 	for i := 0; i <= 3; i++ {
// 		arr[j] = arr2[i]
// 		j = j + 1
// 	}
// 	for i := 0; i <= 9; i++ {
// 		fmt.Println(arr[i])
// 	}

// }

//  55. how to reverse an integer array

// package main

// import (
// 	"fmt"
// )
// func main() {
// 	var arr1 [4]int
// 	var j int = 3
// 	arr := [4]int{1, 2, 3, 4}
// 	for i := 0; i < 4; i++ {
// 		arr1[j] = arr[i]
// 		j = j - 1
// 	} // 	for i := 0; i < 4; i++ {
// 		fmt.Printf("%d ", arr1[i])
// 	}
// }

// 56. print the prime numbers from an integer array

package main

func main() {

}
