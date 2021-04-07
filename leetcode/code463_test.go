package leetcode

import "testing"

func TestIslandPerimeter(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	ret := islandPerimeter(grid)
	if ret != 16 {
		t.Errorf("islandPerimeter return %d expected 16", ret)
	} else {
		t.Log("success!")
	}
}
