package main

import "fmt"

func printArray2(arr []int) {
	for index, value := range arr {
		fmt.Printf("%d \t", value)
		fmt.Printf("%d \n", index)
	}
}
