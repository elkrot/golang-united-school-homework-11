package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res_ := make([]user, 0, n)
	var wg sync.WaitGroup

	k := int(n) / int(pool)

	for i := 0; i < int(pool); i++ {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()
			tmp_ := make([]user, 0, k)
			b := j * k
			d := b + k

			for c := b; c < d; c++ {
				z := int64(c)
				user := getOne(z)
				tmp_ = append(tmp_, user)
			}
			res_ = append(res_, tmp_...)

		}(i)

	}

	wg.Wait()

	return res_
}
