package main

import (
	"binaryTree/tree"
	"fmt"
)

func main(){
	root := tree.NewNode(5)
	fmt.Println(root.FindMinValue(root))
	fmt.Println(root.FindMaxValue(root))
	root.DeleteValue(root,1)
	fmt.Println("层序遍历结果为")
	fmt.Println(root.LevelOrder(root))
	root.AddNode(1)
	root.AddNode(6)
	root.AddNode(4)
	root.AddNode(7)
	root.AddNode(-22)

	fmt.Println("层序遍历结果为")
	fmt.Println(root.LevelOrder(root))
	fmt.Println("搜索结果为")
	fmt.Println(root.SerachValue(root,-1))
	fmt.Println("树里的元素为")
	fmt.Println(root.GetAll())
	fmt.Println("树的最小值为")
	fmt.Println(root.FindMinValue(root))
	fmt.Println("树的最大值为")
	fmt.Println(root.FindMaxValue(root))
	root.DeleteValue(root,1)
	fmt.Println("层序遍历结果为")
	fmt.Println(root.LevelOrder(root))
}