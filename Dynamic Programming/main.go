package main

import (
	"fmt"
	
)


func main() {
 
	//----------------------------------Program 1 -----------------------------------//
	//For FIbbonaci
	// map1 := make(map[int]int)
	// fmt.Println(Fibb(50, map1))
	//----------------------------------Program 2 -----------------------------------//
	//For 2D Matrix Traverse

	// type matrixKey struct{
	// 	one int
	// 	two int

	// }
	// map1 := make(map[matrixKey]int)
	// fmt.Println(twoDmatrixTraverse(2,3,matrixKey, map1))

	// map1 := make(map[string]int)
	// fmt.Println(TwoDmatrixTraverse(1000,40000005, map1))
	//----------------------------------Program 3 -----------------------------------//
	//Returning true if numbers from the array add up to be the number given as input
	m := make(map[int]bool)
	var newarr= []int{7,14}
	fmt.Println(CanSum(10000000, newarr,m))
}
