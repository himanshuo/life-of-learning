//this file is going to help me understand arrays and slices in Go
package main

import "fmt"

//ARRAYS
// the type of an array is [n]T. This is the actual type! NOT just a way to represent the type. The actual type.
//[n]T  is an array of type T with n capacity.


//Arrays are a building block for slices.
//key things about arrays:
//1) Arrays are values. Assigning one array to another copies !!!!!!ALL!!!!! the elements.
//2) if you pass an array to a function, it will receive a copy of the array, not a pointer to it!!!!!!!
//3) The size of an array is part of its type. The types [10]int and [20]int are distinct.

//notably, the first of the three things mentioned is EXPENSIVE. 
//Since arrays are values, you end up having to copy EVERYTHING (2nd property) when passing to a function. 
//if you want C-like behavior and efficiency, you can pass a pointer to the array. And that is slices.

//input param a is a pointer to an array with capacity 3 and type float64. 
//NOTE: the (sum float64)
func Sum(a *[3]float64) (sum float64) {
    /*
    	determines sum of values inside array.
    */

    for _, v := range *a {
        sum += v
    }
    return sum
}

func arrayCode() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) 
	fmt.Println(x)

	var a[2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	//can print array very easily. 
	fmt.Println(array)
}




//SLICES

//A slice points to an array of values and also includes a length.
//[]T is a slice with elements of type T.  [1]T is a ARRAY of type T with capacity 1.
//NOTE: to create an array, you specify the capacity. For slice, you don't.



func sliceBasics(){
	p := []int{2, 3, 5, 7, 11, 13}//slice with initializer values
	fmt.Println("p ==", p)//print p == [2 3 5 7 11 13]
	
	for i := 0; i < len(p); i++ { 
		fmt.Printf("p[%d] == %d\n", i, p[i]) //print each element of p. p[0] == 2 ...
	}
	
	
	var myarr [9]int // array of length 9
	myarr[3] = 3 
	// turn array into slice by using its & syntax.
	//this is more than just taking the reference. A slice also holds the length of the slice.
	var slice = &myarr 

	//dereference the slice and print the underlying array 
	fmt.Println(*slice)
}


//Slices can be re-sliced. This creates a new slice value that points to the same array.
//myslice[lo:hi]  evaluates to a slice of the elements from lo through hi-1, inclusive. 
//myslice[lo:lo] is empty
//myslice[lo:lo+1] has 1 element

func slicingSlices(){

	p := []int{2, 3, 5, 7, 11, 13}//slice.
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4]) //slice of a slice. the upper slice has only elements 1,2, 3. 

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])  //slice of slice has elements 0, 1, 2

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])//slice of slice has elements 4,5
}

//Making slices
//Slices are created with the make function. 
//make() allocates a zeroed array and then returns a slice that refers to that array.
//make takes optional params of length and capacity. 


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func makeSlice(){
	a := make([]int, 6) // make slice of int array. length of slice is 6. since no capacity specified, length=capacity
	printSlice("a", a) //len=capacity=6
	b := make([]int, 0, 5) // make slice of int array. length is 0. capacity is 5. I think this makes the innerarray have capacity of 5.
	printSlice("b", b)//len=0, cap=5
	
	
	c := b[:2] //make a new slice using b. length is 2. capacity is still 5 since that was the capacity of b.
	printSlice("c", c)//len=2, cap=5
	
	
	//can do this because capacity is 5 for c. c however only has length 2 so this is going past the slices length.
	//in this case, the new slice is just takes entirely new values. 
	//len=cap=3  like making a new slice with only the length specified.
	d := c[2:5] 
	printSlice("d", d)//len=3, cap=3
}

//Nil Slices
//zero value of a slice is length=capacity=0. empty array.
//can use a special nil slice

func nilSlice(){
	//z is a slice. since did not use make(), the zero value of the slice is used.
	//the zero value has len=cap=0. the array that it points to is empty.
	var z []int
	fmt.Println(z, len(z), cap(z))// [] 0 0 
	if z == nil {
		fmt.Println("nil!")
	}
}


//Append
//built-in function allows you to append to a slice. function will grow the slice if need be.
func appendSlice(){
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}








func sliceCode(){
	fmt.Println("###slice basics###")
	sliceBasics()
	fmt.Println("###slicing slices###")
	slicingSlices()
	fmt.Println("###making slices###")
	makeSlice()
	fmt.Println("###nil slice###")
	nilSlice()
	fmt.Println("###append slice###")
	appendSlice()
}

func main() {
fmt.Println("----------------array code--------------------------")
arrayCode()
fmt.Println("----------------slide code--------------------------")
sliceCode()

	
}
