package rest

// SkipList is a data structure for creating ordered Sets/Maps/Lists
// It has the same performance/complexity as a B-Tree, O(log n)
// in search, insert, remove, update.
// However, it is much simpler to implement than a B-Tree
// To learn more about Skip lists see: http://epaperpress.com/sortsearch/download/skiplist.pdf
import (
	"math/rand"
	"time"
)

// 32 levels counting from 0
const maxHeight = 31

// skipList struct itself
type skipList struct {
	height int
	head   *skipListNode
}

// A skiplist's node representation
type skipListNode struct {
	ttl  time.Time       // Used to compare nodes and sort them
	key  string          // Used to remove elements from other structures such as lrus and maps
	next []*skipListNode // Used to poit to the next nodes, one for each level
}

func newSkipList() *skipList {
	head := &skipListNode{
		next: make([]*skipListNode, maxHeight),
	}

	return &skipList{
		height: 0,
		head:   head,
	}
}

// Insert a node in the SkipList
func (s *skipList) insert(key string, ttl time.Time) *skipListNode {
	level := 0

	// New random seed
	rand.Seed(time.Now().UnixNano())

	// Like flipping a coin up to the maximum height
	// Level will have a value between 0 and 31
	for level < maxHeight && rand.Intn(2) == 1 {
		level++

		if level > s.height {
			s.height = level
			break
		}
	}

	node := &skipListNode{
		ttl:  ttl,
		key:  key,
		next: make([]*skipListNode, level+1),
	}

	current := s.head

	for i := s.height; i >= 0; i-- {

		for ; current.next[i] != nil; current = current.next[i] {
			// If the ttl of the next element is > than the element to be inserted
			// go down one level.
			if current.next[i].ttl.Sub(node.ttl) > 0 {
				break
			}
		}

		// Just care if we are at the right level or less
		if i <= level {
			node.next[i] = current.next[i]
			current.next[i] = node
		}
	}

	return node
}

// Remove a node from the SkipList
func (s *skipList) remove(node *skipListNode) {
	if node == nil {
		return
	}

	current := s.head

	for i := s.height; i >= 0; i-- {

		// If next is nil, move to the next level
		for ; current.next[i] != nil; current = current.next[i] {

			// If the ttl of the next element is > than the element to be removed,
			// go down one level
			if current.next[i].ttl.Sub(node.ttl) > 0 {
				break
			}

			// If current next points to the node we are trying to remove,
			// change pointers, so current.next will point to node.next
			if current.next[i] == node {
				current.next[i] = node.next[i]
				break
			}

		}

	}
}
