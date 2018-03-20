package main


// First : Hand ins made in time
// Second : Attend Lectures
// Third : Read books
// Fourth : Make exercises

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
