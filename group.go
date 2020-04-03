package helpers

// GroupRectangles returns slice of groups of intersecting rectangles
func GroupRectangles(rects []Rectangle) [][]Rectangle {
	intersections := make(map[int]([]int))
	for i := range rects {
		intersections[i] = []int{}
	}

	for i := range rects {
		for j := range rects {
			if rects[i].Equal(rects[j]) {
				continue
			}

			intersects, _ := rects[i].Intersect(rects[j])
			if intersects {
				intersections[i] = append(intersections[i], j)
			}
		}
	}

	var uniqueNodes []int
	var groups [][]Rectangle
	for node := range intersections {
		if !in(node, uniqueNodes) {
			var subgraph []int
			subgraph, uniqueNodes = getSubgraph(node, intersections, uniqueNodes)
			group := make([]Rectangle, len(subgraph))
			for i, v := range subgraph {
				group[i] = rects[v]
			}

			groups = append(groups, group)
		}
	}

	return groups
}

func getSubgraph(node int, relations map[int]([]int), uniqueNodes []int) ([]int, []int) {
	if !in(node, uniqueNodes) {
		uniqueNodes = append(uniqueNodes, node)
		children := relations[node]

		var chain []int
		for _, child := range children {
			var subgraph []int
			subgraph, uniqueNodes = getSubgraph(child, relations, uniqueNodes)
			chain = append(chain, subgraph...)
		}

		return append(chain, node), uniqueNodes
	}

	return []int{}, uniqueNodes
}

func in(i int, slice []int) bool {
	for _, el := range slice {
		if i == el {
			return true
		}
	}

	return false
}
