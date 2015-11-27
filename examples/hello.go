package main

import (
		"fmt"

		//"github.com/himanshuo/stringutil"
		"github.com/himanshuo/evaler"
)

type himanshu struct{
	rock string
	cool int64
}

func main() {
	//fmt.Printf(stringutil.Reverse("\ndlroW ,olleH1"))
	//var test_expression = "1d+1m*now*mon+3day*(2/7y-3h)"
	var test_expression = "1d+1m*now*mon+3day*((2/4)y-3h)"
	var result, err = evaler.Eval(test_expression)
	if err!=nil{
		fmt.Println("error. input string is:" + test_expression)
		return
	}
	var intResult int64
	intResult, err = evaler.BigratToInt(result)
	if err!=nil{
		fmt.Println("unable to conver to int")
		return
	}
	 fmt.Print("resulting value is:")
	 fmt.Println(intResult)
	 fmt.Println(result)


	 fmt.Println("-----------------------------------")
	var rock []himanshu
	rock = append(rock, himanshu{"rock", 4})
	rock = append(rock, himanshu{"hahaha", 6})
	fmt.Println(rock)
	for i, curhimanshu := range rock{
		curhimanshu.rock = "i wanna touch you"
		rock[i]=curhimanshu
		fmt.Println(i)


	}
	fmt.Println(rock)

}


