package main

import (
	"testing"
)


func TestIfElse(t *testing.T) {

	R := IfElse([]bool{true,true,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = IfElse([]bool{true,true,true,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = IfElse([]bool{true,true,false,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = IfElse([]bool{true,true,false,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = IfElse([]bool{true,false,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = IfElse([]bool{true,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{true,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{true,false,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,true,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,true,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,true,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,true,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,false,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = IfElse([]bool{false,false,false,false})
	if R != FailString(){
		t.Fail()
	}


}

func TestSmallIfElse(t *testing.T) {

	R := SmallIfElse([]bool{true,true,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,true,true,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,true,false,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,true,false,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,false,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{true,false,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,true,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,true,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,true,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,true,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,false,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = SmallIfElse([]bool{false,false,false,false})
	if R != FailString(){
		t.Fail()
	}
}

func TestArraySolution(t *testing.T) {
	i := 1
	R := ArraySolution([]bool{true,true,true,true})
	if R != SuccedString(){
		t.Fail()
		t.Log(i)
		t.Log(R)

	}
	i++;
	R = ArraySolution([]bool{true,true,true,false})
	if R != SuccedString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,true,false,true})
	if R != SuccedString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,true,false,false})
	if R != SuccedString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,false,true,true})
	if R != SuccedString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,false,true,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,false,false,true})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{true,false,false,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,true,true,true})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,true,true,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,true,false,true})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,true,false,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,false,true,true})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,false,true,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,false,false,true})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
	R = ArraySolution([]bool{false,false,false,false})
	if R != FailString(){
		t.Fail()
		t.Log(i)
		t.Log(R)
	}
	i++;
}

func TestTreeSolution(t *testing.T) {
	SetupTreeBeforeHand()

	R := TreeSolution([]bool{true,true,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,true,true,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,true,false,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,true,false,false})
	if R != SuccedString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,false,true,true})
	if R != SuccedString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{true,false,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,true,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,true,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,true,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,true,false,false})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,false,true,true})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,false,true,false})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,false,false,true})
	if R != FailString(){
		t.Fail()
	}
	R = TreeSolution([]bool{false,false,false,false})
	if R != FailString(){
		t.Fail()
	}


}
