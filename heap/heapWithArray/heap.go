package heapWithArray

func ConstructHeap(data []int){
	heapSize := len(data)-1
	build(data,1,heapSize) // 为首元素建立堆
}

func build(data []int,i,heapSize int){
	left := getLeft(i)
	right := getRight(i)
	if left <= heapSize{
		build(data,left,heapSize)
	}
	if right <= heapSize{
		build(data,right,heapSize)
	}
	heapFix(data,i,heapSize)
}

func heapFix(data []int,i,heapSize int){
	left,right := getLeft(i),getRight(i)
	temp := i
	if left <= heapSize && data[left] > data[i]{
		temp = left
	}
	if right <= heapSize && data[right] > data[temp]{
		temp = right
	}
	if temp != i{ // temp保存了根节点，左节点和右节点最大的
		data[temp],data[i] = data[i],data[temp]
		heapFix(data,temp,heapSize)
	}
}

// 获得父节点
func getParent(i int) int{
	return i/2
}

// 获得左子节点
func getLeft(i int) int{
	return 2*i
}

// 获得右子节点
func getRight(i int) int{
	return 2*i+1
}

// 从大根堆里每次获取最大值
func GetMax(data []int) (int,[]int){
	max := data[1]
	for i := 1; i < len(data)-1; i++{
		data[i] = data[i+1]
	}
	data = data[:len(data)-1]
	heapFix(data,1,len(data)-1) // 进行修复
	return max,data
}

