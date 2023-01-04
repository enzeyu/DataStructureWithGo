package main

import "fmt"
import "heap/heapWithArray"

func main(){
	// data是原始数据, -1是哨兵节点，经过处理后输出堆的结构
	data := []int{-1,6,5,3,30,50,21,12,18,24,20}
	max := 1<<63-1
	heapWithArray.ConstructHeap(data)
	fmt.Println(data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
	max,data = heapWithArray.GetMax(data)
	fmt.Println(max,data)
}
