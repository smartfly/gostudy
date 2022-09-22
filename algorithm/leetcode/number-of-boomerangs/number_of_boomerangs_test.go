package number_of_boomerangs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfBoomerangs(t *testing.T) {
	points := [][]int{{0, 0}, {1, 0}, {2, 0}}
	ret := numberOfBoomerangs(points)
	assert.Equal(t, 2, ret)
}
