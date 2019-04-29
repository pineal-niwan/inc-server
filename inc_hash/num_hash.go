package inc_hash

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrHasNoGetNumFromKeyFunc = errors.New("no get num from key function in num hash")
)

//从存储介质中由key获取uint32
type GetNumByKeyFromStoreFunc func(key string) (uint32, error)

//从存储介质中获取uint32
type GetNumFromStoreFunc func() (uint32, error)

//单个的递增unit32
type NumWithLock struct {
	sync.Mutex
	//数字
	num uint32
	//是否被设置
	beSet bool
}

//递增操作--如果当前没有被设置，就去介质中获取
func (nL *NumWithLock) Inc(f GetNumFromStoreFunc) (uint32, error) {
	nL.Lock()
	if nL.beSet {
		//值已经被初始化过了，可以直接加
		nL.num++
		num := nL.num
		nL.Unlock()
		return num, nil
	} else {
		//没有被初始化过，需要从介质中获取
		num, err := f()
		if err != nil {
			nL.Unlock()
			//获取失败，不用设置标记位
			return 0, err
		} else {
			//获取成功，设置值被初始化的标记位
			num++
			nL.num = num
			nL.beSet = true
			nL.Unlock()
			return num, nil
		}
	}
}

//key-递增uint32的hash
type NumIncHash struct {
	numHash    sync.Map
	getNumFunc atomic.Value
}

//设置从介质获取数字的函数指针
func (nH *NumIncHash) Init(f GetNumByKeyFromStoreFunc) {
	nH.getNumFunc.Store(f)
}

func (nH *NumIncHash) Get(key string) (uint32, error) {
	//拿到num的指针，如果拿不到，就新设定一个
	numPtr := &NumWithLock{}
	actualPtr, _ := nH.numHash.LoadOrStore(key, numPtr)
	newNumPtr := actualPtr.(*NumWithLock)
	f := func() (uint32, error) {
		fPtr := nH.getNumFunc.Load()
		if fPtr == nil {
			return 0, ErrHasNoGetNumFromKeyFunc
		} else {
			fFunc := fPtr.(GetNumByKeyFromStoreFunc)
			return fFunc(key)
		}
	}
	return newNumPtr.Inc(f)

}
