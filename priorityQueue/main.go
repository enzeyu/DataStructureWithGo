package main

import (
	"fmt"
	"priorityQueue/pQueue"
)

func main(){
	data := []int{-1,6,5,3,30,50,21,12,18,24,20}
	pQueue.Build(data,1,len(data)-1) // 将data初始化为优先队列
	fmt.Println(pQueue.ExtractMax(data))
	fmt.Println(data)
	data = pQueue.Insert(data,1000)
	fmt.Println(data)
	pQueue.IncreaseKey(data,2,1001)
	fmt.Println(data)
}

