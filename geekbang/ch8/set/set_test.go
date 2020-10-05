package set

import "testing"

func TestSet(t *testing.T) {
	set := map[int]bool{}
	//添加元素
	set[1] = true
	//判断元素是否存在
	set[2] = false
	key := 2
	if _, ok := set[key]; ok {
		t.Log("key", key, "existing")
	} else {
		t.Log("key", key, " is not existing")
	}
	//个数
	t.Log(len(set))
	//删除元素 delete(map,指定删掉的key)
	delete(set, 1)
	t.Log("-----------")
	t.Log(len(set))
}
