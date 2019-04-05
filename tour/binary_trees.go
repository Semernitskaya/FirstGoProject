package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	InnerWalk(t, ch)
	close(ch)
}

func InnerWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		InnerWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		InnerWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	if t1 == t2 {
		return true
	}
	channel1 := make(chan int)
	channel2 := make(chan int)
	go Walk(t1, channel1)
	go Walk(t2, channel2)

	for {
		val1, ok1 := <-channel1
		val2, ok2 := <-channel2
		if ok1 != ok2 {
			return false
		}
		if ok1 == false && ok2 == false {
			return true
		}
		if val1 != val2 {
			return false
		}
	}
}

func incMostRight(t *tree.Tree) {
	right := t
	for ; right.Right != nil; right = right.Right {
	}
	right.Value += 1000

}

func main() {
	channel := make(chan int)
	tree0 := tree.New(1)
	go Walk(tree0, channel)
	for value := range channel {
		fmt.Println(value)
	}

	tree1 := tree.New(1)
	tree2 := tree.New(2)
	tree3 := tree.New(2)
	incMostRight(tree3)

	fmt.Println(Same(tree0, tree0))
	fmt.Println(Same(tree0, tree1))
	fmt.Println(Same(tree0, tree2))
	fmt.Println(Same(tree2, tree3))

}
