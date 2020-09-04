package ch25

import (
	"testing"
	// package起别名
	cm "github.com/easierway/concurrent_map"
)

func Test1(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
