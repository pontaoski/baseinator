package main

import "testing"

func benchBase(base Base, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NameBase(base)
	}
}

func Benchmark420(b *testing.B)    { benchBase(420, b) }
func Benchmark720720(b *testing.B) { benchBase(720720, b) }
func BenchmarkPrime(b *testing.B)  { benchBase(193*197*2, b) }
func BenchmarkBig(b *testing.B)    { benchBase((1<<31)-1, b) }
func BenchmarkBigger(b *testing.B) { benchBase((1<<63)-1, b) }
