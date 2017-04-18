// https://leetcode.com/problems/hamming-distance/#/description
// The Hamming distance between two integers is the number of positions at which
// the corresponding bits are different.
//
// Given two integers x and y, calculate the Hamming distance.
//
// Note:
// 0 ≤ x, y < 2^31.

package main

import "fmt"

func hammingDistance(x int, y int) int {
	var hamming_dist int
	z := x ^ y
	for {
		if v := z & 1; v == 1 {
			hamming_dist += 1
		}
		z = z >> 1
		if z == 0 {
			break
		}
	}
	return hamming_dist
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
