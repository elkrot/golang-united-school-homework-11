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
	var wg sync.WaitGroup
	sem := make(chan user, pool)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		go func(j int) {
			user := getOne(int64(j))
			sem <- user
			res = append(res, <-sem)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
