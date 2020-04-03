package helpers

import (
	"errors"
	"reflect"
	"sort"
)

// GetTotalArea calculates total area of rectangles
func GetTotalArea(rects []Rectangle) (int, error) {
	nonZeroRects := Filter(rects, func(rect Rectangle) bool {
		return rect.Area() != 0
	})

	dividers := getUniqueSortedXSlice(nonZeroRects)
	splitted, err := splitRectsByX(rects, dividers)
	if err != nil {
		return 0, err
	}
	combined, err := combineRectsOnY(splitted, dividers)
	if err != nil {
		return 0, err
	}

	totalArea := 0
	for i := range combined {
		totalArea += combined[i].Area()
	}

	return totalArea, nil
}

func (r *Rectangle) splitAtX(x int) []Rectangle {
	if x <= r.X1 || x >= r.X2 {
		return []Rectangle{*r}
	}

	r1 := Rectangle{
		X1: r.X1,
		Y1: r.Y1,
		X2: x,
		Y2: r.Y2,
	}

	r2 := Rectangle{
		X1: x,
		Y1: r.Y1,
		X2: r.X2,
		Y2: r.Y2,
	}

	return []Rectangle{r1, r2}
}

func getUniqueSortedXSlice(rects []Rectangle) []int {
	uniqueXMap := make(map[int]bool)

	for _, r := range rects {
		uniqueXMap[r.X1] = true
		uniqueXMap[r.X2] = true
	}

	uniqueX := make([]int, len(uniqueXMap))
	i := 0
	for k := range uniqueXMap {
		uniqueX[i] = k
		i++
	}

	sort.Ints(uniqueX)

	return uniqueX
}

func splitRectsByX(rects []Rectangle, dividers []int) ([]Rectangle, error) {
	for i := range rects {
		if rects[i].Area() <= 0 {
			return nil, errors.New("one of the Rectangles has non-positive area")
		}
	}
	totalDivided := rects

	for _, x := range dividers {
		var divided []Rectangle
		for _, r := range totalDivided {
			splitted := r.splitAtX(x)
			divided = append(divided, splitted...)
		}
		totalDivided = divided
	}

	return totalDivided, nil
}

func combineRectsOnY(rects []Rectangle, dividers []int) ([]Rectangle, error) {
	var totalCombined []Rectangle

	for _, x := range dividers {
		filtered := Filter(rects, func(rect Rectangle) bool {
			return rect.X1 == x
		})

		sort.Slice(filtered, func(i int, j int) bool {
			return filtered[i].Y1 < filtered[j].Y1
		})

		if len(filtered) == 0 {
			continue
		}

		first := filtered[0]
		last := filtered[len(filtered)-1]
		if reflect.DeepEqual(first, last) {
			totalCombined = append(totalCombined, first)
			continue
		}

		prev := first
		for _, r := range filtered {
			if prev.Width() != r.Width() {
				return nil, errors.New("two rectangles with same X has different width")
			}

			intersects, _ := prev.Intersect(r)
			if intersects {
				prev = *prev.Union(r)
			} else {
				totalCombined = append(totalCombined, prev)
				prev = r
			}
		}

		totalCombined = append(totalCombined, prev)
	}

	return totalCombined, nil
}
