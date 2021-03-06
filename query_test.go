package main

import "testing"

func benchmarkSearch(search string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		queryRepo(append([]string{}, search))
	}
}
func BenchmarkSearchSimpleTopDown(b *testing.B) {
	config.SortMode = TopDown
	benchmarkSearch("chromium", b)
}

func BenchmarkSearchSimpleBottomUp(b *testing.B) {
	config.SortMode = BottomUp
	benchmarkSearch("chromium", b)
}

func BenchmarkSearchComplexTopDown(b *testing.B) {
	config.SortMode = TopDown
	benchmarkSearch("linux", b)
}
func BenchmarkSearchComplexBottomUp(b *testing.B) {
	config.SortMode = BottomUp
	benchmarkSearch("linux", b)
}
