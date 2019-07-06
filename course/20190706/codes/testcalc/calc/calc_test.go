package calc

import "testing"

func TestAdd(t *testing.T) {
	if 4 != Add(1, 3) {
		t.Error("1 + 3 != 2")
	}
}

func TestAddFlag(t *testing.T) {
	if -1 != Add(-1, 3) {
		t.Error("-1 + any != -1")
	}
}

func BenchmarkFact(b *testing.B) {
	for i := 1; i <= 10000; i++ {
		Fact(i)
	}
}
