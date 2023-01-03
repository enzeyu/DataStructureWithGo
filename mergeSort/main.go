package main

import "fmt"
import "mergeSort/merge"

func main(){
	data := []int{50,4,30,15,21,12,3,18}
	merge.MergeSort(data,0,len(data)-1)
	fmt.Println(data)
}
