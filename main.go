package main

import (
	// "./each"
	// "./hybrid"
	// "./oneof"
	"runtime"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(320000000)
	for i := 0; i < 320000000; i++ {
		func() {
			wg.Done()
		}()
	}
	wg.Wait()
	println(`SPEED ITER (ALL) 320000000 i:`, time.Since(t).String())

	// t = time.Now()
	// wgb := &sync.WaitGroup{}
	// wgb.Add(320000000)
	// for i := 0; i < 320000000; i++ {
	// 	go func() {
	// 		wgb.Done()
	// 	}()
	// }
	// wgb.Wait()
	// println(`NATIVE MIX (ALL) 320000000 tasks:`, time.Since(t).String())

	// t = time.Now()
	// wgc := &sync.WaitGroup{}
	// wgc.Add(320000000)
	// for i := 0; i < 320000000; i++ {
	// 	go wgc.Done()
	// }
	// wgc.Wait()
	// println(`NATIVE MIX 2 (ALL) 320000000 tasks:`, time.Since(t).String())

	// t = time.Now()
	// epool := each.MakePool(320000000)
	// println(`EACH (START) 320000000 procs:`, time.Since(t).String())

	// wg := &sync.WaitGroup{}
	// wg.Add(320000000)
	// tp := time.Now()
	// for i := 0; i < 320000000; i++ {
	// 	epool.Go(func() int {
	// 		defer wg.Done()
	// 		return i
	// 	})
	// }
	// wg.Wait()
	// println(`EACH (PROC) 320000000 tasks:`, time.Since(tp).String())

	// t = time.Now()
	// epool.Stop()
	// println(`EACH (STOP) 320000000 procs:`, time.Since(t).String())

	// //-------------------------------------------------------

	// t = time.Now()
	// opool := oneof.MakePool(320000000)
	// println(`ONEOF (START) 320000000 procs:`, time.Since(t).String())

	// wgo := &sync.WaitGroup{}
	// wgo.Add(320000000)
	// tpo := time.Now()
	// for i := 0; i < 320000000; i++ {
	// 	opool.Go(func() int {
	// 		defer wgo.Done()
	// 		return i
	// 	})
	// }
	// wgo.Wait()
	// println(`ONEOF (PROC) 320000000 tasks:`, time.Since(tpo).String())

	// t = time.Now()
	// opool.Stop()
	// println(`ONEOF (STOP) 320000000 procs:`, time.Since(t).String())

	t = time.Now()

	wgh := &sync.WaitGroup{}

	num := runtime.NumCPU() * 8

	wgh.Add(320000000)

	for i := 0; i < num; i++ {

		ch := make(chan struct{}, num)

		go func() {
			for range ch {
				wgh.Done()
			}
		}()

		go func() {
			for i := 0; i < 320000000; i = i + num {
				ch <- struct{}{}
			}
		}()

	}
	wgh.Wait()

	println(`HYBRID (PROC) 320000000 procs:`, time.Since(t).String())
}
