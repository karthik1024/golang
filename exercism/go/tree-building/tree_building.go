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

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	var nodeList = make([]Node, len(records))
	//var nodeList []Node

	var parent *Node
	for idx, rec := range records {
		if rec.ID == 0 && rec.Parent != 0 {
			return nil, fmt.Errorf("incomplete record set")
		}

		if (idx != rec.ID) || ((rec.ID != 0) && rec.ID <= rec.Parent) {
			return nil, fmt.Errorf("incomplete record set")
		}

		n := Node{ID: rec.ID}

		nodeList[idx] = n
		//nodeList = append(nodeList, n)

		parent = &nodeList[rec.Parent]
		if rec.ID != 0 {
			parent.Children = append(parent.Children, &nodeList[idx])
		}
	}

	return &nodeList[0], nil
}
