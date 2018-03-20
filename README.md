# 2. Compulsory algorithm Assignment - Trees
This repository is for a hand-in for Software development (PBA) Datastructure and Algorithms course. Daniel (cph-dh136)

## Description
This assignment is about trees and requires to make a decision binary tree. The assignment is based on this ressourse: [Here](https://i.gyazo.com/e7ff5d1cca5cc5bd016f018140035347.png)

## Assignment
To start off with, I have an idea of different ways to solve the requirements of the assignment. I know the assigmment is to solve the requirements with a binary tree as a solution. However, i thought it would be cool to see more solutions and actully reflect on different ways of doing the same thing. So the 3 different solutions i had in mind was, alot of nested if-else blocks, to use array as conceptual tree, a linked list kind of thing.

The full assignment/solution is also uploaded at github [Here](https://github.com/DanielHauge/DSAL2)

As an overview i've made a decision tree diagram in Software idea modeller which can be seen below:
[![https://gyazo.com/a1ef72145a322c1b75684090b351a4c6](https://i.gyazo.com/a1ef72145a322c1b75684090b351a4c6.png)](https://gyazo.com/a1ef72145a322c1b75684090b351a4c6)
The full software idea modeller project is on the github repository.

#### If-Else solution
The if-else solution can be found in [IfStatementSolution.go](https://github.com/DanielHauge/DSAL2/blob/master/IfStatementSolution.go)

But essentially it is quite simple, but very long and tedious to program. Here is a cutout from the solution to give a good idea of how it is implemented:
```golang
if Answers[0]{ // First Layer
		if Answers[1]{ // Second Layer
			if Answers[2]{ // Third Layer
				if Answers[3]{ // Fourth Layer
					return SuccedString()
				} else { // Fourth Layer
					return SuccedString()
				}
			} else { // Third Layer
				if Answers[3]{ // Fourth Layer
					return SuccedString()
				} else { // Fourth Layer
					return SuccedString()
				}
			}
			..... Goes on from here
```
This is long and tedious, so i came up with a shorter version, that doesn't realy implement the conceptual idea of a tree.
```go
func SmallIfElse (Answers []bool)string{

	if !Answers[0]{
		return FailString()
	} else if !Answers[1]&&!Answers[2]{
		return FailString()
	} else if !Answers[1]&&!Answers[3]{
		return FailString()
	} else {
		return SuccedString()
	}
}
```
This one basicly just checks the conditions and returns either the failstring or succedstring.

#### Array solution
The Array solution can be found in [ArraySolution.go](https://github.com/DanielHauge/DSAL2/blob/master/ArraySolution.go)

This solution could be defined alot more to be a tree structure than previus. This solution is very slim and not alot of code for it to work, however the tree(Array) is manually constructed.
```go
func ArraySolution(Answers []bool)string{
	Index := 0
	MyArrayAsConceptualTree := []string{"root", "yes", "no", "yes", "no", "yes", "no", "yes", "no", SuccedString(), SuccedString(), SuccedString(), SuccedString(), SuccedString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString(), FailString()}

	for i, a := range Answers {
		if a{
			Index += 2*i
		}else {
			Index += 2*i+1
		}
	}
	return MyArrayAsConceptualTree[Index]

}
```
As seen above, this is all that it required. What it basicly does, it to iterate over the answers which come in a form of a boolean array, and jumps(Set index) to specific places in the array based on the answers given.

#### Binary tree solution
The Array solution can be found in [TreeSolution.go](https://github.com/DanielHauge/DSAL2/blob/master/TreeSolution.go)

This one takes advantage of recursiveness and linking structs together. The struct for the tree which is used is as below:
```go
type TreeBranch struct {
	Value string
	Right *TreeBranch
	Left *TreeBranch
}
```
With recursiveness, it can make a balanced tree with the function as below:
```go
RootBranchNode := TreeBranch{"root", MakeBranch(3), MakeBranch(3)}
```
```go
func MakeBranch(i int)*TreeBranch{
	if i == 0 {
		return &TreeBranch{FailString(),nil,nil}
	} else {
		return &TreeBranch{"N/A", MakeBranch(i-1), MakeBranch(i-1)}
	}
}
```
With this construction of the tree, i've set all leafs to be the failstring, i've used this function to set the decisions afterwards, i've done this manually tho.
```go
func SetDecisions(MyTree TreeBranch){
	MyTree.Left.Left.Left.Left.Value = SuccedString()
	MyTree.Left.Left.Left.Right.Value = SuccedString()
	MyTree.Left.Left.Right.Left.Value = SuccedString()
	MyTree.Left.Left.Right.Right.Value = SuccedString()
	MyTree.Left.Right.Left.Left.Value = SuccedString()

	RootBranchNode = &MyTree
}
```
To traverse this tree, i'm starting at the root, and traversing down the tree, by moving the pointer to the children brances, until there is no more answers to check, and then returns the value of the leaf.
```go
    var Leaf *TreeBranch
	Leaf = &RootBranchNode

	for _, a := range Answers {
		if a{
			Leaf = Leaf.Left
		}else {
			Leaf = Leaf.Right
		}
	}
	return Leaf.Value
```
With this, it can return the value of the leaf which was manualy set to the succed or recursevely to fail string.

### Results
I've made a small example program that just tests each solution. In [Main.go](https://github.com/DanielHauge/DSAL2/blob/master/TreeSolution.go)

The main.go example is:
```go
    log.Println(TreeSolution([]bool{true,true,true,true}))
	log.Println(TreeSolution([]bool{true,true,false,true}))
	log.Println(TreeSolution([]bool{true,true,true,false}))
	log.Println(TreeSolution([]bool{true,false,true,true}))
	log.Println(TreeSolution([]bool{true,false,false,true}))
	log.Println(TreeSolution([]bool{true,false,false,false}))
	log.Println(TreeSolution([]bool{false,false,false,true}))
```
The console output is:
```
2018/03/16 16:02:44 You should be able to pass the exam
2018/03/16 16:02:44 You should be able to pass the exam
2018/03/16 16:02:44 You should be able to pass the exam
2018/03/16 16:02:44 You should be able to pass the exam
2018/03/16 16:02:44 You could easily fail the exam
2018/03/16 16:02:44 You could easily fail the exam
2018/03/16 16:02:44 You could easily fail the exam
```


## Testing
I wanted to make sure all solutions worked, therefor i've made some tests to check if they work as inteded.

All test code can be found in [Tests_test.go](https://github.com/DanielHauge/DSAL2/blob/master/Tests_test.go)

I tested all possible configurations of the answers that you can give on all 3 solutions including short version of if-else. Everything worked exactly as inteded.
[![https://gyazo.com/3a136181f3f89e28436011cbd5023679](https://i.gyazo.com/3a136181f3f89e28436011cbd5023679.png)](https://gyazo.com/3a136181f3f89e28436011cbd5023679)

A cutout of some testing:
```go
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
	........... And so on. for all configurations.
```

## Benchmarking
Benchmarking code can be found in [Benchmarking_test.go](https://github.com/DanielHauge/DSAL2/blob/master/Benchmarking_test.go)
#### Benchmark 1.
The initial benchmarking showed a few expected results, and some other unexpected results. First of all, the expected results i was expecting was, that the shorter if-else was faster than the longer one. Since it just checks a few things and know what to spitout then. And also the array solution was expected, since it needs to go through some initialization and "jumpin" arround in the array to find the right place in the end.

But what i didn't expect was that the tree solution was taking very long time. But i think i know why. It is because i was actully benchmarking the whole process of constructing the tree and then traversion it and more. Where the if statements are staic and does not use recursiveness and more, and so on.

Benchmarking results:
```go
goos: windows
goarch: amd64
500000000	         3.87 ns/op    // If-else long
2000000000	         1.11 ns/op    // If-else small
100000000	        22.9 ns/op     // Array
1000000	      1172 ns/op           // Tree
PASS
ok  	_/C_/Users/Animc/Desktop/Repos/DSAL2	8.136s
```
It goes from top the button. Long if-else, small if-else, array, tree.

I will refractor the code to have more usefull and relevant results.

#### Benchmarking 2.
Now i changed the code as follows: 
```go
func SetupTreeBeforeHand(){
	RootBranchNode := TreeBranch{"root", MakeBranch(3), MakeBranch(3)}
	SetDecisions(RootBranchNode)
}
```
I made a setuptreebeforehand which basicly constructs the tree and all, where the call i'm benchmarking on, only looks at searching the tree.

With this change. the results are alot closer to my expectations:
```go
500000000	         4.40 ns/op    // If-else long
2000000000	         1.92 ns/op    // If-else small
100000000	        22.9 ns/op     // Array
200000000	         7.54 ns/op    // Tree
PASS
ok  	_/C_/Users/Animc/Desktop/Repos/DSAL2	11.285s
```
Now it is only roughly 7,5 nano seconds instead of around 1000.

### Remarks/notes
I find the requirements of the assignment a little funny, since as i've understood it, it is essentially just conditioning. Where i've gotten the impression that binary tree's and also other kinds of tree structures is about data, and how it is very fast to search for data in a tree structure.
So i think a just as interesting or more interesting requirement, would be to have something to do with searching for values rather than a Conditional decision tree. However i think it is good, since it made me think that in all circumstances there is alot of tools and solution that can furfill the requirements, it is up to me to choose the best one.

Another note is, that the tree solution contains alot of ReferencePointing/Targeting and more, to avoid reconstructing trees. So it doesn't generate or construct lots of trees for any operation, this makes it more efficient. It is annotated with '&' for targeting and '*' for pointer/reference. This enables, recursiveness, things for concurrency and more.


## Conclusion
As it looks to me, when it is purely conditioning on predetermined answers, it is far more efficient to do if-else. But if the decision becomes large, then it is tedious to program, aswell as we get very smelly code very fast with alot of if-else statements. However in some cases it is possible to shorten it down, which may help.

But when it comes to storing data, it is a little different. Doing if-else statements realy doesn't search in stored data, it is static. With an array as a conceptual tree can work better this way. This way, it is just a matter of editing or configuring the array in the right manner for it be effective to search it, aswell with linking structs as the tree solution. But the tree structure is more effective in the way i've programmed it, so it might indicate the linking of structs are more efficient, than jumping and keeping track of indexes in an array.