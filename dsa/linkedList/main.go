package main

type node struct {
	data int
	next *node
}

type linkedList struct {
	data node
	next
}
