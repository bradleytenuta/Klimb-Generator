package common

/*
T ----
The import path is like this:
"<module>/<folder-name>"

The above imports all the go files in the package. One package can exist within a folder. So folders sync with packages.
Then to use an exported object from a go file you use:

<package-name>.<object-name>
T ----
*/
import (
	"fmt"
	"generator/climb"
)

/*
Is a linked list that handles the creation of rows and their sequential order.
This stuct can access the head and tail row. Each row is linked to the one before and after it.
*/
type Controller struct {
	Head *climb.Row
	Tail *climb.Row
}

/*
T ----
A method, a function that belongs to a struct.
T ----

A controller method to insert a new row.
*/
func (controller *Controller) Insert(newRow climb.Row) {

	newRow.Next = controller.Head

	// If the head of the controller is not null, update the rows previous row with this new row.
	if controller.Head != nil {
		controller.Head.Prev = &newRow
	}
	controller.Head = &newRow

	// Gets the row of the head controller.
	// If the head row's next row is not null then set the head row's next row as the controllers tail.
	var headRow = controller.Head
	for headRow.Next != nil {
		headRow = headRow.Next
	}
	controller.Tail = headRow
}

/*
Goes through each row and displays its nodes types.
*/
func (controller *Controller) DisplayTypes() {
	headRow := controller.Head

	for headRow != nil {
		// TODO: Add loop here to loop through nodes when I make it a list.
		fmt.Printf("%+v -> ", headRow.Nodes.Type)
		headRow = headRow.Next
	}
	fmt.Println()
}

func (controller *Controller) Reverse() {
	var previousRow *climb.Row

	// Sets the head row of the controller as the current row and last row in controller.
	currentRow := controller.Head
	controller.Tail = controller.Head

	// Loops through the current Row while its not nil.
	for currentRow != nil {
		nextRow := currentRow.Next
		currentRow.Next = previousRow
		previousRow = currentRow

		// Updates the current row. This will eventually become nil when there is no more rows.
		currentRow = nextRow
	}

	controller.Head = previousRow
	Show(controller.Head)
}

func Show(row *climb.Row) {
	// While the row is not nil
	for row != nil {
		// Prints the type.
		// TODO: loop through the nodes and print all types.
		fmt.Printf("%v ->", row.Nodes.Type)

		// Updates the row with the next row, will eventually become nil when there is no more rows.
		row = row.Next
	}
	fmt.Println()
}
