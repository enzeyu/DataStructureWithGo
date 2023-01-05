package pQueue

// 返回队列优先级最高元素
func GetMax(data []int) int {
	return data[1]
}

// 返回队列优先级最高元素，将该元素删除
func ExtractMax(data []int) int{
	max := data[1]
	data[1] = data[len(data)-1] // 将底层最右边元素放到根节点上
	data = data[:len(data)-1]
	heapFix(data,1,len(data))
	return max
}

// 将索引为i的元素优先级提升到key级
func IncreaseKey(data []int,i,key int) {
	data[i] = key
	for i > 1 && data[i] > data[getParent(i)]{
		data[i], data[getParent(i)] = data[getParent(i)], data[i]
		i = getParent(i)
	}
}

// 插入一个优先级为key的元素
func Insert(data []int,key int) []int{
	data = append(data,-1<<63)
	data[len(data)-1] = key
	IncreaseKey(data,len(data)-1,key)
	return data
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

func Build(data []int,i,heapSize int){
	left := getLeft(i)
	right := getRight(i)
	if left <= heapSize{
		Build(data,left,heapSize)
	}
	if right <= heapSize{
		Build(data,right,heapSize)
	}
	heapFix(data,i,heapSize)
}