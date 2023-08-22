package benchmark

import (
	"testing"
)

//var result []int
//
//const size = 5

const iterations = 1e8

func getRatingPlural(count int) string {
	if count == 1 {
		return "Rating"
	}
	return "Ratings"
}

func getReviewsPlural(count int) string {
	if count == 1 {
		return "Review"
	}
	return "Reviews"
}

func get(s string, count int) string {
	if count == 1 {
		return s
	}
	return s + "s"
}

func BenchmarkSeparate(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		getRatingPlural(1)
		getRatingPlural(2)
		getReviewsPlural(1)
		getReviewsPlural(2)
	}
}

func BenchmarkCommon(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		get("Rating", 1)
		get("Rating", 2)
		get("Review", 1)
		get("Review", 2)
	}
}
