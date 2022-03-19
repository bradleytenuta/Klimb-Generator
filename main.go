package main

import (
	"fmt"
	"generator/climb"
	"generator/common"
)

/*
T ----
When creating a struct, if you do not declare a variable that is in the struct, it will be nil.
T ----
*/
func main() {
	controller := common.Controller{}

	node1 := climb.Node{
		Type: "Standard",
	}
	row1 := climb.Row{
		Nodes: &node1,
	}

	node2 := climb.Node{
		Type: "Type2",
	}
	row2 := climb.Row{
		Nodes: &node2,
	}

	node3 := climb.Node{
		Type: "Type3",
	}
	row3 := climb.Row{
		Nodes: &node3,
	}

	controller.Insert(row1)
	controller.Insert(row2)
	controller.Insert(row3)

	fmt.Printf("Head: %v\n", controller.Head.Nodes.Type)
	fmt.Printf("Tail: %v\n", controller.Tail.Nodes.Type)
	controller.DisplayTypes()
	controller.Reverse()
}
