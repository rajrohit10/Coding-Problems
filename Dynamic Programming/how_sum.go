package main

func HowSum(val int, param []int, myMap map[int]int) int {
	//Optimised
	if val == 0 {
		return 1
	}
	if val < 0 {
		return 0
	}
	if _, exists := myMap[val]; exists {
		return myMap[val]
	}
	counter := 0
	for _, x := range param {
		diff := val - x
		//------------------------
		//The commented Part was earliar there. it was wrong
		// myMap[val] = HowSum(diff, param, myMap)
		// counter=counter+myMap[val]
		//------------------------
		counter += HowSum(diff, param, myMap)

	}
	myMap[val] = counter
	return myMap[val]
}

func HowSumTab(val int, param []int) int {

	newArr:=make([]int,val+1)
	newArr[0]=1

	for i,_:=range newArr{

		if  newArr[i]>0 {

			for _,data:=range param{

				if i+data<=val{
	
					newArr[data+i]+=newArr[i]
				}
	
	
			}

		}
	}
	return newArr[val]







}