/*
使用二叉树排序Slice
*/
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// 前序遍历二叉树
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 建立二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		// 法1：
		//t = new(tree)
		//t.value = value

		// 法2：
		t = &tree{value: value}
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	arr := []int{1024, 10, 9, 8, 11, 34, 1, 3, 2, 11}
	Sort(arr)
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}
