package tree

import "sort"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node

	// feel free to add fields as you see fit
}

func AddSortedChildren(parent *Node, newChild *Node) {
	//x := sort.Search(newChild.ID, func(i int) bool { return newChild.ID >= parent.Children[i].ID })
	parent.Children = append(parent.Children, newChild)
	//copy(parent.Children[x+1:], parent.Children[x:])
	//parent.Children[x] = newChild
	sort.Slice(parent.Children, func(i, j int) bool { return parent.Children[i].ID < parent.Children[j].ID })
}

func Build(records []Record) (*Node, error) {
	nodeMap := make(map[int]*Node)

	for _, r := range records {

		// Add key if it is not present already.
		n, ok := nodeMap[r.ID]
		if !ok {
			n = &Node{ID: r.ID}
			nodeMap[r.ID] = n
		}
		// If root key, no need for anything else except entry
		// into nodeMap.
		if r.ID == 0 {
			continue
		}

		// Add parent key if required, and update it.
		val, ok := nodeMap[r.Parent]
		if ok {
			// Don't add duplicate child entries.
			for _, c := range val.Children {
				if n.ID == c.ID {
					panic("Duplicate error.")
				}
			}
			AddSortedChildren(val, n)
			//val.Children = append(val.Children, n)
		} else {
			p := &Node{ID: r.Parent}
			nodeMap[r.Parent] = p
			p.Children = append(p.Children, n)
		}
	}
	return nodeMap[0], nil
}
