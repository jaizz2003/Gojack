package main

import (
	"fmt"
	"os"
)

func add () {
	fmt.Println("Input the number of variables you have")
	var num int
	fmt.Scanf("%d", &num)
	arr := make([]int, num)
	for i := 0; i<num ;i++{
		fmt.Scanf("%d", &arr[i])
	}

	sum := 0
	for i := 0; i<num ;i++{
		sum+=arr[i]
	}

	fmt.Println(sum)
}

func main(){
	fmt.Println("Choose what you want? \n1. Add  \n2. Subtract \n3. Multiply \n4. Divide \n5. Exit")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice{

		case 1: add()
		case 2: subtract()
		case 3: multiply()
		case 4: divide()
		default: os.Exit(3)
			
	}
}

