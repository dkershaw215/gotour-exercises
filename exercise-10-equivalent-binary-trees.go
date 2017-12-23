package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"reflect"
	"time"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, done chan bool) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch, nil)
	}
	if t.Right != nil {
		Walk(t.Right, ch, nil)
	}
	if done != nil {
		done <- true
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	done1 := make(chan bool)
	go Walk(t1, ch1, done1)

	ch2 := make(chan int)
	done2 := make(chan bool)
	go Walk(t2, ch2, done2)

	t1Nodes := collectTreeNodes(ch1, done1)
	t2Nodes := collectTreeNodes(ch2, done2)

	fmt.Println(t1Nodes)
	fmt.Println(t2Nodes)
	return reflect.DeepEqual(t1Nodes, t2Nodes)

}

func collectTreeNodes(ch chan int, done chan bool) []int {

	nodes := make([]int, 0)

	for {
		select {
		case i := <-ch:
			nodes = append(nodes, i)
		case <-done:
			sort.Sort(sort.IntSlice(nodes))
			return nodes
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
