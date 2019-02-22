package rsb

import (
	"math/rand"
	"time"
)

type Object struct {
	i int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func knuth(n int, fk float32) []Object {
	k := int(float32(n) * fk)
	objs := genData(n)
	res := make([]Object, 0, k)
	t, m := 0, 0
	for m < k {
		if float64(n-t)*rand.Float64() < float64(k-m) {
			res = append(res, objs[t])
			m++
		}
		t++
	}
	return res
}

func shufflePartial(n int, fk float32) []Object {
	k := int(float32(n) * fk)
	objs := genData(n)
	res := make([]Object, 0, k)
	for i := 1; i <= k; i++ {
		pos := rand.Int() % (n - i)
		res = append(res, objs[i+pos])
		objs[i], objs[i+pos] = objs[i+pos], objs[i]
	}
	return res
}

func genData(n int) []Object {
	objs := make([]Object, 0, n)
	for i := 0; i < n; i++ {
		objs = append(objs, Object{
			i: i,
		})
	}
	return objs
}
