package bug

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptimize(t *testing.T) {
	type args struct {
		v    costs
		w    weights
		bugW int
	}
	type want struct {
		cost    float64
		objects []int
	}
	tests := []struct {
		args args
		want want
	}{
		{args: args{v: []float64{60, 100, 120}, w: []int{1, 2, 3}, bugW: 5}, want: want{cost: 220, objects: []int{2, 1}}},
		{args: args{v: []float64{1, 2, 3, 5, 2}, w: []int{1, 2, 1, 4, 3}, bugW: 7}, want: want{cost: 10, objects: []int{1, 3, 2}}},
	}
	for _, tt := range tests {
		cost, objects := Optimize(tt.args.v, tt.args.w, tt.args.bugW)

		assert.Equal(t, tt.want.cost, cost)
		assert.Equal(t, tt.want.objects, objects)
	}
}
