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
	if head == nil{
		return nil
	}
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

// 先找到要删除的节点
// 如果要删除的节点同时有左子树和右子树，则将这个值变成右子树最小值，再删去原来的右子树最小值
// 如果要删除节点没有右子树，直接把左子树提上来即可，相反也是同理
func (n *Node) DeleteValue(head *Node, value int) *Node{
	if head == nil{
		return nil
	}
	if value > head.Value{
		head.Right = head.Right.DeleteValue(head.Right,value)
	}else if value < head.Value{
		head.Left = head.Left.DeleteValue(head.Left,value)
	}else if head.Left != nil && head.Right != nil{
		head.Value = head.FindMinValue(head.Right)
		head.Right = head.Right.DeleteValue(head.Right,head.Value)
	}else if head.Left != nil{
		head = head.Left
	}else{
		head = head.Right
	}
	return head
}

func (n *Node) FindMinValue(head *Node) int{
	if head.Left != nil{
		return head.Left.FindMinValue(head.Left)
	}
	return head.Value
}

func (n *Node) FindMaxValue(head *Node) int{
	if head.Right != nil{
		return head.Right.FindMaxValue(head.Right)
	}
	return head.Value
}

// 获得二叉树所有元素
func (n *Node) GetAll() []int{
	ans := make([]int,0)
	return n.getElements(ans)
}

// 获得所有元素的子程序
func (n *Node) getElements(eles []int) []int{
	if n != nil{
		eles = append(n.Left.getElements(eles))
		eles = append(eles,n.Value)
		eles = append(n.Right.getElements(eles))
	}
	return eles
}

