package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Skiplist(t *testing.T) {
	list := newSkipList()
	for i := 0; i < 10; i++ {
		list.insert(rand.Intn(20), 1)
	}
	list.Print()
	//list.search(15)

	fmt.Println("\n--------------------------------------")

	//list.delete(10)
	//list.Print()
	//
	//fmt.Println("\n--------------------------------------")
}
