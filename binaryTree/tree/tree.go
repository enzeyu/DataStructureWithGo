package tree

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func NewNode(val int) *Node{
	return &Node{Value: val}
}

// 递归添加元素
func (n *Node) AddNode(val int) *Node{
	if n == nil{
		return NewNode(val)
	}
	if n.Value < val{
		n.Right = n.Right.AddNode(val)
	}else{
		n.Left = n.Left.AddNode(val)
	}
	return n
}

// 找元素
func (n *Node) SerachValue(head *Node, val int) *Node {
	if head == nil{
		return nil
	}
	if head.Value == val{
		return head
	}else if val > head.Value{
		return n.SerachValue(head.Right,val)
	}else {
		return n.SerachValue(head.Left,val)
	}
}

// 层序遍历
func (n *Node) LevelOrder(head *Node) [][]int{
	ans := make([][]int,0)
	queue := []*Node{ head }
	for len(queue) > 0{
		temp := make([]int,0)
		length := len(queue)
		for length > 0{
			h := queue[0]
			queue = queue[1:]
			length--
			temp = append(temp,h.Value)
			if h.Left != nil{
				queue = append(queue,h.Left)
			}
			if h.Right != nil {
				queue = append(queue, h.Right)
			}
		}
		ans = append(ans,temp)
	}
	return ans
}

func (n *Node) DeleteValue(head *Node, value int) *Node{
	// 如果删除的元素是根元素，则右子树节点变成根节点
	
	return head
}

// 获得二叉树所有元素
func (n *Node) GetAll() []int{
	ans := make([]int,0)
	return n.getElements(ans)
}

func (n *Node) getElements(eles []int) []int{
	if n != nil{
		eles = append(n.Left.getElements(eles))
		eles = append(eles,n.Value)
		eles = append(n.Right.getElements(eles))
	}
	return eles
}

