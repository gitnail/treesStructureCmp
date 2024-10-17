package main

/*
  Code that compares two trees for structure equality.
*/

import (
	"fmt"
)

type node struct {
	id       int
	children []*node
}

func calc(v *node, sz map[int]uint, levels map[uint][]uint, depth uint) {
	if v == nil {
		return
	}

	sz[v.id] = 1

	for _, c := range v.children {
		calc(c, sz, levels, depth+1)
		sz[v.id] += sz[c.id]
	}

	levels[depth] = append(levels[depth], sz[v.id])
}

func main() {
	t1 := &node{
		id: 1,
		children: []*node{
			{id: 2},
			{id: 13},
			{
				children: []*node{
					{id: 4},
					{id: 5},
				},
			},
		},
	}

	t2 := &node{
		id: 1,
		children: []*node{
			{id: 2},
			{id: 3},
			{
				children: []*node{
					{id: 34},
					{id: 55},
				},
			},
		},
	}

	sz1 := make(map[int]uint)
	sz2 := make(map[int]uint)

	levels1 := make(map[uint][]uint)
	levels2 := make(map[uint][]uint)

	calc(t1, sz1, levels1, 0)
	calc(t2, sz2, levels2, 0)

	if len(levels1) != len(levels2) {
		fmt.Println("Not equal")
		return
	}

	for k, v1 := range levels1 {
		v2 := levels2[k]

		if len(v1) != len(v2) {
			fmt.Println("Not equal")
			return
		}

		cnt := make(map[uint]uint)
		for _, sz := range v1 {
			cnt[sz]++
		}
		for _, sz := range v2 {
			if cnt[sz] == 0 {
				fmt.Println("Not equal")
				return
			}

			cnt[sz]--
		}
		for _, sz := range cnt {
			if sz != 0 {
				fmt.Println("Not equal")
				return
			}
		}
	}

	fmt.Println("Equal")
}
