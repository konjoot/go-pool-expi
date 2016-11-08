package hybrid

import (
	"sync"
)

func Make() chan func() {
	in := make(chan func())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var j func()

		wg.Done()
		for j = range in {
			j()
		}
	}()
	wg.Wait()

	return in
}
