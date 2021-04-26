package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var rwmu sync.RWMutex

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	rwmu.RLock()
	if v, ok := o.visitIPs[ip]; ok {
		rwmu.RUnlock()
		if time.Now().Before(v.Add(3 * time.Minute)) {
			return true
		}
	}
	rwmu.RUnlock()
	rwmu.Lock()
	o.visitIPs[ip] = time.Now()
	rwmu.Unlock()
	return false
}
func main() {
	var success int64
	ban := NewBan()
	//mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					//mu.Lock()
					atomic.AddInt64(&success, 1)
					//mu.Unlock()
				}
			}(j)
		}

	}
	wg.Wait()
	fmt.Println("success:", success)
}
