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
	Y1: 4,
	X2: 6,
	Y2: 6,
}

func TestArea(t *testing.T) {
	assert := assert.New(t)
	area := rect1.Area()

	assert.Equal(9, area)
}

func TestGetRectUnion(t *testing.T) {
	assert := assert.New(t)
	union := GetRectUnion(rect1, rect2)

	correct := Rectangle{1, 2, 5, 7}

	assert.Equal(correct, *union)
}

func TestGetRectIntersection(t *testing.T) {
	assert := assert.New(t)
	intersects, intersection := GetRectIntersection(rect1, rect2)

	correct := Rectangle{2, 3, 4, 5}

	assert.True(intersects)
	assert.Equal(correct, *intersection)

	intersects, intersection = GetRectIntersection(rect1, rect3)

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
	assert.Equal(25, area)
}
