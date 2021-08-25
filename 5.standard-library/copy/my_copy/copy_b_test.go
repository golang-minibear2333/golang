package my_copy

import "testing"

func BenchmarkFoo_Duplicate(b *testing.B) {
	b.StopTimer()
	var a = 1
	var t1 = Foo{IntPtr: &a}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = t1.Duplicate()
	}
}

func BenchmarkDeepCopyByGob(b *testing.B) {
	b.StopTimer()
	var a = 1
	var t1 = Foo{IntPtr: &a}
	t2 := Foo{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = DeepCopyByGob(&t2, t1)
	}
}

func BenchmarkDeepCopyByJson(b *testing.B) {
	b.StopTimer()
	var a = 1
	var t1 = Foo{IntPtr: &a}
	t2 := Foo{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = DeepCopyByJson(&t2, t1)
	}
}
