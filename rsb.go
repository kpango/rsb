package rsb

import (
	"math/rand"

	"github.com/kpango/fastime"
)

type Object struct {
	i int
}

var (
	f = fastime.New()
)

func knuth(n int, fk float32) []Object {
	k := int(float32(n) * fk)
	objs := genData(n)
	res := make([]Object, 0, k)
	t, m := 0, 0
	for m < k {
		rand.Seed(f.UnixNanoNow())
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
		rand.Seed(f.UnixNanoNow())
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
