package inc_hash

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count uint32

func _getByKey(key string) (uint32, error) {
	atomic.AddUint32(&count, 1)
	time.Sleep(20*time.Millisecond)
	return 0, nil
}

func TestNumHash_Get(t *testing.T) {
	var nh NumIncHash
	var keyList [][]string
	var waitGrp sync.WaitGroup

	nh.Init(_getByKey)

	lKey := 100

	keyList = make([][]string, lKey)
	//key数量
	for i := 0; i < 100; i++ {
		k := i
		for j := 0; j < 100; j++ {
			if k >= lKey {
				k = 0
			}
			keyList[i] = append(keyList[i], fmt.Sprintf("%d", k))
			k++
		}
	}

	waitGrp.Add(1000)

	//go程数量
	for i := 0; i < 1000; i++ {
		go func() {
			k := i % lKey
			l := len(keyList[k])
			for j := 0; j < l; j++ {
				nh.Get(keyList[k][j])
			}
			waitGrp.Done()
		}()
	}

	waitGrp.Wait()
	x := atomic.LoadUint32(&count)
	t.Logf("all get key from store time:%+v", x)

}