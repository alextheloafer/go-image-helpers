// Package helpers provides types, functions and methods
// that often used in image processing
package helpers

import (
	"math"
)

// BBoxToRect transforms BBox coordinates to Rectangle
func BBoxToRect(tlx, bry, w, h int) *Rectangle {
	return &Rectangle{
		X1: tlx - int(float64(w)/2),
		Y1: bry - int(float64(h)/2),
		X2: tlx + int(float64(w)/2),
		Y2: bry + int(float64(h)/2),
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

// Rectangle is a struct that implements rectangle
type Rectangle struct {
	X1, Y1, X2, Y2 int
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() int {
	return r.Width() * r.Height()
}

// Equal returns if rectangle equal r1 or not
func (r *Rectangle) Equal(r1 Rectangle) bool {
	return r.X1 == r1.X1 && r.X2 == r1.X2 && r.Y1 == r1.Y1 && r.Y2 == r1.Y2
}

// Height returns height of the rectangle
func (r *Rectangle) Height() int {
	return r.Y2 - r.Y1
}

// Intersect returns if two rectangles intersects
// and intersection rectangle.
func (r *Rectangle) Intersect(r1 Rectangle) (bool, *Rectangle) {
	x1 := int(math.Max(float64(r.X1), float64(r1.X1)))
	x2 := int(math.Min(float64(r.X2), float64(r1.X2)))
	y1 := int(math.Max(float64(r.Y1), float64(r1.Y1)))
	y2 := int(math.Min(float64(r.Y2), float64(r1.Y2)))
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

// Union returns union of two rectangles
func (r *Rectangle) Union(r1 Rectangle) *Rectangle {
	return &Rectangle{
		X1: int(math.Min(float64(r.X1), float64(r1.X1))),
		X2: int(math.Max(float64(r.X2), float64(r1.X2))),
		Y1: int(math.Min(float64(r.Y1), float64(r1.Y1))),
		Y2: int(math.Max(float64(r.Y2), float64(r1.Y2))),
	}
}

// Width returns width of the rectangle
func (r *Rectangle) Width() int {
	return r.X2 - r.X1
}
