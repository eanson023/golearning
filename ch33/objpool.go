package ch33

import (
	"errors"
	"time"
)

//Object ... 对边定义的对象
type Object struct {
	Name string
}

//对象池
type ObjPool struct {
	bufChannel chan *Object
}

func NewObjectPool(poolSize int) *ObjPool {
	if poolSize <= 0 {
		panic(errors.New("poolsize must greater than 0"))
	}
	channel := make(chan *Object, poolSize)
	// 创poolsize个对象
	for i := 0; i < poolSize; i++ {
		channel <- &Object{}
	}
	return &ObjPool{channel}
}

//返回对象指针
func (pool *ObjPool) GetObj(timeout time.Duration) (*Object, error) {
	select {
	case ret := <-pool.bufChannel:
		// 只有返回空接口类型 不能返回其指针
		return ret, nil
		// 超时控制
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (pool *ObjPool) ReleaseObj(obj *Object) error {
	select {
	case pool.bufChannel <- obj:
		return nil
		// 当放不进去的时候上面会阻塞 那么就走下面 立即返回异常
	default:
		return errors.New("overflow")
	}
}
