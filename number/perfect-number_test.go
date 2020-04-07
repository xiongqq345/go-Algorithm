package number

import "testing"

func TestPerfectNum(t *testing.T) {
	if !IsPerfectNum(6) {
		t.Error(`IsPerfectNum(6) = false`)
	}
	if !IsPerfectNum(28) {
		t.Error(`IsPerfectNum(28) = false`)
	}
	if !IsPerfectNum(496) {
		t.Error(`IsPerfectNum(496) = false`)
	}
}

func BenchmarkIsPerfectNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPerfectNum(8128)
	}
}
