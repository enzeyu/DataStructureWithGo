package main

import (
	"binaryTree/tree"
	"fmt"
)

func main(){
	root := tree.NewNode(5)
	root.AddNode(1)
	root.AddNode(6)
	root.AddNode(4)
	root.AddNode(7)

	fmt.Println(root.LevelOrder(root))
	fmt.Println(root.SerachValue(root,5))
	fmt.Println(root.GetAll())
}