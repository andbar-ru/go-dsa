package dsa

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	keys := []int{231, 321, 212, 321, 433, 262}
	vals := []int{123, 432, 523, 43, 423, 111}
	ht := NewHashtable(len(keys))

	for i := 0; i < len(keys); i++ {
		ht.Insert(keys[i], vals[i])
	}
	ht.Delete(12)

	htStr := fmt.Sprint(ht)
	want := `[0]: 231: 123
[1]: /
[2]: 212: 523
[3]: 262: 111
[4]: /
[5]: /
[6]: 321: 43 -> 433: 423
`
	if htStr != want {
		t.Errorf("hashtable is %s, want %s", htStr, want)
	}
	if ht.Size() != 5 {
		t.Errorf("hashtable size = %d, want 5", ht.Size())
	}

	ht.Delete(212)
	htStr = fmt.Sprint(ht)
	want = `[0]: 231: 123
[1]: /
[2]: /
[3]: 262: 111
[4]: /
[5]: /
[6]: 321: 43 -> 433: 423
`
	if htStr != want {
		t.Errorf("hashtable after deleting is %s, want %s", htStr, want)
	}
	if ht.Size() != 4 {
		t.Errorf("hashtable size = %d, want 4", ht.Size())
	}

	if _, ok := ht.Get(12); ok {
		t.Errorf("ht.Get(12) is ok, want not ok")
	}
	if v, _ := ht.Get(262); v != 111 {
		t.Errorf("ht.Get(262) = %d, want 111", v)
	}
	if v, _ := ht.Get(433); v != 423 {
		t.Errorf("ht.Get(433) = %d, want 423", v)
	}
}
