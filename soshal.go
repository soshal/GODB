package main

import ("fmt")



func find ( n int, r map[int]int ) bool {

	if r[n]==n{
		return true
	}
	return false

}

func main () {

	arr := []int{1,2,3,4,5,6,7,8}
	m := make(map[int]int)
	for i :=0 ; i < len(arr); i++{
	  m[i] = arr[i]
    }

	fmt.Print(find(5, m) )


			
	} 
