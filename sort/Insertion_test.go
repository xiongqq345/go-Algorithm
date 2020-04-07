package sort

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"AAA",args{[]int{5,8,1,2,9,5,1,23,4,12,4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insertion(tt.args.arr)
		})
		fmt.Println(tt.args.arr)
	}
}
