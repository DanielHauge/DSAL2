package main

type TreeBranch struct {
	Value string
	Right *TreeBranch
	Left *TreeBranch
}

var RootBranchNode *TreeBranch


func SetupTreeBeforeHand(){
	RootBranchNode := TreeBranch{"root", MakeBranch(3), MakeBranch(3)}
	SetDecisions(RootBranchNode)
}

func TreeSolution(Answers []bool)string{

	var Leaf *TreeBranch
	Leaf = RootBranchNode

	for _, a := range Answers {
		if a{
			Leaf = Leaf.Left
		}else {
			Leaf = Leaf.Right
		}
	}
	return Leaf.Value
}

func MakeBranch(i int)*TreeBranch{
	if i == 0 {
		return &TreeBranch{FailString(),nil,nil}
	} else {
		return &TreeBranch{"N/A", MakeBranch(i-1), MakeBranch(i-1)}
	}
}

func SetDecisions(MyTree TreeBranch){
	MyTree.Left.Left.Left.Left.Value = SuccedString()
	MyTree.Left.Left.Left.Right.Value = SuccedString()
	MyTree.Left.Left.Right.Left.Value = SuccedString()
	MyTree.Left.Left.Right.Right.Value = SuccedString()
	MyTree.Left.Right.Left.Left.Value = SuccedString()

	RootBranchNode = &MyTree
}