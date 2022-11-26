package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// AddSortedChildren adds a new child and then sorts children based on ID.
func AddSortedChildren(parent *Node, newChild *Node) {
	parent.Children = append(parent.Children, newChild)
	sort.Slice(parent.Children,
		func(i, j int) bool { return parent.Children[i].ID < parent.Children[j].ID })
}

// ChildAlreadyPresent checks if a child is already present in the parent node.
func ChildAlreadyPresent(parent *Node, newChild *Node) bool {
	for _, c := range parent.Children {
		if newChild.ID == c.ID {
			return true
		}
	}
	return false
}

// ContinuityCheck checks if the tree has a continuous sequence of record IDs.
func ContinuityCheck(nodeMap map[int]*Node, recordLen int) error {
	maxKey := 0
	for k := range nodeMap {
		if k > maxKey {
			maxKey = k
		}
	}
	if recordLen != 1+maxKey {
		return fmt.Errorf("non-continuous records")
	}

	return nil
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	rootSeen := false
	nodeMap := make(map[int]*Node)

	for _, r := range records {

		// Add key if it is not present already.
		n, ok := nodeMap[r.ID]
		if !ok {
			n = &Node{ID: r.ID}
			nodeMap[r.ID] = n
		} else {
			if r.ID == 0 && rootSeen {
				return nil, fmt.Errorf("duplicate root detected")
			}
		}

		// If root key, ensure no parent and continue with next record.
		if r.ID == 0 {
			rootSeen = true
			if r.Parent != 0 {
				return nil, fmt.Errorf("root cannot have a parent %d", r.Parent)
			} else {
				continue
			}
		}

		// Parent ID must be lower than child ID.
		if r.ID <= r.Parent {
			return nil, fmt.Errorf("parent id %d must be greater than child id %d", r.Parent, r.ID)
		}

		// Add parent key if required, and update its children.
		parent, ok := nodeMap[r.Parent]
		if ok {
			// Don't add duplicate child entries.
			if ChildAlreadyPresent(parent, n) {
				return nil, fmt.Errorf("child with ID %d to parent %d already exists", n.ID, parent.ID)
			}
			AddSortedChildren(parent, n)
		} else {
			p := &Node{ID: r.Parent}
			nodeMap[r.Parent] = p
			p.Children = append(p.Children, n)
		}
	}

	root, ok := nodeMap[0]
	if !ok {
		return nil, fmt.Errorf("tree not connected to root node")
	}

	if !rootSeen {
		return nil, fmt.Errorf("root node not explicitly specified")
	}

	err := ContinuityCheck(nodeMap, len(records))
	if err != nil {
		return nil, err
	}

	return root, nil
}
