package main

func IfElse (Answers []bool)string{
	results := ""

	// First : Hand ins made in time
	// Second : Attend Lectures
	// Third : Read books
	// Fourth : Make exercises


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
		} else { // Second Layer
			if Answers[2]{ // Third Layer
				if Answers[3]{ // Fourth Layer
					return SuccedString()
				} else { // Fourth Layer
					return FailString()
				}
			} else { // Third Layer
				if Answers[3]{ // Fourth Layer
					return FailString()
				} else { // Fourth Layer
					return FailString()
				}
			}
		}
	} else { // First Layer

		if Answers[1]{ // Second Layer
			if Answers[2]{ // Third Layer
				if Answers[3]{ // Fourth Layer
					return FailString()
				} else { // Fourth Layer
					return FailString()
				}
			} else { // Third Layer
				if Answers[3]{ // Fourth Layer
					return FailString()
				} else { // Fourth Layer
					return FailString()
				}
			}
		} else { // Second Layer
			if Answers[2]{ // Third Layer
				if Answers[3]{ // Fourth Layer
					return FailString()
				} else { // Fourth Layer
					return FailString()
				}
			} else { // Third Layer
				if Answers[3]{ // Fourth Layer
					return FailString()
				} else { // Fourth Layer
					return FailString()
				}
			}
		}
	}


	return results
}

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


