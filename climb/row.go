package climb

// TODO: Change nodes to a array of 4 nodes.
type Row struct {
	Prev  *Row
	Next  *Row
	Nodes *Node
}
