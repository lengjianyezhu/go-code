package main

import "fmt"

/**
题目描述：
	给定一个n个数的数组（n <= 10,000,000）,
	以及一个数字k，请输出：数组中出现最频繁的k个数。
	如：数组[2,3,1,5,2,1,2,4,3,2,3]， k=3
	出现最频繁的数分别是2和3（2出现4次，3出现3次），其次是1（出
	现2次） 所以输出1,2,3这3个数即可，输出顺序随意。

思路 :
	这是一个topk问题，用数组实现的小根堆解决
**/
func main() {
	//测试用的数组
	testSlice := []int{2, 3, 11, 4, 5, 6, 2, 1, 123, 43, 1, 32, 3, 5, 2, 1, 8, 54, 4, 0}
	//K指TOPK中的K
	K := 5

	//初始化小根堆，这里用数组存储小根堆
	arr := testSlice[0:K]

	//获取最后一个有子节点的索引位置
	index := len(arr)/2 - 1

	//生成小根堆
	for i := index; i >= 0; i-- {
		head(arr, i)
	}
	// fmt.Println(arr)

	//开始遍历剩下的所有元素
	for i := K; i < len(testSlice); i++ {
		//每遍历一个则跟堆顶元素进行比较大小
		//如果比栈顶元素大，就替换栈顶元素，重新排列小根堆
		if testSlice[i] > arr[0] {
			arr[0] = testSlice[i]
			head(arr, 0)
		}
	}
	fmt.Printf("前%d个最大的值:%v", K, arr)
}

//表示从 index对应的下标开始调节 小根堆的 子节点
func head(input []int, index int) {
	//index元素对应的左节点
	left := index<<1 + 1
	//index 元素对应的右节点
	right := index<<1 + 2
	var min, temp int
	//下标越界就返回
	if left >= len(input) {
		return
	}
	//min表示 子节点种最小值 对应的下标
	if (right < len(input)) && input[right] < input[left] {
		min = right
	} else {
		min = left
	}
	//父节点 和 子节点的交换
	if input[index] > input[min] {
		temp = input[index]
		input[index] = input[min]
		input[min] = temp
		head(input, min)
	}
}
