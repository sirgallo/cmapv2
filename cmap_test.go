package cmap

import (
	"testing"
)

func TestCMap(t *testing.T) {
	cMap := NewMap()

	t.Run("test put", func(t *testing.T) {
		cMap.Put([]byte("hello"), []byte("world"))
		cMap.Put([]byte("new"), []byte("wow!"))
		cMap.Put([]byte("again"), []byte("test!"))
		cMap.Put([]byte("woah"), []byte("random entry"))
		cMap.Put([]byte("key"), []byte("Saturday!"))
		cMap.Put([]byte("sup"), []byte("6"))
		cMap.Put([]byte("final"), []byte("the!"))
		cMap.Put([]byte("6"), []byte("wow!"))
		cMap.Put([]byte("asdfasdf"), []byte("add 10"))
		cMap.Put([]byte("asdfasdf"), []byte("123123")) // note same key, will update value
		cMap.Put([]byte("asd"), []byte("queue!"))
		cMap.Put([]byte("fasdf"), []byte("interesting"))
		cMap.Put([]byte("yup"), []byte("random again!"))
		cMap.Put([]byte("asdf"), []byte("hello"))
		cMap.Put([]byte("asdffasd"), []byte("uh oh!"))
		cMap.Put([]byte("fasdfasdfasdfasdf"), []byte("error message"))
		cMap.Put([]byte("fasdfasdf"), []byte("info!"))
		cMap.Put([]byte("woah"), []byte("done"))

		rootBitMap := cMap.Root().Bitmap()

		t.Logf("cMap after inserts")
		cMap.Root().PrintChildren()

		expectedBitMap := uint32(542198999)
		t.Logf("actual root bitmap: %d, expected root bitmap: %d\n", rootBitMap, expectedBitMap)
		t.Logf("actual root bitmap: %032b, expected root bitmap: %032b\n", rootBitMap, expectedBitMap)
		if expectedBitMap != rootBitMap {
			t.Errorf("actual bitmap does not match expected bitmap: actual(%032b), expected(%032b)\n", rootBitMap, expectedBitMap)
		}
	})

	t.Run("test retrieve", func(t *testing.T) {
		val1 := cMap.Get([]byte("hello"))
		expVal1 := "world"
		t.Logf("actual: %s, expected: %s", val1, expVal1)
		if string(val1) != expVal1 {
			t.Errorf("val 1 does not match expected val 1: actual(%s), expected(%s)\n", val1, expVal1)
		}

		val2 := cMap.Get([]byte("new"))
		expVal2 := "wow!"
		t.Logf("actual: %s, expected: %s", val2, expVal2)
		if string(val2) != expVal2 {
			t.Errorf("val 1 does not match expected val 1: actual(%s), expected(%s)\n", val2, expVal2)
		}

		val3 := cMap.Get([]byte("asdf"))
		expVal3 := "hello"
		t.Logf("actual: %s, expected: %s", val3, expVal3)
		if string(val3) != expVal3 {
			t.Errorf("val 1 does not match expected val 1: actual(%s), expected(%s)", val3, expVal3)
		}

		val4 := cMap.Get([]byte("asdfasdf"))
		expVal4 := "123123"
		t.Logf("actual: %s, expected: %s", val4, expVal4)
		if string(val4) != expVal4 {
			t.Errorf("val 1 does not match expected val 1: actual(%s), expected(%s)", val4, expVal4)
		}
	})

	t.Run("test delete", func(t *testing.T) {
		cMap.Delete([]byte("hello"))
		cMap.Delete([]byte("yup"))
		cMap.Delete([]byte("asdf"))
		cMap.Delete([]byte("asdfasdf"))
		cMap.Delete([]byte("new"))
		cMap.Delete([]byte("6"))

		rootBitMapAfterDelete := cMap.Root().Bitmap()
		t.Logf("bitmap of root after deletes: %032b\n", rootBitMapAfterDelete)
		t.Logf("bitmap of root after deletes: %d\n", rootBitMapAfterDelete)

		t.Log("hamt after deletes")
		cMap.Root().PrintChildren()

		expectedRootBitmapAfterDelete := uint32(536956102)
		t.Log("actual bitmap:", rootBitMapAfterDelete, "expected bitmap:", expectedRootBitmapAfterDelete)
		if expectedRootBitmapAfterDelete != rootBitMapAfterDelete {
			t.Errorf("actual bitmap does not match expected bitmap: actual(%032b), expected(%032b)\n", rootBitMapAfterDelete, expectedRootBitmapAfterDelete)
		}
	})

	t.Run("test get missing", func(t *testing.T) {
		m := NewMap()
		v := m.Get([]byte("does-not-exist"))
		if len(v) != 0 {
			t.Errorf("expected miss, got value=%q", v)
		}
	})

	t.Log("done")
}
