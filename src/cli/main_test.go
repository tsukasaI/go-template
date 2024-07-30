package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestSampleMax(t *testing.T) {
	tests := map[string]struct {
		a    int
		b    int
		want int
	}{
		"1 and 2": {
			a:    1,
			b:    2,
			want: 2,
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			res := SampleMax(tt.a, tt.b)
			assert.Equal(t, tt.want, res)
		})
	}
}
