package main

import (
	"fmt"
	"quickSort/quick"
)

func main(){
	data := []int{50,4,30,15,21,12,3,18}
	fmt.Println(quick.QuickSort(data,0,len(data)-1))
}

//func swap(data []int, a,b int){
//	data[a],data[b] = data[b],data[a]
//}
