package main

import (
	"fmt"
	"sort"
)

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

// func Build(records []Record) *Node {

// 	var tree = &Node{}
// 	for _, val := range records {
// 		fmt.Println(val.Parent)

// 		if val.ID == 0 && val.Parent == 0 {
// 			tree.ID = val.ID
// 		} else {
// 			tree.Children = append(tree.Children, &Node{ID: val.ID})
// 		}
// 	}

// 	fmt.Println(tree)
// 	return &Node{
// 		ID: 0,
// 		Children: []*Node{
// 			{ID: 1},
// 			{ID: 2},
// 		},
// 	}
// }

//String converts a record to a string.
func (r *Record) String() string {
	return fmt.Sprintf("(ID: %d, Parent: %d)", r.ID, r.Parent)
}

//Interfaces for storting records.
type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		if r == 0 && (rec.ID != 0 || rec.Parent != 0) {
			return nil, fmt.Errorf("Not a valid root record %s", rec.String())
		} else if r == 0 {
			continue
		} else if r != rec.ID || rec.ID <= rec.Parent {
			return nil, fmt.Errorf("Not a valid record: %s", rec.String())
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}

func main() {

	records := []Record{
		{ID: 5, Parent: 1},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2},
	}
	Build(records)
}
