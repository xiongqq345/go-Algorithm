package backtracking

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	for h := uint(0); h < 12; h++ {
		for m := uint(0); m < 60; m++ {
			if bits.OnesCount(h)+bits.OnesCount(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}
