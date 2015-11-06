package sensorserver

import (
	"bytes"
	"math"
	"math/rand"
	"testing"
	"time"
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

func TestConvertFloat(t *testing.T) {
	a := []float32{math.MaxFloat32, math.SmallestNonzeroFloat32, 0.0, -0.0}
	for _, x := range a {
		xBytes := Float32Bytes(x)
		xFloat := BytesToFloat32(xBytes)

		if xFloat != x {
			t.Errorf("%f != %f", xFloat, x)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		x := r.Float32()
		xBytes := Float32Bytes(x)
		xFloat := BytesToFloat32(xBytes)

		if xFloat != x {
			t.Errorf("%f != %f", xFloat, x)
		}
	}
}

func TestConvertInt(t *testing.T) {
	a := []int{math.MaxInt32, math.MinInt32, 0, -40, -145}
	for _, x := range a {
		xBytes := IntBytes(x)
		xInt := BytesToInt(xBytes)

		if xInt != x {
			t.Errorf("%d != %d", xInt, x)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		x := int(r.Int31())
		xBytes := IntBytes(x)
		xInt := BytesToInt(xBytes)

		if xInt != x {
			t.Errorf("%d != %d", xInt, x)
		}
	}
}
