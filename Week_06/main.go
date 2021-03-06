package main

import (
	"fmt"
	"week_06/lc221"
	"week_06/lc64"
)

func main() {
	fmt.Println(lc64.MinPathSum([][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {4, 5, 6, 7}}))
	fmt.Println(lc221.MaximalSquare([][]byte{{'1', '1', '1', '0', '0', '1'}, {'1', '1', '1', '0', '0', '1'}, {'1', '1', '1', '0', '0', '1'}, {'1', '1', '1', '0', '0', '1'}, {'1', '1', '1', '0', '0', '1'}}))
}
