package quick

func QuickSort(data []int,l,r int) []int{
	if l < r{
		pivot := partition(data,l,r)
		QuickSort(data,l,pivot-1)
		QuickSort(data,pivot+1,r)
	}
	return data
}

// 按pivot将data分为三个部分，小于pivot的部分，pivot和大于pivot的部分
func partition(data []int,l,r int) int{
	pivot := data[r] // 选中pivot
	i := l-1 // i是小于pivot元素区域的最右端, i+1是第一个大于pivot的元素
	for j := l; j < r; j++{
		if data[j] < pivot{ // 找到小的则将i和j指向元素互换
			i++
			data[i],data[j] = data[j],data[i]
		}
	}
	data[i+1],data[r] = data[r],data[i+1] // 遍历后将pivot和i+1
	return i+1
}