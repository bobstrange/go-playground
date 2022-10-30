package main

type Node struct {
	name   string
	depth  int
	parent *Node
}

type Status int

const (
	// iota + 1 ではなく int のゼロ値 0 になるようにしておく
	DefaultStatus Status = iota
	ActiveStatus
	CloseStatus
)

type Visitor struct {
	Status Status
}
