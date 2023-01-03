package merge


func merge(data []int,l,m,r int) {
	n1 := m-l+1 // 左半部分长度
	n2 := r-m // 右半部分长度
	array1 := make([]int,n1+1)
	array2 := make([]int,n2+1)
	maxNumber := 1<<63-1
	array1[n1],array2[n2] = maxNumber,maxNumber
	for i := 0; i < n1; i++{
		array1[i] = data[i+l]
	}
	for i := 0; i < n2; i++{
		array2[i] = data[i+m+1] // +1确保在长度为2的情况下（下标为0），array2的元素和array1不同
	}
	i,j := 0,0
	for k := l; k <= r; k++{
		if array1[i] < array2[j]{
			data[k] = array1[i]
			i++
		}else{
			data[k] = array2[j]
			j++
		}
	}
}

func MergeSort(data []int,l,r int){
	if l < r{
		m := (l+r)/2
		MergeSort(data,l,m)
		MergeSort(data,m+1,r)
		merge(data,l,m,r)
	}
}
