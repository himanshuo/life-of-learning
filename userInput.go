package main

import ( 
	"fmt"
	//"strings"
	"strconv"
	"text/scanner"
)


func main(){
	
	//get a single integer
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(a)
	
		
	//get a string (no spaces)
	var b string
	fmt.Scanf("%s",&b)
	fmt.Println(b)
	
	//list of n ints
	var myList []int
	fmt.Print("Enter list: ")
    var input string
    fmt.Scanln(&input)
    fmt.Print(input)
	
	
	var s scanner.Scanner
	s.Init(input)
	tok := s.Scan()
	for tok != scanner.EOF {
		// do something with tok
		tok = s.Scan()
		asInt,_ := strconv.Atoi(tok)
		myList = append(myList, asInt)
	}
	
	fmt.Println(myList)

}
