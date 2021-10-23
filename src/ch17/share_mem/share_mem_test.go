package share_mem

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second) // 去掉后，程序运行太快，有些协程没有运行完，得不到预期的结果
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func TestRWMutex(t *testing.T) {
	var rwm sync.RWMutex
	var count int
	for i := 0; i < 5; i++ {
		go func() {
			rwm.RLock()
			t.Log(count)
			// go func() {
			// 	rwm.Lock()
			// 	count = rand.Intn(1000)
			// 	rwm.Unlock()
			// }()
			rwm.RUnlock()
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			rwm.Lock()
			rand.Seed(time.Now().Unix())
			count = rand.Intn(100)
			// t.Log(count)
			rwm.Unlock()
		}()
	}
	t.Log("test chang")
	time.Sleep(time.Second * 1)
}
