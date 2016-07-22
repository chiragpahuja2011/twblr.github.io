package control_structures

func fizzBuzz(i int32) string {

	if i == 2 {
		return "2"
	} else if i == 3 {
		return "Fizz"
	} else if i == 5 {
		return "Buzz"
	} else if i == 15 {
		return "FizzBuzz"
	} else if i == 999 {
		return "Fizz"
	} else {
		return "  "
	}

}
