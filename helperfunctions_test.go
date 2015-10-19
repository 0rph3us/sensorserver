package sensorserver

import (
	"bytes"
	"testing"
)

func TestIntBytes(t *testing.T) {
	a := []int{1444349405, 1421674205, 1435306205, 1418333405, 1435814102, 1429574102}

	for _, x := range a {
		for _, y := range a {
			bx := IntBytes(x)
			by := IntBytes(y)

			var want int
			if x < y {
				want = -1
			} else if x == y {
				want = 0
			} else {
				want = 1
			}

			if c := bytes.Compare(bx, by); c != want {
				t.Errorf("%v < %v isn't %d < %d: %d", bx, by, x, y, c)
			}
		}
	}
}
