package word

import (
	"testing"
)

func TestWord(t *testing.T) {
	filters := []string{
		"abc",
		"def",
		"efg",
		"가나다",
	}

	w := NewWordFilter()

	for _, f := range filters {
		w.addFilterWord(f)
	}

	if w.IsInclude("abdezzdxdegg12") == true {
		t.Error("err! 0")
	}
	if w.IsInclude("abCb123EfG33") == false {
		t.Error("err! 1")
	}
	if w.IsInclude("간다가나다 라마바사1233  321") == false {
		t.Error("err! 2")
	}
	rep1 := w.ToFilter("abcdefaddddabc")
	if rep1 != "******adddd***" {
		t.Error("err! 3", rep1)
	}
	rep2 := w.ToFilter("나다가나다가가가abc123")
	if rep2 != "나다***가가가***123" {
		t.Error("err! 4", rep2)
	}
}

func BenchmarkWord(b *testing.B) {
	filters := []string{
		"abc",
		"def",
		"efg",
		"가나다",
	}

	w := NewWordFilter()

	for _, f := range filters {
		w.addFilterWord(f)
	}

	for n := 0; n < b.N; n++ {
		rep2 := w.ToFilter("나다가나다가가가abc123")
		if rep2 != "나다***가가가***123" {
			b.Error("err! 4", rep2)
		}
	}
}
