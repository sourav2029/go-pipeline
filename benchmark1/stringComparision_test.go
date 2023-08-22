package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringOps(b *testing.B) {

	tests := []struct {
		description   string
		first, second string
	}{
		{
			description: "both strings equal",
			first:       "aaaadqwqw",
			second:      "aaaadqwdqwdd",
		},
		{
			description: "one string in caps",
			first:       "aaaa",
			second:      "AAAA",
		},
		{
			description: "weird casing situation",
			first:       "aAaA",
			second:      "AaAa",
		},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("%s::equality op", tt.description), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkStringEqualsOperation(tt.first, tt.second, b)
			}
		})

		b.Run(fmt.Sprintf("%s::strings equal fold", tt.description), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkStringsEqualFold(tt.first, tt.second, b)
			}
		})
	}
}

func benchmarkStringEqualsOperation(first, second string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.ToLower(first) == strings.ToLower(second)
	}
}

func benchmarkStringsEqualFold(first, second string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.EqualFold(first, second)
	}
}
