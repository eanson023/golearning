package ch33

import (
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjectPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			if i == 0 {
				v.Name = "Eanson"
			}
			t.Logf("%p,%v\n", v, v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}

func Test2(t *testing.T) {
	var obj Object = *new(Object)
	obj.Name = "Eanson"
	t.Logf("%p,%v", &obj, obj)
	var obj2 *Object = new(Object)
	obj2.Name = "Eanson2"
	t.Logf("%p,%v\n", obj2, obj2)
}
