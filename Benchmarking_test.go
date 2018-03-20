package main

import (
	"testing"
	"time"
	"math/rand"
)



func BenchmarkIfElse(b *testing.B) {

	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	Answers := []bool{r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0 }

	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		IfElse(Answers)
	}
}

func BenchmarkSmallIfElse(b *testing.B) {
	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	Answers := []bool{r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0 }

	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		SmallIfElse(Answers)
	}
}

func BenchmarkArraySolution(b *testing.B) {

	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	Answers := []bool{r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0 }

	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		ArraySolution(Answers)
	}

}

func BenchmarkTreeSolution(b *testing.B) {

	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	Answers := []bool{r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0,r.Intn(2)!= 0 }
	SetupTreeBeforeHand()

	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		TreeSolution(Answers)
	}
}

