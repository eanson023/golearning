package map_test

import "testing"

func TestInitMap(t *testing.T) {
	//初始化方式1
	map1 := map[string]int{"eanson": 999, "wdh": 998}
	t.Log(map1["eanson"])

	//初始化方式2
	map2 := map[string]int{}
	map2["rnm"] = 10
	t.Log(map2)
	//capacity
	map3 := make(map[string]int, 4)
	map3["1"] = 1
	t.Log(map3, len(map3))
}

func TestAccessNotExistingKey(t *testing.T) {
	map1 := map[int]int{}
	//输出结果是0
	t.Log(map1[1])
	map1[2] = 0
	//还是0
	t.Log(map1[2], len(map1))
	//问题:map中的值到底是不存在还是默认就是0值???
	//解决: v:值 ok:bool型 如果存在返回true 不存在不返回false
	key := 3
	if v, ok := map1[key]; ok {
		t.Logf("key %d value=%d", key, v)
	} else {
		t.Logf("key %d is not existing", key)
	}
}
func TestAccessNotExistingKey2(t *testing.T) {
	map1 := map[int]string{}
	t.Log(map1[1], len(map1))
	//map1[1] = ""
	//t.Log(map1[1], len(map1))
	if v, ok := map1[1]; ok {
		t.Log(v)
	} else {
		t.Log("not found")
	}
}

func TestTravelMap(t *testing.T) {
	map1 := map[string]int{"eanson": 999, "wdh": 998}
	for key, value := range map1 {
		t.Log(key, value)
	}
}
