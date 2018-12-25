package main

import (
	"errors"
	"fmt"
)

/**
 题目：
	 给定一个二叉树，每个节点上有一个数字。
	 对于每一个叶子节点，输出从该节点到根
	节点路径上所有数的和。
思路：
	实际上是一个遍历问题，递归查找叶子节点，利用回溯的思想，用栈存储 找到节点的路径
    当找到节点后，计算栈中的数据和。
**/
func main() {
	//构建题目中的树
	rootNode := GenerateTree()

	HeadNode := &Node{0, nil, nil}
	//构建的栈
	stack := &Stack{HeadNode, HeadNode}
	//叶子节点对应的值
	target := 13
	//主要逻辑，递归查找叶子节点，同时将查找的路径存入stack
	getLeaveNodeWithPathSum(target, rootNode, stack)
	//遍历 stack
	sum := stack.getAllSum()
	fmt.Printf("叶子节点 %d 对应的和是 %d", target, sum)
}

/**
  说明：查找一个叶子结点，并存储对应的路径
  参数：
	 target :叶子节点的值
	 root : 二叉树的根
	 stack: 存储路径的栈
  返回的值：
	 bool : 表示有没有找到叶子节点
**/
func getLeaveNodeWithPathSum(target int, root *TreeNode, stack *Stack) bool {

	if root == nil {
		return false
	}
	//入栈当前的值
	stack.Push(root.value)

	if root.Left == nil && root.Right == nil && root.value == target {
		//说明找到了对应的子节点，返回true
		return true
	}
	//分别从左子树和右子树种递归去找
	flag1 := getLeaveNodeWithPathSum(target, root.Left, stack)
	flag2 := getLeaveNodeWithPathSum(target, root.Right, stack)

	if !flag1 && !flag2 {
		//如果两个子树都没有找到，需要去掉栈顶的内容
		stack.Pop()
		return false
	}
	return true
}

//节点
type TreeNode struct {
	value int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree() *TreeNode {
	root := &TreeNode{5, nil, nil}

	// 7,2,11
	tempRoot := &TreeNode{11, nil, nil}
	tempRoot.Left = &TreeNode{7, nil, nil}
	tempRoot.Right = &TreeNode{2, nil, nil}
	// 4
	tempRoot = &TreeNode{4, tempRoot, nil}
	//左边构建完成
	root.Left = tempRoot
	//1
	tempRoot = &TreeNode{4, nil, nil}
	tempRoot.Right = &TreeNode{1, nil, nil}

	tempRoot = &TreeNode{8, nil, tempRoot}
	tempRoot.Left = &TreeNode{13, nil, nil}
	root.Right = tempRoot
	return root
}
func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%v=>", root.value)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

//用链表实现一个栈
type Node struct {
	val  int
	next *Node
	prev *Node
}

//栈
type Stack struct {
	Head *Node //栈首
	Tail *Node //初始化为head
}

//入栈
func (this *Stack) Push(val int) {

	newNode := &Node{val, nil, nil}
	//放到最后一个节点后面
	this.Tail.next = newNode
	newNode.prev = this.Tail
	this.Tail = newNode
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//如果到了栈首，就不出栈了
	if this.Tail == this.Head {
		fmt.Println("没有元素可以出栈了")
		return 0, errors.New("没有元素可以出栈了")
	}

	val = this.Tail.val
	this.Tail = this.Tail.prev
	this.Tail.next = nil
	return val, nil
}

//计算栈中所有数据的和
func (this *Stack) getAllSum() int {
	helper := this.Head
	sum := 0
	if this.Head.next == nil {
		return 0
	}
	for {
		//从栈首打印到栈尾
		if helper.next == nil {
			break
		}
		helper = helper.next
		sum += helper.val
	}
	return sum
}
