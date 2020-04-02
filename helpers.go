package helpers

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

// Rectangle is a struct that implements rectangle
type Rectangle struct {
	X1, Y1, X2, Y2 int
}

// BBoxToRect transforms BBox coordinates to Rectangle
func BBoxToRect(tlx, bry, w, h int) *Rectangle {
	return &Rectangle{
		X1: tlx - int(float64(w)/2),
		Y1: bry - int(float64(h)/2),
		X2: tlx + int(float64(w)/2),
		Y2: bry + int(float64(h)/2),
	}
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() int {
	return r.Width() * r.Height()
}

// Width returns width of the rectangle
func (r *Rectangle) Width() int {
	return r.X2 - r.X1
}

// Height returns height of the rectangle
func (r *Rectangle) Height() int {
	return r.Y2 - r.Y1
}

// GetRectUnion returns union of two rectangles
func GetRectUnion(r1, r2 Rectangle) *Rectangle {
	return &Rectangle{
		X1: int(math.Min(float64(r1.X1), float64(r2.X1))),
		X2: int(math.Max(float64(r1.X2), float64(r2.X2))),
		Y1: int(math.Min(float64(r1.Y1), float64(r2.Y1))),
		Y2: int(math.Max(float64(r1.Y2), float64(r2.Y2))),
	}
}

// GetRectIntersection returns if two rectangles intersects
// and intersection rectangle.
func GetRectIntersection(r1, r2 Rectangle) (bool, *Rectangle) {
	x1 := int(math.Max(float64(r1.X1), float64(r2.X1)))
	x2 := int(math.Min(float64(r1.X2), float64(r2.X2)))
	y1 := int(math.Max(float64(r1.Y1), float64(r2.Y1)))
	y2 := int(math.Min(float64(r1.Y2), float64(r2.Y2)))
	if x1 >= x2 || y1 >= y2 {
		return false, nil
	}
	return true, &Rectangle{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

// Filter filters array of rectangles based on condition
func Filter(rects []Rectangle, filterFunc func(rect Rectangle) bool) []Rectangle {
	var filtered []Rectangle
	for i := range rects {
		if filterFunc(rects[i]) {
			filtered = append(filtered, rects[i])
		}
	}

	return filtered
}

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
	for k := range uniqueXMap {
		uniqueX = append(uniqueX, k)
	}

	sort.Ints(uniqueX)

	return uniqueX
}

func splitRectsByX(rects []Rectangle, dividers []int) ([]Rectangle, error) {
	var totalDivided []Rectangle
	for i := range rects {
		if rects[i].Area() <= 0 {
			return nil, errors.New("one of the Rectangles has non-positive area")
		}
	}

	for _, x := range dividers {
		var divided []Rectangle
		for _, r := range rects {
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
		last := filtered[0]
		if reflect.DeepEqual(first, last) {
			totalCombined = append(totalCombined, first)
			continue
		}

		prev := first
		for _, r := range filtered {
			if prev.Width() != r.Width() {
				return nil, errors.New("two rectangles with same X has different width")
			}

			intersects, _ := GetRectIntersection(prev, r)
			if intersects {
				prev = *GetRectUnion(prev, r)
			} else {
				totalCombined = append(totalCombined, prev)
				prev = r
			}
		}

		totalCombined = append(totalCombined, prev)
	}

	return totalCombined, nil
}
