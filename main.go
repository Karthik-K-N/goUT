package main

import (
	"fmt"
	"goUT/client"
)

type myStruct struct {
}

func main() {
	myVariable := myStruct{}
	num1 := 1
	num2 := 2
	res, err := processInput(myVariable, num1,num2)
	if err != nil{
		panic(err)
	}else {
		fmt.Printf("Sum of %d and %d is  %d\n", num1, num2, res)
	}
}

func processInput(input client.Client, num1, num2 int) (int, error){
	res, err := input.GetSum(num1, num2)
	if err != nil{
		return 0, err
	}
	return res, nil
}

func (input myStruct) GetSum( num1, num2 int) (int, error) {
	return num1+num2, nil
}

//func main() {
//	myVariable := myStruct{}
//	num1 := 1
//	num2 := 2
//	res, err := myVariable.GetSum(num1, num2)
//	if err != nil{
//		panic(err)
//	}else {
//		fmt.Printf("Sum of %d and %d is  %d\n", num1, num2, res)
//	}
//}
