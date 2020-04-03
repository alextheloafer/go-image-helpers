package helpers_test

import (
	"testing"

	. "github.com/alextheloafer/go-image-helpers"
	"github.com/stretchr/testify/assert"
)

var rect1 = Rectangle{
	X1: 1,
	Y1: 2,
	X2: 4,
	Y2: 5,
}

var rect2 = Rectangle{
	X1: 2,
	Y1: 3,
	X2: 5,
	Y2: 7,
}

var rect3 = Rectangle{
	X1: 8,
	Y1: 3,
	X2: 11,
	Y2: 5,
}

var rect4 = Rectangle{
	X1: 3,
	Y1: 6,
	X2: 6,
	Y2: 8,
}

func TestArea(t *testing.T) {
	assert := assert.New(t)
	area := rect1.Area()

	assert.Equal(9, area)
}

func TesUnion(t *testing.T) {
	assert := assert.New(t)
	union := rect1.Union(rect2)

	correct := Rectangle{1, 2, 5, 7}

	assert.Equal(correct, *union)
}

func TestGetRectIntersection(t *testing.T) {
	assert := assert.New(t)
	intersects, intersection := rect1.Intersect(rect2)

	correct := Rectangle{2, 3, 4, 5}

	assert.True(intersects)
	assert.Equal(correct, *intersection)

	intersects, intersection = rect1.Intersect(rect3)

	assert.False(intersects)
	assert.Nil(intersection)
}

func TestFilter(t *testing.T) {
	rects := []Rectangle{rect1, rect2, rect3}

	filtered := Filter(rects, func(rect Rectangle) bool {
		return rect.X1 <= 5
	})

	correct := []Rectangle{rect1, rect2}

	assert.Equal(t, correct, filtered)
}

func TestGetTotalArea(t *testing.T) {
	assert := assert.New(t)
	rects := []Rectangle{rect1, rect2, rect3, rect4}

	area, err := GetTotalArea(rects)
	assert.Nil(err)
	assert.Equal(27, area)
}

// func TestGroupRectangles(t *testing.T) {
// 	assert := assert.New(t)
//
// 	rects := []Rectangle{rect1, rect2, rect3, rect4}
//
// 	correctGroup1 := []Rectangle{rect1, rect2, rect4}
// 	correctGroup2 := []Rectangle{rect3}
// 	correct := [][]Rectangle{correctGroup1, correctGroup2}
//
// 	GroupRectangles(rects)
// }
